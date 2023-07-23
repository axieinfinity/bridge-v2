package task

import (
	"crypto/ecdsa"
	"github.com/axieinfinity/bridge-v2/stats"
	"math/big"
	"sort"
	"strings"
	"time"

	ethGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	roninGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
	roninValidator "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/validator"
	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/models"
	"github.com/axieinfinity/bridge-v2/stores"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/signer/core"
)

const (
	ErrSigAlreadySubmitted   = "execution reverted: BOsGovernanceRelay: query for outdated period"
	ErrOutdatedPeriod        = "execution reverted: BOsGovernanceProposal: query for outdated period"
	ErrOperatorsAlreadyVoted = "execution reverted: BridgeOperatorsBallot: bridge operator set is already voted"
)

type task struct {
	util              utils.Utils
	task              *models.Task
	store             stores.BridgeStore
	validator         *ecdsa.PrivateKey
	client            bind.ContractBackend
	ethClient         bind.ContractBackend
	contracts         map[string]string
	chainId           *big.Int
	maxTry            int
	taskType          string
	listener          bridgeCore.Listener
	releaseTasksCh    chan int
	gasLimitBumpRatio uint64
}

func newTask(
	listener bridgeCore.Listener,
	client,
	ethClient bind.ContractBackend,
	store stores.BridgeStore,
	chainId *big.Int,
	contracts map[string]string,
	maxTry int,
	taskType string,
	releaseTasksCh chan int,
	util utils.Utils,
	gasLimitBumpRatio uint64,
) *task {
	return &task{
		util:              util,
		task:              nil,
		store:             store,
		client:            client,
		ethClient:         ethClient,
		contracts:         contracts,
		chainId:           chainId,
		maxTry:            maxTry,
		taskType:          taskType,
		listener:          listener,
		releaseTasksCh:    releaseTasksCh,
		gasLimitBumpRatio: gasLimitBumpRatio,
	}
}

func (r *task) collectTask(t *models.Task) {
	log.Debug("Received new task", "id", t.ID, "status", t.Status, "type", t.Type)
	if t.Type == r.taskType {
		r.task = t
	}
}

func (r *task) send() {
	switch r.taskType {
	case VOTE_BRIDGE_OPERATORS_TASK:
		r.sendTransaction(r.voteBridgeOperatorsBySignature)
	case RELAY_BRIDGE_OPERATORS_TASK:
		r.sendTransaction(r.relayBridgeOperators)
	}
}

func (r *task) sendTransaction(sendTx func(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction)) {
	if r.task == nil {
		return
	}

	var txHash string

	doneTasks, processingTasks, failedTasks, transaction := sendTx(r.task)

	if transaction != nil {
		log.Debug("[task] Transaction", "type", r.taskType, "hash", transaction.Hash().Hex())
		go updateTasks(r.store, processingTasks, STATUS_PROCESSING, transaction.Hash().Hex(), time.Now().Unix(), r.releaseTasksCh)
		_ = metrics.Pusher.IncrGauge(metrics.ProcessingTaskMetric, len(processingTasks))
	}
	go updateTasks(r.store, doneTasks, STATUS_DONE, txHash, 0, r.releaseTasksCh)
	go updateTasks(r.store, failedTasks, STATUS_FAILED, txHash, 0, r.releaseTasksCh)
	_ = metrics.Pusher.IncrCounter(metrics.SuccessTaskMetric, len(doneTasks))
	_ = metrics.Pusher.IncrCounter(metrics.FailedTaskMetric, len(failedTasks))
}

