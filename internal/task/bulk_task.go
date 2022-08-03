package task

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/stores"
	"github.com/ethereum/go-ethereum/signer/core"
	"math/big"
	"time"

	"github.com/axieinfinity/bridge-contracts/generated_contracts/ethereum/gateway"
	roninGateway "github.com/axieinfinity/bridge-contracts/generated_contracts/ronin/gateway"
	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type bulkTask struct {
	util           utils.Utils
	tasks          []*models.Task
	store          stores.ListenHandlerStore
	validator      *ecdsa.PrivateKey
	client         *ethclient.Client
	contracts      map[string]string
	chainId        *big.Int
	maxTry         int
	taskType       string
	listener       bridgeCore.Listener
	releaseTasksCh chan int
}

func newBulkTask(listener bridgeCore.Listener, client *ethclient.Client, store stores.ListenHandlerStore, chainId *big.Int, validator *ecdsa.PrivateKey, contracts map[string]string, ticker time.Duration, maxTry int, taskType string, releaseTasksCh chan int, util utils.Utils) *bulkTask {
	return &bulkTask{
		util:           util,
		tasks:          make([]*models.Task, 0),
		store:          store,
		validator:      validator,
		client:         client,
		contracts:      contracts,
		chainId:        chainId,
		maxTry:         maxTry,
		taskType:       taskType,
		listener:       listener,
		releaseTasksCh: releaseTasksCh,
	}
}

func (r *bulkTask) collectTask(t *models.Task) {
	if t.Type == r.taskType {
		r.tasks = append(r.tasks, t)
	}
}

func (r *bulkTask) send() {
	log.Info("[bulkTask] sending bulk", "type", r.taskType, "tasks", len(r.tasks))
	if len(r.tasks) == 0 {
		return
	}
	switch r.taskType {
	case DEPOSIT_TASK:
		r.sendBulkTransactions(r.sendDepositTransaction)
	case WITHDRAWAL_TASK:
		r.sendBulkTransactions(r.sendWithdrawalSignaturesTransaction)
	case ACK_WITHDREW_TASK:
		r.sendBulkTransactions(r.sendAckTransactions)
	}
}

func (r *bulkTask) sendBulkTransactions(sendTxs func(tasks []*models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction)) {
	start, end := 0, len(r.tasks)
	for start < end {
		var (
			txHash string
			next   int
		)
		if start+defaultLimitRecords < end {
			next = start + defaultLimitRecords
		} else {
			next = end
		}
		log.Info("[bulkTask][sendBulkTransactions] start sending txs", "start", start, "end", end, "type", r.taskType)
		doneTasks, processingTasks, failedTasks, transaction := sendTxs(r.tasks[start:next])

		if transaction != nil {
			go updateTasks(r.store, processingTasks, STATUS_PROCESSING, transaction.Hash().Hex(), time.Now().Unix(), r.releaseTasksCh)
			metrics.Pusher.IncrCounter(metrics.ProcessingTaskMetric, 1)
		}
		go updateTasks(r.store, doneTasks, STATUS_DONE, txHash, 0, r.releaseTasksCh)
		go updateTasks(r.store, failedTasks, STATUS_FAILED, txHash, 0, r.releaseTasksCh)
		metrics.Pusher.IncrCounter(metrics.SuccessTaskMetric, len(doneTasks))
		metrics.Pusher.IncrCounter(metrics.FailedTaskMetric, len(failedTasks))
		start = next
	}
}

