package task

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/axieinfinity/bridge-v2/generated_contracts/ethereum/gateway"
	gateway2 "github.com/axieinfinity/bridge-v2/generated_contracts/ronin/gateway"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"sync"
	"time"
)

const (
	defaultLimitRecords = 50
	defaultMaxTry       = 5
	defaultReceiptCheck = 30
)

var defaultTaskInterval = 1 * time.Second

type RoninTask struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	util utils.IUtils

	listener types.IListener
	store    types.IMainStore

	taskInterval    time.Duration
	txCheckInterval time.Duration
	secret          *types.Secret

	client    *ethclient.Client
	contracts map[string]string

	validator *ecdsa.PrivateKey
	relayer   *ecdsa.PrivateKey

	limitQuery int
	chainId    *big.Int
}

func NewRoninTask(listener types.IListener, util utils.IUtils) (*RoninTask, error) {
	config := listener.Config()
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		return nil, err
	}
	chainId, err := listener.GetChainID()
	if err != nil {
		return nil, err
	}
	newCtx, cancelFunc := context.WithCancel(listener.Context())
	task := &RoninTask{
		ctx:             newCtx,
		cancelFunc:      cancelFunc,
		listener:        listener,
		store:           listener.GetStore(),
		taskInterval:    defaultTaskInterval,
		txCheckInterval: defaultTaskInterval,
		secret:          config.Secret,
		contracts:       config.Contracts,
		client:          client,
		chainId:         chainId,
		util:            util,
		validator:       listener.(types.IEthListener).GetValidatorKey(),
		relayer:         listener.(types.IEthListener).GetRelayerKey(),
		limitQuery:      defaultLimitRecords,
	}
	if config.TaskInterval > 0 {
		task.taskInterval = config.TaskInterval * time.Second
	}
	if config.TransactionCheckPeriod > 0 {
		task.txCheckInterval = config.TransactionCheckPeriod * time.Second
	}
	if config.MaxTasksQuery > 0 {
		task.limitQuery = config.MaxTasksQuery
	}
	return task, nil
}

func (r *RoninTask) Start() {
	taskTicker := time.NewTicker(r.taskInterval)
	processingTicker := time.NewTicker(time.Duration(defaultReceiptCheck) * time.Second)
	for {
		select {
		case <-taskTicker.C:
			go func() {
				if err := r.processPending(); err != nil {
					log.Error("[RoninTask] error while process tasks", "err", err)
				}
			}()
		case <-processingTicker.C:
			go func() {
				if err := r.checkProcessingTasks(); err != nil {
					log.Error("[RoninTask] error while checking processing tasks", "err", err)
				}
			}()
		case <-r.ctx.Done():
			r.client.Close()
			r.Close()
			return
		}
	}
}

func (r *RoninTask) SetLimitQuery(limit int) {
	r.limitQuery = limit
}

func (r *RoninTask) Close() {
	r.cancelFunc()
}

func (r *RoninTask) GetListener() types.IListener {
	return r.listener
}

func (r *RoninTask) processPending() error {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[RoninTask][processPending]recover from panic", "err", err)
		}
	}()
	tasks, err := r.store.GetTaskStore().GetTasks(hexutil.EncodeBig(r.chainId), types.STATUS_PENDING, r.limitQuery, 10)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return nil
	}

	bulkDepositTask := NewBulkTask(r.listener, r.client, r.store, r.chainId, r.validator, r.contracts, r.txCheckInterval, defaultMaxTry, types.DEPOSIT_TASK, r.util)
	bulkSubmitWithdrawalSignaturesTask := NewBulkTask(r.listener, r.client, r.store, r.chainId, r.validator, r.contracts, r.txCheckInterval, defaultMaxTry, types.WITHDRAWAL_TASK, r.util)
	ackWithdrewTasks := NewBulkTask(r.listener, r.client, r.store, r.chainId, r.validator, r.contracts, r.txCheckInterval, defaultMaxTry, types.ACK_WITHDREW_TASK, r.util)
	for _, task := range tasks {
		// collect tasks for bulk deposits
		bulkDepositTask.collectTask(task)

		// collect tasks for bulk withdrawal signature
		bulkSubmitWithdrawalSignaturesTask.collectTask(task)

		// collect tasks for acknowledge withdrawal
		ackWithdrewTasks.collectTask(task)
	}
	bulkDepositTask.send()
	bulkSubmitWithdrawalSignaturesTask.send()
	ackWithdrewTasks.send()
	return nil
}