func (r *task) voteBridgeOperatorsBySignature(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	log.Info("[RoninTask][BridgeOperatorSetCallback] Processing task")

	event, err := r.unpackBridgeOperatorSetUpdatedEvent(task)
	if err != nil {
		task.LastError = err.Error()
		stats.SendErrorToStats(r.listener, err)
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	sort.Sort(BridgeOperatorsSorter(event.BridgeOperators))

	// create caller
	roninGovernanceTransactor, err := roninGovernance.NewGovernance(common.HexToAddress(r.contracts[GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		stats.SendErrorToStats(r.listener, err)
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	voted, err := roninGovernanceTransactor.BridgeOperatorsVoted(nil, event.Period, event.Epoch, r.listener.GetVoterSign().GetAddress())
	if err != nil {
		task.LastError = err.Error()
		stats.SendErrorToStats(r.listener, err)
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	if voted {
		log.Debug("[RoninTask][BridgeOperatorSetCallback] Bridge already voted", "period", event.Period)
		doneTasks = append(doneTasks, task)
		return
	}

	syncedInfo, err := roninGovernanceTransactor.LastSyncedBridgeOperatorSetInfo(nil)
	if err != nil {
		task.LastError = err.Error()
		stats.SendErrorToStats(r.listener, err)
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	sort.Sort(BridgeOperatorsSorter(syncedInfo.Operators))

	isValidatorSetShouldUpdate := event.Period.Cmp(syncedInfo.Period) >= 0 && event.Epoch.Cmp(syncedInfo.Epoch) > 0
	log.Info("[RoninTask][BridgeOperatorSetCallback] Should vote", "value", isValidatorSetShouldUpdate, "event", event, "syncedInfo", syncedInfo)
	if !isValidatorSetShouldUpdate {
		doneTasks = append(doneTasks, task)
		return doneTasks, nil, nil, nil
	}

	// otherwise add task to processingTasks to adjust after sending transaction
	processingTasks = append(processingTasks, task)

	var bridgeOperators []interface{}
	for _, address := range event.BridgeOperators {
		bridgeOperators = append(bridgeOperators, address.Hex())
	}
	opts := &signDataOpts{
		ChainID: r.chainId,
		SignTypedDataCallback: func(typedData core.TypedData) (hexutil.Bytes, error) {
			return r.util.SignTypedData(typedData, r.listener.GetVoterSign())
		},
	}
	signature, err := signBridgeOperatorsBallot(opts, event.Period.Int64(), event.Epoch.Int64(), bridgeOperators)
	if err != nil {
		task.LastError = err.Error()
		stats.SendErrorToStats(r.listener, err)
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	signatureStruct := parseSignatureAsRsv(signature)
	log.Debug("[RoninTask][BridgeOperatorSetCallback] Prepared data", "r", common.Bytes2Hex(signatureStruct.R[:]), "s", common.Bytes2Hex(signatureStruct.S[:]), "v", signatureStruct.V, "period", event.Period.Int64(), "bridgeOperators", bridgeOperators)

	tx, err = r.util.SendContractTransaction(r.listener.GetVoterSign(), r.chainId, r.gasLimitBumpRatio, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return roninGovernanceTransactor.VoteBridgeOperatorsBySignatures(opts, roninGovernance.BridgeOperatorsBallotBridgeOperatorSet{
			Period:    event.Period,
			Epoch:     event.Epoch,
			Operators: event.BridgeOperators,
		}, []roninGovernance.SignatureConsumerSignature{
			signatureStruct,
		})
	})
	if err != nil {
		stats.SendErrorToStats(r.listener, err)
		switch err.Error() {
		case ErrOutdatedPeriod:
			log.Warn("[RoninTask][BridgeOperatorSetCallback] Bridge operators period outdated")
			doneTasks = append(doneTasks, task)
			return doneTasks, nil, nil, nil
		case ErrOperatorsAlreadyVoted:
			log.Warn("[RoninTask][BridgeOperatorSetCallback] Bridge operators have already voted")
			doneTasks = append(doneTasks, task)
			return doneTasks, nil, nil, nil
		default:
			log.Error("[RoninTask][BridgeOperatorSetCallback] Send transaction error", "err", err)
			task.LastError = err.Error()
			failedTasks = append(failedTasks, task)
			return nil, nil, failedTasks, nil
		}
	}

	log.Debug("[RoninTask][BridgeOperatorSetCallback] Vote bridge operators", "hash", tx.Hash().Hex())
	return
}

func (r *task) relayBridgeOperators(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	log.Info("[RoninTask][BridgeOperatorsApprovedCallback] Processing task")
	// create caller
	roninGovernanceCaller, err := roninGovernance.NewGovernanceCaller(common.HexToAddress(r.contracts[GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		stats.SendErrorToStats(r.listener, err)
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	ethGovernanceTransactor, err := ethGovernance.NewGovernance(common.HexToAddress(r.contracts[ETH_GOVERNANCE_CONTRACT]), r.ethClient)
	if err != nil {
		stats.SendErrorToStats(r.listener, err)
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	event, err := r.unpackBridgeOperatorsApprovedEvent(task)
	if err != nil {
		stats.SendErrorToStats(r.listener, err)
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	log.Debug("[RoninTask][BridgeOperatorsApprovedCallback] Unpacked event", "event", event)
	sort.Sort(BridgeOperatorsSorter(event.Operators))

	// Ethereum call
	ethSyncedInfo, err := ethGovernanceTransactor.LastSyncedBridgeOperatorSetInfo(nil)
	if err != nil {
		stats.SendErrorToStats(r.listener, err)
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	sort.Sort(BridgeOperatorsSorter(ethSyncedInfo.Operators))

	isValidatorSetEquals := EqualOperatorSet(event.Operators, ethSyncedInfo.Operators)
	log.Info("[RoninTask][BridgeOperatorsApprovedCallback] Check operators set", "isEqual", isValidatorSetEquals, "event", event, "ethSyncedInfo", ethSyncedInfo)
	if isValidatorSetEquals {
		doneTasks = append(doneTasks, task)
		return doneTasks, nil, nil, nil
	}

	// otherwise add task to processingTasks to adjust after sending transaction
	processingTasks = append(processingTasks, task)

	signatures, err := roninGovernanceCaller.GetBridgeOperatorVotingSignatures(nil, event.Period, event.Epoch)
	if err != nil {
		stats.SendErrorToStats(r.listener, err)
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	log.Info("[RoninTask][BridgeOperatorsApprovedCallback] Bridge operator voting signatures", "signatures", signatures.Signatures, "voters", signatures.Voters)

	for i := 0; i < len(signatures.Signatures)-1; i++ {
		for j := i + 1; j < len(signatures.Signatures); j++ {
			if strings.ToLower(signatures.Voters[j].Hex()) < strings.ToLower(signatures.Voters[i].Hex()) {
				signatures.Voters[j], signatures.Voters[i] = signatures.Voters[i], signatures.Voters[j]
				signatures.Signatures[j], signatures.Signatures[i] = signatures.Signatures[i], signatures.Signatures[j]
			}
		}
	}

	var ethSignatures []ethGovernance.SignatureConsumerSignature
	for _, sig := range signatures.Signatures {
		if sig.V > 0 {
			ethSignatures = append(ethSignatures, ethGovernance.SignatureConsumerSignature{
				V: sig.V,
				R: sig.R,
				S: sig.S,
			})
		}
	}

	ethListener := r.listener.GetListener("Ethereum")
	ethChainId, err := ethListener.GetChainID()
	if err != nil {
		task.LastError = err.Error()
		stats.SendErrorToStats(r.listener, err)
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	tx, err = r.util.SendContractTransaction(r.listener.GetRelayerSign(), ethChainId, r.gasLimitBumpRatio, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return ethGovernanceTransactor.RelayBridgeOperators(opts, ethGovernance.BridgeOperatorsBallotBridgeOperatorSet{
			Period:    event.Period,
			Epoch:     event.Epoch,
			Operators: event.Operators,
		}, ethSignatures)
	})
	if err != nil {
		stats.SendErrorToStats(r.listener, err)
		// Prevent retry submit signature if the signature was already submitted
		switch err.Error() {
		case ErrSigAlreadySubmitted:
			log.Warn("[RoninTask][BridgeOperatorsApprovedCallback] Bridge operators already submitted")
			doneTasks = append(doneTasks, task)
			return doneTasks, nil, nil, nil
		default:
			log.Error("[RoninTask][BridgeOperatorsApprovedCallback] Send transaction error", "err", err)
			task.LastError = err.Error()
			failedTasks = append(failedTasks, task)
			return nil, nil, failedTasks, nil
		}
	}

	log.Debug("[RoninTask][BridgeOperatorsApprovedCallback] Relay bridge operators", "hash", tx.Hash().Hex())
	return
}

func (r *task) unpackBridgeOperatorSetUpdatedEvent(task *models.Task) (*roninValidator.ValidatorBridgeOperatorSetUpdated, error) {
	roninEvent := new(roninValidator.ValidatorBridgeOperatorSetUpdated)
	roninValidatorAbi, err := roninValidator.ValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	log.Debug("Bridge operator set updated", "data", task.Data)
	if err = r.util.UnpackLog(*roninValidatorAbi, roninEvent, "BridgeOperatorSetUpdated", common.Hex2Bytes(task.Data)); err != nil {
		return nil, err
	}

	return roninEvent, nil
}

func (r *task) unpackBridgeOperatorsApprovedEvent(task *models.Task) (*roninGovernance.GovernanceBridgeOperatorsApproved, error) {
	roninEvent := new(roninGovernance.GovernanceBridgeOperatorsApproved)
	roninGovernanceAbi, err := roninGovernance.GovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	log.Debug("Bridge operators approved", "data", task.Data)
	if err = r.util.UnpackLog(*roninGovernanceAbi, roninEvent, "BridgeOperatorsApproved", common.Hex2Bytes(task.Data)); err != nil {
		return nil, err
	}

	return roninEvent, nil
}

type signDataOpts struct {
	ChainID               *big.Int
	SignTypedDataCallback func(typedData core.TypedData) (hexutil.Bytes, error)
}

func createSalt(chainID *big.Int) (*common.Hash, error) {
	uint256ty, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return nil, err
	}
	stringTy, err := abi.NewType("string", "string", nil)
	if err != nil {
		return nil, err
	}
	args := abi.Arguments{
		{
			Type: stringTy,
		},
		{
			Type: uint256ty,
		},
	}

	bytes, err := args.Pack("RONIN_GOVERNANCE_ADMIN", chainID)
	if err != nil {
		return nil, err
	}

	hash := crypto.Keccak256Hash(bytes)

	return &hash, nil
}

func signBridgeOperatorsBallot(opts *signDataOpts, period, epoch int64, bridgeOperators interface{}) ([]byte, error) {
	salt, err := createSalt(opts.ChainID)
	if err != nil {
		log.Error("Can not encode salt key", "error", err)
		return nil, err
	}
	log.Debug("Generated salt key", "salt", salt.Hex(), "chainID", opts.ChainID)

	bridgeOperatorsBallotTypes := core.TypedData{
		Types: core.Types{
			"EIP712Domain": []core.Type{
				{
					Name: "name", Type: "string",
				},
				{
					Name: "version", Type: "string",
				},
				{
					Name: "salt", Type: "bytes32",
				},
			},
			"BridgeOperatorsBallot": []core.Type{
				{
					Name: "period", Type: "uint256",
				},
				{
					Name: "epoch", Type: "uint256",
				},
				{
					Name: "operators", Type: "address[]",
				},
			},
		},
		PrimaryType: "BridgeOperatorsBallot",
		Domain: core.TypedDataDomain{
			Name:    "GovernanceAdmin",
			Version: "2",
			Salt:    salt.Hex(),
		},
		Message: core.TypedDataMessage{
			"period":    math.NewHexOrDecimal256(period),
			"epoch":     math.NewHexOrDecimal256(epoch),
			"operators": bridgeOperators,
		},
	}

	return opts.SignTypedDataCallback(bridgeOperatorsBallotTypes)
}
