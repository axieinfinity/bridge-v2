package task

import (
	"context"
	"math/big"
	"sync"
	"time"

	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/models"
	"github.com/axieinfinity/bridge-v2/stores"
	bridgeUtils "github.com/axieinfinity/bridge-v2/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
)

const (
	defaultLimitRecords       = 50
	defaultMaxTry             = 5
	defaultMaxProcessingTasks = 200
)

var (
	defaultTaskInterval = 10 * time.Second
	defaultReceiptCheck = 50 * time.Second
)

type RoninTask struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	util utils.Utils

	listener bridgeCore.Listener
	store    stores.BridgeStore

	taskInterval    time.Duration
	txCheckInterval time.Duration
	secret          *bridgeCore.Secret

	client    *ethclient.Client
	contracts map[string]string

	limitQuery int
	chainId    *big.Int

	releaseTasksCh chan int

	processingIdsMap   sync.Map
	maxProcessingTasks int
}

func NewRoninTask(listener bridgeCore.Listener, db *gorm.DB, util utils.Utils) (*RoninTask, error) {
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
		ctx:                newCtx,
		cancelFunc:         cancelFunc,
		listener:           listener,
		store:              stores.NewBridgeStore(db),
		taskInterval:       defaultTaskInterval,
		txCheckInterval:    defaultReceiptCheck,
		secret:             config.Secret,
		contracts:          config.Contracts,
		client:             client,
		chainId:            chainId,
		util:               util,
		limitQuery:         defaultLimitRecords,
		releaseTasksCh:     make(chan int, defaultLimitRecords),
		maxProcessingTasks: defaultMaxProcessingTasks,
	}
	if config.TaskInterval > 0 {
		task.taskInterval = config.TaskInterval
	}
	if config.TransactionCheckPeriod > 0 {
		task.txCheckInterval = config.TransactionCheckPeriod
	}
	if config.MaxTasksQuery > 0 {
		task.limitQuery = config.MaxTasksQuery
	}
	if config.MaxProcessingTasks > 0 {
		task.maxProcessingTasks = config.MaxProcessingTasks
	}
	return task, nil
}