func (r *RoninTask) checkProcessingTasks() error {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] recover from panic", "err", err)
		}
	}()
	tasks, err := r.store.GetTaskStore().GetTasks(hexutil.EncodeBig(r.chainId), types.STATUS_PROCESSING, r.limitQuery, 2)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return nil
	}

	failedTasksMap := sync.Map{}
	successTasks := sync.Map{}
	processedTx := make(map[string][]*models.Task)

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, t := range tasks {
		if _, ok := processedTx[t.TransactionHash]; ok {
			wg.Done()
			processedTx[t.TransactionHash] = append(processedTx[t.TransactionHash], t)
			continue
		}
		processedTx[t.TransactionHash] = append(processedTx[t.TransactionHash], t)

		go func(task *models.Task) {
			defer wg.Done()
			// check transaction receipt status
			log.Info("[RoninTask][checkProcessingTasks] Start checking transaction status", "tx", task.TransactionHash)
			receipt, err := r.listener.GetReceipt(common.HexToHash(task.TransactionHash))
			if err != nil {
				failedTasksMap.Store(task.TransactionHash, struct{}{})
				return
			}
			if receipt != nil {
				// start confirmation step
				// start 3 times confirmation
				for i := 0; i < int(r.listener.Config().SafeBlockRange); i++ {
					time.Sleep(r.listener.Config().TransactionCheckPeriod * time.Second)
					confirmedReceipt, _ := r.listener.GetReceipt(common.HexToHash(task.TransactionHash))
					if confirmedReceipt == nil { // receipt is not found, then reorg may happen then break and retry again
						failedTasksMap.Store(task.TransactionHash, struct{}{})
						return
					}
				}
				// if receipt status is failed then add tx to failed tasks map
				if receipt.Status == 1 {
					successTasks.Store(task.TransactionHash, receipt.Status)
					return
				}
			}
			failedTasksMap.Store(task.TransactionHash, struct{}{})
		}(t)
	}
	wg.Wait()

	// collect successTasks
	var (
		successTxs []string
		failedTxs  []string
	)

	successTasks.Range(func(key interface{}, value interface{}) bool {
		if value.(uint64) == 1 {
			successTxs = append(successTxs, key.(string))
		} else {
			failedTxs = append(failedTxs, key.(string))
		}
		return true
	})

	// collect droppedTasks, retryTasks
	var (
		droppedTaskIds []int
		retryTaskIds   []int
	)
	failedTasksMap.Range(func(key interface{}, value interface{}) bool {
		failedTasks := processedTx[key.(string)]
		for _, task := range failedTasks {
			if task.Retries+1 >= 10 {
				droppedTaskIds = append(droppedTaskIds, task.ID)
			} else {
				retryTaskIds = append(retryTaskIds, task.ID)
			}
		}
		return true
	})

	// update success tasks with transaction's status = 1 (success)
	if len(successTxs) > 0 {
		if err = r.store.GetTaskStore().UpdateTasksWithTransactionHash(successTxs, 1, types.STATUS_DONE); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while update tasks with success transactions", "err", err)
		}
	}

	// update success tasks with transaction's status = 0 (failed)
	if len(failedTxs) > 0 {
		if err = r.store.GetTaskStore().UpdateTasksWithTransactionHash(failedTxs, 0, types.STATUS_DONE); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while update tasks with failed transactions", "err", err)
		}
	}

	// update failed tasks
	if len(droppedTaskIds) > 0 {
		if err = r.store.GetTaskStore().UpdateTaskWithIds(droppedTaskIds, 0, types.STATUS_FAILED); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while update tasks with success transactions", "err", err)
		}
	}

	// update retries tasks
	if len(retryTaskIds) > 0 {
		if err = r.store.GetTaskStore().IncrementRetries(retryTaskIds); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while increment retries", "err", err)
		}
	}

	return nil
}

type BulkTask struct {
	util      utils.IUtils
	tasks     []*models.Task
	store     types.IMainStore
	validator *ecdsa.PrivateKey
	client    *ethclient.Client
	contracts map[string]string
	chainId   *big.Int
	ticker    time.Duration
	maxTry    int
	taskType  string
	listener  types.IListener
}

func NewBulkTask(listener types.IListener, client *ethclient.Client, store types.IMainStore, chainId *big.Int, validator *ecdsa.PrivateKey, contracts map[string]string, ticker time.Duration, maxTry int, taskType string, util utils.IUtils) *BulkTask {
	return &BulkTask{
		util:      util,
		tasks:     make([]*models.Task, 0),
		store:     store,
		validator: validator,
		client:    client,
		contracts: contracts,
		chainId:   chainId,
		ticker:    ticker,
		maxTry:    maxTry,
		taskType:  taskType,
		listener:  listener,
	}
}

func (r *BulkTask) collectTask(t *models.Task) {
	if t.Type == r.taskType {
		r.tasks = append(r.tasks, t)
	}
}

