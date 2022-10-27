package task

import (
	"crypto/ecdsa"
	ethGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/governance"
	roninGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
	roninGovernance "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/governance"
	roninTrustedOrg "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/trusted_org"
	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/models"
	"github.com/axieinfinity/bridge-v2/stores"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
	"time"
)

type task struct {
	util           utils.Utils
	task           *models.Task
	store          stores.BridgeStore
	validator      *ecdsa.PrivateKey
	client         *ethclient.Client
	contracts      map[string]string
	chainId        *big.Int
	maxTry         int
	taskType       string
	listener       bridgeCore.Listener
	releaseTasksCh chan int
}

func newTask(listener bridgeCore.Listener, client *ethclient.Client, store stores.BridgeStore, chainId *big.Int, contracts map[string]string, maxTry int, taskType string, releaseTasksCh chan int, util utils.Utils) *task {
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
	if t.Type == r.taskType {
		r.task = t
	}
}

func (r *task) send() {
	log.Info("[bulk] sending transaction", "type", r.taskType)
	switch r.taskType {
	case VOTE_BRIDGE_OPERATORS_TASK:
		r.sendTransaction(r.voteBridgeOperatorsBySignature)
	case RELAY_BRIDGE_OPERATORS_TASK:
		r.sendTransaction(r.relayBridgeOperators)
	}
}

func (r *task) sendTransaction(sendTx func(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction)) {
	var txHash string

	log.Info("[task][sendTransaction] start sending tx", "type", r.taskType)
	doneTasks, processingTasks, failedTasks, transaction := sendTx(r.task)

	if transaction != nil {
		go updateTasks(r.store, processingTasks, STATUS_PROCESSING, transaction.Hash().Hex(), time.Now().Unix(), r.releaseTasksCh)
		metrics.Pusher.IncrGauge(metrics.ProcessingTaskMetric, len(processingTasks))
	}
	go updateTasks(r.store, doneTasks, STATUS_DONE, txHash, 0, r.releaseTasksCh)
	go updateTasks(r.store, failedTasks, STATUS_FAILED, txHash, 0, r.releaseTasksCh)
	metrics.Pusher.IncrCounter(metrics.SuccessTaskMetric, len(doneTasks))
	metrics.Pusher.IncrCounter(metrics.FailedTaskMetric, len(failedTasks))
}

func (r *task) voteBridgeOperatorsBySignature(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	// create caller
	transactor, err := roninGovernance.NewGatewayTransactor(common.HexToAddress(r.contracts[GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	// create caller
	if err != nil {
		// append all success tasks into failed tasks
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	event, err := r.unpackBridgeOperatorsUpdatedEvent(task)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	// otherwise add task to processingTasks to adjust after sending transaction
	processingTasks = append(processingTasks, task)

	signatures, err := r.signBridgeOperatorsBallot(event.Period, event.BridgeOperators)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	tx, err = r.util.SendContractTransaction(r.listener.GetValidatorSign(), r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return transactor.VoteBridgeOperatorsBySignatures(opts, event.Period, event.BridgeOperators, parseSignatureAsRsv(signatures))
	})

	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	doneTasks = append(doneTasks, task)
	return
}

func (r *task) relayBridgeOperators(task *models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	// create caller
	roninTrustedCaller, err := roninTrustedOrg.NewGatewayCaller(common.HexToAddress(r.contracts[TRUSTED_ORGS_CONTRACT]), r.client)
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
	log.Debug("[RoninListener][BridgeOperatorsApprovedCallback] Trusted organization", "trustedOrgs", trustedOrgs)

	var voters []common.Address
	for _, node := range trustedOrgs {
		voters = append(voters, node.BridgeVoter)
	}

	roninGovernanceCaller, err := roninGovernance.NewGatewayCaller(common.HexToAddress(r.contracts[GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	ethGovernanceTransactor, err := ethGovernance.NewGatewayTransactor(common.HexToAddress(r.contracts[ETH_GOVERNANCE_CONTRACT]), r.client)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	event, err := r.unpackBridgeOperatorsApprovedEvent(task)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}

	// otherwise add task to processingTasks to adjust after sending transaction
	processingTasks = append(processingTasks, task)

	signatures, err := roninGovernanceCaller.GetBridgeOperatorVotingSignatures(nil, event.Period, voters)
	if err != nil {
		task.LastError = err.Error()
		failedTasks = append(failedTasks, task)
		return nil, nil, failedTasks, nil
	}
	log.Debug("[RoninListener][BridgeOperatorsApprovedCallback] Voting signatures", "signatures", signatures)

	tx, err = r.util.SendContractTransaction(r.listener.GetValidatorSign(), r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return ethGovernanceTransactor.RelayBridgeOperators(opts, event.Period, event.BridgeOperators, signatures)
	})

	doneTasks = append(doneTasks, task)
	return
}

func (r *task) unpackBridgeOperatorsUpdatedEvent(task *models.Task) (*roninGateway.GatewayBridgeOperatorsUpdated, error) {
	ronEvent := new(roninGateway.GatewayBridgeOperatorsUpdated)
	ronGatewayAbi, err := roninGateway.GatewayMetaData.GetAbi()
	if err != nil {
		return ronEvent, err
	}

	if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "BridgeOperatorsUpdated", common.Hex2Bytes(task.Data)); err != nil {
		return ronEvent, err
	}

	return ronEvent, nil
}

func (r *task) unpackBridgeOperatorsApprovedEvent(task *models.Task) (*roninGateway.GatewayBridgeOperatorsApproved, error) {
	ronEvent := new(roninGateway.GatewayBridgeOperatorsApproved)
	ronGatewayAbi, err := roninGateway.GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "BridgeOperatorsApproved", common.Hex2Bytes(task.Data)); err != nil {
		return nil, err
	}

	return ronEvent, nil
}

func (r *task) signBridgeOperatorsBallot(period *big.Int, bridgeOperators []common.Address) ([]byte, error) {
	bridgeOperatorsBallotTypes := core.TypedData{
		Types: core.Types{
			"BridgeOperatorsBallot": []core.Type{
				{
					Name: "period", Type: "uint256",
				},
				{
					Name: "operators", Type: "address[]",
				},
			},
		},
		PrimaryType: "RoninGatewayV2",
		Domain: core.TypedDataDomain{
			Name:              "RoninGatewayV2",
			Version:           "2",
			ChainId:           math.NewHexOrDecimal256(r.chainId.Int64()),
			VerifyingContract: r.contracts[GATEWAY_CONTRACT],
		},
		Message: core.TypedDataMessage{
			"period":    period,
			"operators": bridgeOperators,
		},
	}

	signature, err := r.util.SignTypedData(bridgeOperatorsBallotTypes, r.listener.GetValidatorSign())
	if err != nil {
		return nil, err
	}

	return signature, nil
}