func (r *RoninTask) Start() {
	log.Info("[RoninTask] starting ronin task", "taskInterval", r.taskInterval, "txCheckInterval", r.txCheckInterval, "maxProcessingTasks", r.maxProcessingTasks)
	taskTicker := time.NewTicker(r.taskInterval)
	processingTicker := time.NewTicker(r.txCheckInterval)

	ethConfig := r.listener.GetListener(bridgeUtils.Ethereum).Config()
	ethClient, _ := ethclient.Dial(ethConfig.RpcUrl)

	for {
		select {
		case <-taskTicker.C:
			go func() {
				if err := r.processPending(ethClient); err != nil {
					log.Error("[RoninTask] error while process tasks", "err", err)
				}
			}()
		case <-processingTicker.C:
			go func() {
				if err := r.checkProcessingTasks(); err != nil {
					log.Error("[RoninTask] error while checking processing tasks", "err", err)
				}
			}()
		case id := <-r.releaseTasksCh:
			r.unlockTask(id)

		case <-r.ctx.Done():
			r.client.Close()
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

func (r *RoninTask) GetListener() bridgeCore.Listener {
	return r.listener
}

// getLimitQuery gets limitQuery which makes sure total processing tasks are within maxProcessingTasks
func (r *RoninTask) getLimitQuery(numberOfExcludedIds int) int {
	if numberOfExcludedIds >= r.maxProcessingTasks {
		return 0
	}
	if r.maxProcessingTasks-numberOfExcludedIds < r.limitQuery {
		return r.maxProcessingTasks - numberOfExcludedIds
	}
	return r.limitQuery
}

func (r *RoninTask) processPending(ethClient *ethclient.Client) error {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[RoninTask][processPending] recover from panic", "err", err)
		}
	}()
	// load processing tasks into excluded list
	var excludeIds []int
	r.processingIdsMap.Range(func(key, value interface{}) bool {
		excludeIds = append(excludeIds, key.(int))
		return true
	})
	// get limitQuery, if limitQuery = 0 then do nothing and wait until processing tasks are released
	limitQuery := r.getLimitQuery(len(excludeIds))
	if limitQuery == 0 {
		return nil
	}
	tasks, err := r.store.GetTaskStore().GetTasks(hexutil.EncodeBig(r.chainId), STATUS_PENDING, limitQuery, 10, 0, excludeIds)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return nil
	}
	metrics.Pusher.IncrCounter(metrics.PendingTaskMetric, len(tasks))

	bulkDepositTask := newBulkTask(r.listener, r.client, r.store, r.chainId, r.contracts, r.txCheckInterval, defaultMaxTry, DEPOSIT_TASK, r.releaseTasksCh, r.util)
	bulkSubmitWithdrawalSignaturesTask := newBulkTask(r.listener, r.client, r.store, r.chainId, r.contracts, r.txCheckInterval, defaultMaxTry, WITHDRAWAL_TASK, r.releaseTasksCh, r.util)
	ackWithdrewTasks := newBulkTask(r.listener, r.client, r.store, r.chainId, r.contracts, r.txCheckInterval, defaultMaxTry, ACK_WITHDREW_TASK, r.releaseTasksCh, r.util)
	voteBridgeOperatorsTask := newTask(r.listener, r.client, ethClient, r.store, r.chainId, r.contracts, defaultMaxTry, VOTE_BRIDGE_OPERATORS_TASK, r.releaseTasksCh, r.util)
	relayBridgeOperatorsTask := newTask(r.listener, r.client, ethClient, r.store, r.chainId, r.contracts, defaultMaxTry, RELAY_BRIDGE_OPERATORS_TASK, r.releaseTasksCh, r.util)

	for _, task := range tasks {
		// lock task
		r.lockTask(task)

		// collect tasks for bulk deposits
		bulkDepositTask.collectTask(task)

		// collect tasks for bulk withdrawal signature
		bulkSubmitWithdrawalSignaturesTask.collectTask(task)

		// collect tasks for acknowledge withdrawal
		ackWithdrewTasks.collectTask(task)

		// collect task for vote bridge operators
		voteBridgeOperatorsTask.collectTask(task)

		// collect task for relay bridge operators
		relayBridgeOperatorsTask.collectTask(task)
	}
	bulkDepositTask.send()
	bulkSubmitWithdrawalSignaturesTask.send()
	ackWithdrewTasks.send()
	voteBridgeOperatorsTask.send()
	relayBridgeOperatorsTask.send()
	return nil
}

func (r *RoninTask) lockTask(t *models.Task) {
	if t != nil {
		r.processingIdsMap.Store(t.ID, struct{}{})
	}
}

func (r *RoninTask) unlockTask(id int) {
	r.processingIdsMap.Delete(id)
}

func (r *RoninTask) checkProcessingTasks() error {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[RoninTask][checkProcessingTask] recover from panic", "err", err)
		}
	}()
	before := time.Now().Unix() - int64(r.listener.Config().SafeBlockRange*uint64(r.listener.Config().LoadInterval.Seconds()))
	tasks, err := r.store.GetTaskStore().GetTasks(hexutil.EncodeBig(r.chainId), STATUS_PROCESSING, r.limitQuery, 2, before, nil)
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
	metrics.Pusher.IncrGauge(metrics.ProcessingTaskMetric, -len(tasks))
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
		metrics.Pusher.IncrCounter(metrics.SuccessTaskMetric, len(successTxs))
		if err = r.store.GetTaskStore().UpdateTasksWithTransactionHash(successTxs, 1, STATUS_DONE); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while update tasks with success transactions", "err", err)
		}
	}

	// update success tasks with transaction's status = 0 (failed)
	if len(failedTxs) > 0 {
		metrics.Pusher.IncrCounter(metrics.FailedTaskMetric, len(failedTxs))

		if err = r.store.GetTaskStore().UpdateTasksWithTransactionHash(failedTxs, 0, STATUS_FAILED); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while update tasks with failed transactions", "err", err)
		}
	}

	// update all tasks contain failed txs to pending
	if len(resetToPending) > 0 {
		log.Info("[RoninTask][checkProcessingTasks] reset tasks to pending", "transactionHash", resetToPending)
		if err = r.store.GetTaskStore().ResetTo(resetToPending, STATUS_PENDING); err != nil {
			log.Error("[RoninTask][checkProcessingTasks] error while reset tasks to pending", "err", err)
		}
	}

	return nil
}