func (r *BulkTask) send() {
	log.Info("[BulkTask] sending bulk", "type", r.taskType, "tasks", len(r.tasks))
	if len(r.tasks) == 0 {
		return
	}
	switch r.taskType {
	case types.DEPOSIT_TASK:
		r.sendBulkTransactions(r.sendDepositTransaction)
	case types.WITHDRAWAL_TASK:
		r.sendBulkTransactions(r.sendWithdrawalSignaturesTransaction)
	case types.ACK_WITHDREW_TASK:
		r.sendBulkTransactions(r.SendAckTransactions)
	}
}

func (r *BulkTask) sendBulkTransactions(sendTxs func(tasks []*models.Task) (successTasks []*models.Task, failedTasks []*models.Task, tx *ethtypes.Transaction)) {
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
		log.Info("[BulkTask][sendBulkTransactions] start sending txs", "start", start, "end", end, "type", r.taskType)
		successTasks, failedTasks, transaction := sendTxs(r.tasks[start:next])

		if transaction != nil {
			txHash = transaction.Hash().Hex()
			go updateTasks(r.store, successTasks, types.STATUS_PROCESSING, txHash)
		}
		go updateTasks(r.store, failedTasks, types.STATUS_FAILED, txHash)
		start = next
	}

}

func (r *BulkTask) sendDepositTransaction(tasks []*models.Task) (successTasks []*models.Task, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	var (
		receipts []gateway2.TransferReceipt
	)
	// create caller
	caller, err := gateway2.NewGatewayCaller(common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.client)
	if err != nil {
		return
	}

	for _, t := range tasks {
		ethEvent := new(gateway.GatewayDepositRequested)
		ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		data := common.Hex2Bytes(t.Data)
		if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "DepositRequested", data); err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		// append task into success tasks
		successTasks = append(successTasks, t)

		// check deposit vote
		result, err := caller.DepositVote(nil, ethEvent.Receipt.Mainchain.ChainId, ethEvent.Receipt.Id)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		if result.Status == types.VoteStatusExecuted {
			continue
		}

		// check if current validator has been voted for this deposit or not
		voted, err := caller.DepositVoted(nil, ethEvent.Receipt.Mainchain.ChainId, ethEvent.Receipt.Id, crypto.PubkeyToAddress(r.validator.PublicKey))
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		if voted {
			continue
		}

		// append new receipt into receipts slice
		receipts = append(receipts, gateway2.TransferReceipt{
			Id:   ethEvent.Receipt.Id,
			Kind: ethEvent.Receipt.Kind,
			Mainchain: gateway2.TokenOwner{
				Addr:      ethEvent.Receipt.Mainchain.Addr,
				TokenAddr: ethEvent.Receipt.Mainchain.TokenAddr,
				ChainId:   ethEvent.Receipt.Mainchain.ChainId,
			},
			Ronin: gateway2.TokenOwner{
				Addr:      ethEvent.Receipt.Ronin.Addr,
				TokenAddr: ethEvent.Receipt.Ronin.TokenAddr,
				ChainId:   ethEvent.Receipt.Ronin.ChainId,
			},
			Info: gateway2.TokenInfo{
				Erc:      ethEvent.Receipt.Info.Erc,
				Id:       ethEvent.Receipt.Info.Id,
				Quantity: ethEvent.Receipt.Info.Quantity,
			},
		})
	}
	// send tx with above receipts
	c, err := gateway2.NewGatewayTransactor(common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.client)
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks, nil
	}
	tx, err = r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		return c.TryBulkDepositFor(opts, receipts)
	})
	if err != nil {
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks, nil
	}
	return
}

