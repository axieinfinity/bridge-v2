package task

import (
	"crypto/ecdsa"
	ethGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	roninGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
	roninTrustedOrganization "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/trusted_organization"
	roninValidator "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/validator"
	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/models"
	"github.com/axieinfinity/bridge-v2/stores"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
	"time"
)

const Salt = "0xe3922a0bff7e80c6f7465bc1b150f6c95d9b9203f1731a09f86e759ea1eaa306"
const ErrSigAlreadySubmitted = "BOsGovernanceRelay: query for outdated period"

type task struct {
	util           utils.Utils
	task           *models.Task
	store          stores.BridgeStore
	validator      *ecdsa.PrivateKey
	client         bind.ContractBackend
	contracts      map[string]string
	chainId        *big.Int
	maxTry         int
	taskType       string
	listener       bridgeCore.Listener
	releaseTasksCh chan int
}

func newTask(listener bridgeCore.Listener, client bind.ContractBackend, store stores.BridgeStore, chainId *big.Int, contracts map[string]string, maxTry int, taskType string, releaseTasksCh chan int, util utils.Utils) *task {
	return &task{
		util:           util,
		task:           nil,
		store:          store,
		client:         client,
		contracts:      contracts,
		chainId:        chainId,
		maxTry:         maxTry,
		taskType:       taskType,
		listener:       listener,
		releaseTasksCh: releaseTasksCh,
	}
}

func (r *task) collectTask(t *models.Task) {
	log.Debug("Received new task", "hash", t.TransactionHash, "type", t.Type)
	if t.Type == r.taskType {
		r.task = t
	}
}

func (r *task) send() {
	log.Info("[task] Sending transaction", "type", r.taskType)
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
		log.Debug("[task] Transaction", "hash", transaction.Hash().Hex())
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
	// create caller
	roninGovernanceTransactor, err := roninGovernance.NewGovernanceTransactor(common.HexToAddress(r.contracts[GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	if err != nil {
		// append all success tasks into failed tasks
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	event, err := r.unpackBridgeOperatorSetUpdatedEvent(task)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	// otherwise add task to processingTasks to adjust after sending transaction
	processingTasks = append(processingTasks, task)

	var bridgeOperators []interface{}
	for _, address := range event.BridgeOperators {
		bridgeOperators = append(bridgeOperators, address.Hex())
	}
	opts := &signDataOpts{
		SignTypedDataCallback: func(typedData core.TypedData) (hexutil.Bytes, error) {
			return r.util.SignTypedData(typedData, r.listener.GetValidatorSign())
		},
	}
	signature, err := signBridgeOperatorsBallot(opts, event.Period.Int64(), bridgeOperators)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	signatureStruct := parseSignatureAsRsv(signature)
	log.Debug("[RoninTask][BridgeOperatorSetCallback] Prepared data", "r", common.Bytes2Hex(signatureStruct.R[:]), "s", common.Bytes2Hex(signatureStruct.S[:]), "v", signatureStruct.V, "period", event.Period.Int64(), "bridgeOperators", bridgeOperators)

	tx, err = r.util.SendContractTransaction(r.listener.GetValidatorSign(), r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return roninGovernanceTransactor.VoteBridgeOperatorsBySignatures(opts, event.Period, event.BridgeOperators, []roninGovernance.SignatureConsumerSignature{
			signatureStruct,
		})
	})
	if err != nil {
		log.Error("[RoninTask][BridgeOperatorSetCallback] Send transaction error", "err", err)
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	log.Debug("[RoninTask][BridgeOperatorSetCallback] Vote bridge operators", "hash", tx.Hash().Hex())

	doneTasks = append(doneTasks, task)
	return
}

func (r *task) relayBridgeOperators(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	log.Info("[RoninTask][BridgeOperatorsApprovedCallback] Processing task")
	// create caller
	roninTrustedCaller, err := roninTrustedOrganization.NewTrustedOrganizationCaller(common.HexToAddress(r.contracts[TRUSTED_ORGANIZATION_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	roninGovernanceCaller, err := roninGovernance.NewGovernanceCaller(common.HexToAddress(r.contracts[GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	ethGovernanceTransactor, err := ethGovernance.NewGovernanceTransactor(common.HexToAddress(r.contracts[ETH_GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	trustedOrgs, err := roninTrustedCaller.GetAllTrustedOrganizations(nil)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	var voters []common.Address
	for _, node := range trustedOrgs {
		voters = append(voters, node.BridgeVoter)
	}

	event, err := r.unpackBridgeOperatorsApprovedEvent(task)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	log.Debug("[RoninTask][BridgeOperatorsApprovedCallback] Unpacked event", "event", event)

	// otherwise add task to processingTasks to adjust after sending transaction
	processingTasks = append(processingTasks, task)

	signatures, err := roninGovernanceCaller.GetBridgeOperatorVotingSignatures(nil, event.Period, voters)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	log.Debug("[RoninTask][BridgeOperatorsApprovedCallback] Voting signatures", "voters", voters)

	var ethSignatures []ethGovernance.SignatureConsumerSignature
	for _, sig := range signatures {
		if sig.V > 0 {
			ethSignatures = append(ethSignatures, ethGovernance.SignatureConsumerSignature{
				V: sig.V,
				R: sig.R,
				S: sig.S,
			})
		}
	}

	tx, err = r.util.SendContractTransaction(r.listener.GetValidatorSign(), r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return ethGovernanceTransactor.RelayBridgeOperators(opts, event.Period, event.Operators, ethSignatures)
	})
	// Prevent retry submit signature if the signature was already submitted
	if err.Error() == ErrSigAlreadySubmitted {
		log.Debug("[RoninTask][BridgeOperatorsApprovedCallback] Bridge operators already submitted")
		doneTasks = append(doneTasks, task)
		return
	}
	if err != nil {
		log.Error("[RoninTask][BridgeOperatorsApprovedCallback] Send transaction error", "err", err)
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	log.Debug("[RoninTask][BridgeOperatorsApprovedCallback] Relay bridge operators", "hash", tx.Hash().Hex())

	doneTasks = append(doneTasks, task)
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
	SignTypedDataCallback func(typedData core.TypedData) (hexutil.Bytes, error)
}

func signBridgeOperatorsBallot(opts *signDataOpts, period int64, bridgeOperators interface{}) ([]byte, error) {
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
					Name: "operators", Type: "address[]",
				},
			},
		},
		PrimaryType: "BridgeOperatorsBallot",
		Domain: core.TypedDataDomain{
			Name:    "GovernanceAdmin",
			Version: "1",
			Salt:    Salt,
		},
		Message: core.TypedDataMessage{
			"period":    math.NewHexOrDecimal256(period),
			"operators": bridgeOperators,
		},
	}

	return opts.SignTypedDataCallback(bridgeOperatorsBallotTypes)
}
