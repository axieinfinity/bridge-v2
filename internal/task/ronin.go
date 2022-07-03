package task

import (
	"context"
	"crypto/ecdsa"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sync"
	"time"
)

const (
	defaultLimitRecords = 50
	defaultMaxTry       = 5
	defaultReceiptCheck = 50
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
	tasks, err := r.store.GetTaskStore().GetTasks(hexutil.EncodeBig(r.chainId), types.STATUS_PENDING, r.limitQuery, 10, 0)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return nil
	}

	bulkDepositTask := newBulkTask(r.listener, r.client, r.store, r.chainId, r.validator, r.contracts, r.txCheckInterval, defaultMaxTry, types.DEPOSIT_TASK, r.util)
	bulkSubmitWithdrawalSignaturesTask := newBulkTask(r.listener, r.client, r.store, r.chainId, r.validator, r.contracts, r.txCheckInterval, defaultMaxTry, types.WITHDRAWAL_TASK, r.util)
	ackWithdrewTasks := newBulkTask(r.listener, r.client, r.store, r.chainId, r.validator, r.contracts, r.txCheckInterval, defaultMaxTry, types.ACK_WITHDREW_TASK, r.util)
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
	before := time.Now().Unix() - int64(r.listener.Config().SafeBlockRange*uint64(r.listener.Config().LoadInterval.Seconds()))
	tasks, err := r.store.GetTaskStore().GetTasks(hexutil.EncodeBig(r.chainId), types.STATUS_PROCESSING, r.limitQuery, 2, before)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return nil
	}

	failedTasks := sync.Map{}
	successTasks := sync.Map{}
	processedTx := make(map[string]struct{})

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, t := range tasks {
		if _, ok := processedTx[t.TransactionHash]; ok {
			wg.Done()
			continue
		}
		processedTx[t.TransactionHash] = struct{}{}

		go func(task *models.Task) {
			defer wg.Done()

			// check transaction receipt status
			log.Info("[RoninTask][checkProcessingTasks] Start checking transaction status", "tx", task.TransactionHash)
			receipt, err := r.listener.GetReceipt(common.HexToHash(task.TransactionHash))
			if err != nil || receipt == nil {
				failedTasks.Store(task, struct{}{})
				return
			}

			// add task and transaction's status into successTasks
			successTasks.Store(task, receipt.Status)
		}(t)
	}
	wg.Wait()

	var (
		successTxs     []string
		failedTxs      []string
		resetToPending []string
	)

	// loop through successTasks, if receipt is failed then reset to pending and retry if retry is not reached to 10
	successTasks.Range(func(key interface{}, value interface{}) bool {
		task := key.(*models.Task)
		if value.(uint64) == 1 {
			successTxs = append(successTxs, task.TransactionHash)
		} else {
			if task.Retries+1 >= 10 {
				// append to failedTxs and update all tasks with this transactionHash to failed
				failedTxs = append(failedTxs, task.TransactionHash)
			} else {
				// append to resetToPending and update all tasks with this transactionHash to pending
				resetToPending = append(resetToPending, task.TransactionHash)
			}
		}
		return true
	})

	// loop through failedTasksMap, reset to pending and retry if retry is not reached to 10
	failedTasks.Range(func(key interface{}, value interface{}) bool {
		task := key.(*models.Task)
		if task.Retries+1 >= 10 {
			// append to failedTxs and update all tasks with this transactionHash to failed
			failedTxs = append(failedTxs, task.TransactionHash)
		} else {
			// append to resetToPending and update all tasks with this transactionHash to pending
			resetToPending = append(resetToPending, task.TransactionHash)
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
		if err = r.store.GetTaskStore().UpdateTasksWithTransactionHash(failedTxs, 0, types.STATUS_FAILED); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while update tasks with failed transactions", "err", err)
		}
	}

	// update all tasks contain failed txs to pending
	if len(resetToPending) > 0 {
		log.Info("[RoninTask][checkProcessingTasks] reset tasks to pending", "transactionHash", resetToPending)
		if err = r.store.GetTaskStore().ResetTo(resetToPending, types.STATUS_PENDING); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while reset tasks to pending", "err", err)
		}
	}

	return nil
}