func (r *bulkTask) sendDepositTransaction(tasks []*models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	var (
		receipts []roninGateway.TransferReceipt
	)
	// create caller
	caller, err := roninGateway.NewGatewayCaller(common.HexToAddress(r.contracts[GATEWAY_CONTRACT]), r.client)
	if err != nil {
		for _, t := range tasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, nil, failedTasks, nil
	}

	// create transactor
	transactor, err := roninGateway.NewGatewayTransactor(common.HexToAddress(r.contracts[GATEWAY_CONTRACT]), r.client)
	if err != nil {
		for _, t := range tasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, nil, failedTasks, nil
	}

	for _, t := range tasks {
		ok, receipt, err := r.validateDepositTask(caller, t)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		if receipt.Id != nil {
			// store receiptId to processed receipt
			if err := r.store.GetProcessedReceiptStore().Save(t.ID, receipt.Id.Int64()); err != nil {
				log.Error("[bulkTask][sendDepositTransaction] error while saving processed receipt", "err", err)
			}
		}

		// if deposit request is executed or voted (ok) then do nothing and add to doneTasks
		if ok {
			doneTasks = append(doneTasks, t)
			continue
		}

		// otherwise add task to processingTasks to adjust after sending transaction
		processingTasks = append(processingTasks, t)

		// append new receipt into receipts slice
		receipts = append(receipts, roninGateway.TransferReceipt{
			Id:   receipt.Id,
			Kind: receipt.Kind,
			Mainchain: roninGateway.TokenOwner{
				Addr:      receipt.Mainchain.Addr,
				TokenAddr: receipt.Mainchain.TokenAddr,
				ChainId:   receipt.Mainchain.ChainId,
			},
			Ronin: roninGateway.TokenOwner{
				Addr:      receipt.Ronin.Addr,
				TokenAddr: receipt.Ronin.TokenAddr,
				ChainId:   receipt.Ronin.ChainId,
			},
			Info: roninGateway.TokenInfo{
				Erc:      receipt.Info.Erc,
				Id:       receipt.Info.Id,
				Quantity: receipt.Info.Quantity,
			},
		})
	}
	metrics.Pusher.IncrCounter(metrics.DepositTaskMetric, len(tasks))

	if len(receipts) > 0 {
		tx, err = r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return transactor.TryBulkDepositFor(opts, receipts)
		})
		if err != nil {
			for _, t := range processingTasks {
				t.LastError = err.Error()
				failedTasks = append(failedTasks, t)
			}
			return doneTasks, nil, failedTasks, nil
		}
	}
	return
}

func (r *bulkTask) sendWithdrawalSignaturesTransaction(tasks []*models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	var (
		ids        []*big.Int
		signatures [][]byte
	)
	//create transactor
	transactor, err := roninGateway.NewGatewayTransactor(common.HexToAddress(r.contracts[GATEWAY_CONTRACT]), r.client)
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range tasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, nil, failedTasks, nil
	}
	// create caller
	caller, err := roninGateway.NewGatewayCaller(common.HexToAddress(r.contracts[GATEWAY_CONTRACT]), r.client)
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range tasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, nil, failedTasks, nil
	}
	for _, t := range tasks {
		result, receipt, err := r.validateWithdrawalTask(caller, t)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		if receipt.Id != nil {
			// store receiptId to processed receipt
			if err := r.store.GetProcessedReceiptStore().Save(t.ID, receipt.Id.Int64()); err != nil {
				log.Error("[bulkTask][sendWithdrawalSignaturesTransaction] error while saving processed receipt", "err", err)
			}
		}
		// if validated then do nothing and add to doneTasks
		if result {
			doneTasks = append(doneTasks, t)
			continue
		}
		// otherwise add to processingTasks
		sigs, err := r.signWithdrawalSignatures(receipt)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		processingTasks = append(processingTasks, t)
		signatures = append(signatures, sigs)
		ids = append(ids, receipt.Id)
	}
	metrics.Pusher.IncrCounter(metrics.WithdrawalTaskMetric, len(tasks))

	if len(ids) > 0 {
		tx, err = r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return transactor.BulkSubmitWithdrawalSignatures(opts, ids, signatures)
		})
		if err != nil {
			// append all success tasks into failed tasks
			for _, t := range processingTasks {
				t.LastError = err.Error()
				failedTasks = append(failedTasks, t)
			}
			return doneTasks, nil, failedTasks, nil
		}
	}
	return
}