func (r *BulkTask) sendWithdrawalSignaturesTransaction(tasks []*models.Task) (successTasks []*models.Task, failedTasks []*models.Task, tx *ethtypes.Transaction) {

	var (
		ids        []*big.Int
		signatures [][]byte
	)

	for _, t := range tasks {
		// Unpack event from data
		ronEvent := new(gateway2.GatewayMainchainWithdrew)
		ronGatewayAbi, err := gateway2.GatewayMetaData.GetAbi()
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		if err = ronGatewayAbi.UnpackIntoInterface(ronEvent, "MainchainWithdrew", common.Hex2Bytes(t.Data)); err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		receipt := ronEvent.Receipt

		// try getting withdrawal data from database by receipt.id, do nothing if withdrawal is found
		withdrawal, _ := r.store.GetWithdrawalStore().GetWithdrawalById(receipt.Id.Int64())
		if withdrawal != nil && withdrawal.ID > 0 {
			continue
		}

		// try checking on smart contract
		// create caller
		caller, err := gateway2.NewGatewayCaller(common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.client)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		result, err := caller.MainchainWithdrew(nil, receipt.Id)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		if !result {
			typedData := apitypes.TypedData{
				Types: apitypes.Types{
					"EIP712Domain": []apitypes.Type{
						{Name: "name", Type: "string"},
						{Name: "version", Type: "string"},
						{Name: "chainId", Type: "uint256"},
						{Name: "verifyingContract", Type: "address"},
					},
					"Receipt": []apitypes.Type{
						{Name: "id", Type: "uint256"},
						{Name: "kind", Type: "uint8"},
						{Name: "mainchain", Type: "TokenOwner"},
						{Name: "ronin", Type: "TokenOwner"},
						{Name: "info", Type: "TokenInfo"},
					},
					"TokenOwner": []apitypes.Type{
						{Name: "addr", Type: "address"},
						{Name: "tokenAddr", Type: "address"},
						{Name: "chainId", Type: "uint256"},
					},
					"TokenInfo": []apitypes.Type{
						{Name: "erc", Type: "uint8"},
						{Name: "id", Type: "uint256"},
						{Name: "quantity", Type: "uint256"},
					},
				},
				Domain: apitypes.TypedDataDomain{
					Name:              "MainchainGatewayV2",
					Version:           "2",
					ChainId:           math.NewHexOrDecimal256(receipt.Mainchain.ChainId.Int64()),
					VerifyingContract: r.contracts[types.ETH_GATEWAY_CONTRACT],
				},
				PrimaryType: "Receipt",
				Message: apitypes.TypedDataMessage{
					"id":   receipt.Id.String(),
					"kind": fmt.Sprintf("%d", receipt.Kind),
					"mainchain": apitypes.TypedDataMessage{
						"addr":      receipt.Mainchain.Addr.Hex(),
						"tokenAddr": receipt.Mainchain.TokenAddr.Hex(),
						"chainId":   receipt.Mainchain.ChainId.String(),
					},
					"ronin": apitypes.TypedDataMessage{
						"addr":      receipt.Ronin.Addr.Hex(),
						"tokenAddr": receipt.Ronin.TokenAddr.Hex(),
						"chainId":   receipt.Ronin.ChainId.String(),
					},
					"info": apitypes.TypedDataMessage{
						"erc":      fmt.Sprintf("%d", receipt.Info.Erc),
						"id":       receipt.Info.Id.String(),
						"quantity": receipt.Info.Quantity.String(),
					},
				},
			}

			sigs, err := r.util.SignTypedData(typedData, r.validator)
			if err != nil {
				t.LastError = err.Error()
				failedTasks = append(failedTasks, t)
				continue
			}
			signatures = append(signatures, sigs)
			ids = append(ids, receipt.Id)
		}
		successTasks = append(successTasks, t)
	}
	c, err := gateway2.NewGatewayTransactor(common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.client)
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks, nil
	}
	if len(ids) > 0 {
		tx, err = r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.BulkSubmitWithdrawalSignatures(opts, ids, signatures)
		})
		if err != nil {
			// append all success tasks into failed tasks
			for _, t := range successTasks {
				t.LastError = err.Error()
				failedTasks = append(failedTasks, t)
			}
			return nil, failedTasks, nil
		}
	}
	return
}

func (r *BulkTask) SendAckTransactions(tasks []*models.Task) (successTasks []*models.Task, failedTasks []*models.Task, tx *ethtypes.Transaction) {
	var ids []*big.Int
	c, err := gateway2.NewGatewayTransactor(common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.client)
	if err != nil {
		for _, t := range tasks {
			t.LastError = err.Error()
		}
		return nil, tasks, nil
	}
	for _, t := range tasks {
		// Unpack event data
		ethEvent := new(gateway.GatewayWithdrew)
		ethGatewayAbi, err := gateway.GatewayMetaData.GetAbi()
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		if err = ethGatewayAbi.UnpackIntoInterface(ethEvent, "Withdrew", common.Hex2Bytes(t.Data)); err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		// try checking on smart contract
		// create caller
		caller, err := gateway2.NewGatewayCaller(common.HexToAddress(r.contracts[types.GATEWAY_CONTRACT]), r.client)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}
		result, err := caller.MainchainWithdrew(nil, ethEvent.Receipt.Id)
		if err != nil {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
			continue
		}

		if !result {
			ids = append(ids, ethEvent.Receipt.Id)
		}
		successTasks = append(successTasks, t)
	}
	tx, err = r.util.SendContractTransaction(r.validator, r.chainId, func(opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
		if ids != nil {
			return c.TryBulkAcknowledgeMainchainWithdrew(opts, ids)
		}
		return nil, errors.New("empty withdraw ids list")
	})
	if err != nil {
		// append all success tasks into failed tasks
		for _, t := range successTasks {
			t.LastError = err.Error()
			failedTasks = append(failedTasks, t)
		}
		return nil, failedTasks, nil
	}
	return
}

func updateTasks(store types.IMainStore, tasks []*models.Task, status, txHash string) {
	// update tasks with given status
	// note: if task.retries < 10 then retries++ and status still be processing
	for _, t := range tasks {
		if status == types.STATUS_FAILED {
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
	}
}