func (r *bulkTask) sendAckTransactions(tasks []*models.Task) (doneTasks, processingTasks, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	var (
		ids []*big.Int
	)
	// create transactor
	transactor, err := roninGateway.NewGatewayTransactor(common.HexToAddress(r.contracts[GATEWAY_CONTRACT]), r.client)
	if err != nil {
		for _, t := range tasks {
			t.LastError = err.Error()
		}
		return nil, nil, tasks, nil
	}

	// create caller
	caller, err := roninGateway.NewGatewayCaller(common.HexToAddress(r.contracts[GATEWAY_CONTRACT]), r.client)
	if err != nil {
		for _, t := range tasks {
			t.LastError = err.Error()
		}
		return nil, nil, tasks, nil
	}

	// loop through tasks, check if they are qualified to send ack transaction or not
	for _, t := range tasks {
		result, id, err := r.validateAckWithdrawalTask(caller, t)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		if id != nil {
			// store receiptId to processed receipt
			if err := r.store.GetProcessedReceiptStore().Save(t.ID, id.Int64()); err != nil {
				log.Error("[bulkTask][sendAckTransactions] error while saving processed receipt", "err", err)
			}
		}

		// if validated then do nothing and add to doneTasks
		if result {
			doneTasks = append(doneTasks, t)
			continue
		}

		// otherwise add id to ids and add task to processingTasks
		ids = append(ids, id)
		processingTasks = append(processingTasks, t)
	}

	metrics.Pusher.IncrCounter(metrics.AckWithdrawalTaskMetric, len(tasks))
	if len(ids) > 0 {
		tx, err = r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return transactor.TryBulkAcknowledgeMainchainWithdrew(opts, ids)
		})
		if err != nil {
			// append all success tasks into failed tasks
			for _, t := range processingTasks {
				t.LastError = err.Error()
				failedTasks = append(failedTasks, t)
			}
			return doneTasks, nil, failedTasks, nil
		}
	}
	return
}

// ValidateDepositTask validates if:
// - current signer has been voted for a deposit request or not
// - deposit request has been executed or not
// also returns transfer receipt
func (r *bulkTask) validateDepositTask(caller *roninGateway.GatewayCaller, task *models.Task) (bool, gateway.TransferReceipt, error) {
	ethEvent := new(gateway.GatewayDepositRequested)
	ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
	if err != nil {
		return false, ethEvent.Receipt, err
	}

	data := common.Hex2Bytes(task.Data)
	if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "DepositRequested", data); err != nil {
		return false, ethEvent.Receipt, err
	}

	// check deposit vote
	result, err := caller.DepositVote(nil, ethEvent.Receipt.Mainchain.ChainId, ethEvent.Receipt.Id)
	if err != nil {
		return false, ethEvent.Receipt, err
	}
	if result.Status == VoteStatusExecuted {
		return true, ethEvent.Receipt, nil
	}

	// check if current validator has been voted for this deposit or not
	voted, err := caller.DepositVoted(nil, ethEvent.Receipt.Mainchain.ChainId, ethEvent.Receipt.Id, crypto.PubkeyToAddress(r.validator.PublicKey))
	if err != nil {
		return false, ethEvent.Receipt, err
	}
	return voted, ethEvent.Receipt, nil
}

// ValidateAckWithdrawalTask validates if:
// - signer has been voted for a withdrawal or not
// - withdrawal request is executed or not
// returns true if withdraw is executed or voted
// also returns receipt id
func (r *bulkTask) validateAckWithdrawalTask(caller *roninGateway.GatewayCaller, task *models.Task) (bool, *big.Int, error) {
	// Unpack event data
	ethEvent := new(gateway.GatewayWithdrew)
	ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
	if err != nil {
		return false, nil, err
	}

	if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "Withdrew", common.Hex2Bytes(task.Data)); err != nil {
		return false, nil, err
	}

	// check if withdraw has been executed or not
	result, err := caller.MainchainWithdrew(nil, ethEvent.Receipt.Id)
	if err != nil {
		return false, nil, err
	}
	if result {
		return true, ethEvent.Receipt.Id, nil
	}

	// check if withdrew has been voted or not
	voted, err := caller.MainchainWithdrewVoted(nil, ethEvent.Receipt.Id, crypto.PubkeyToAddress(r.validator.PublicKey))
	if err != nil {
		return false, nil, err
	}
	return voted, ethEvent.Receipt.Id, nil
}

// ValidateWithdrawalTask validates if:
// - Withdrawal request is executed or not
// returns true if it is executed
// also returns transfer receipt
func (r *bulkTask) validateWithdrawalTask(caller *roninGateway.GatewayCaller, task *models.Task) (bool, roninGateway.TransferReceipt, error) {
	// Unpack event from data
	ronEvent := new(roninGateway.GatewayMainchainWithdrew)
	ronGatewayAbi, err := roninGateway.GatewayMetaData.GetAbi()
	if err != nil {
		return false, ronEvent.Receipt, err
	}
	if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "MainchainWithdrew", common.Hex2Bytes(task.Data)); err != nil {
		return false, ronEvent.Receipt, err
	}
	receipt := ronEvent.Receipt

	// try getting withdrawal data from database by receipt.id, do nothing if withdrawal is found
	withdrawal, _ := r.store.GetWithdrawalStore().GetWithdrawalById(receipt.Id.Int64())
	if withdrawal != nil && withdrawal.ID > 0 {
		return true, ronEvent.Receipt, nil
	}

	// check if withdraw has been executed or not
	result, err := caller.MainchainWithdrew(nil, receipt.Id)
	if err != nil {
		return false, ronEvent.Receipt, err
	}
	return result, ronEvent.Receipt, nil
}

func updateTasks(store stores.ListenHandlerStore, tasks []*models.Task, status, txHash string, timestamp int64, releaseTasksCh chan int) {
	// update tasks with given status
	// note: if task.retries < 10 then retries++ and status still be processing
	for _, t := range tasks {
		if timestamp > 0 {
			t.TxCreatedAt = timestamp
		}
		if status == STATUS_FAILED {
			if t.Retries+1 >= 10 {
				t.Status = status
			} else {
				t.Retries += 1
			}
		} else {
			t.Status = status
			t.TransactionHash = txHash
		}
		if err := store.GetTaskStore().Update(t); err != nil {
			log.Error("error while update task", "id", t.ID, "err", err)
		}
		releaseTasksCh <- t.ID
	}
}

func (r *bulkTask) signWithdrawalSignatures(receipt roninGateway.TransferReceipt) (hexutil.Bytes, error) {
	typedData := core.TypedData{
		Types: core.Types{
			"EIP712Domain": []core.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"Receipt": []core.Type{
				{Name: "id", Type: "uint256"},
				{Name: "kind", Type: "uint8"},
				{Name: "mainchain", Type: "TokenOwner"},
				{Name: "ronin", Type: "TokenOwner"},
				{Name: "info", Type: "TokenInfo"},
			},
			"TokenOwner": []core.Type{
				{Name: "addr", Type: "address"},
				{Name: "tokenAddr", Type: "address"},
				{Name: "chainId", Type: "uint256"},
			},
			"TokenInfo": []core.Type{
				{Name: "erc", Type: "uint8"},
				{Name: "id", Type: "uint256"},
				{Name: "quantity", Type: "uint256"},
			},
		},
		Domain: core.TypedDataDomain{
			Name:              "MainchainGatewayV2",
			Version:           "2",
			ChainId:           math.NewHexOrDecimal256(receipt.Mainchain.ChainId.Int64()),
			VerifyingContract: r.contracts[ETH_GATEWAY_CONTRACT],
		},
		PrimaryType: "Receipt",
		Message: core.TypedDataMessage{
			"id":   receipt.Id.String(),
			"kind": fmt.Sprintf("%d", receipt.Kind),
			"mainchain": core.TypedDataMessage{
				"addr":      receipt.Mainchain.Addr.Hex(),
				"tokenAddr": receipt.Mainchain.TokenAddr.Hex(),
				"chainId":   receipt.Mainchain.ChainId.String(),
			},
			"ronin": core.TypedDataMessage{
				"addr":      receipt.Ronin.Addr.Hex(),
				"tokenAddr": receipt.Ronin.TokenAddr.Hex(),
				"chainId":   receipt.Ronin.ChainId.String(),
			},
			"info": core.TypedDataMessage{
				"erc":      fmt.Sprintf("%d", receipt.Info.Erc),
				"id":       receipt.Info.Id.String(),
				"quantity": receipt.Info.Quantity.String(),
			},
		},
	}
	return r.util.SignTypedData(typedData, r.validator)
}
