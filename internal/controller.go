package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/stores"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	defaultWorkers          = 2048
	defaultMaxQueueSize     = 4096
	defaultCoolDownDuration = 1
)

type Controller struct {
	lock       sync.Mutex
	ctx        context.Context
	cancelFunc context.CancelFunc

	listeners   map[string]types.IListener
	HandlerABIs map[string]*abi.ABI
	utilWrapper utils.IUtils

	Workers []*Worker

	// message backoff
	MaxRetry int32
	BackOff  int32

	// coolDownDuration is used to sleep for a while when a channel reaches its size
	coolDownDuration int

	// Queue holds a list of worker
	Queue chan chan types.IJob

	// JobChan receives new job
	JobChan        chan types.IJob
	SuccessJobChan chan types.IJob
	FailedJobChan  chan types.IJob
	PrepareJobChan chan types.IJob

	jobId         int32
	processedJobs sync.Map

	MaxQueueSize int
	cfg          *types.Config

	store types.IMainStore
}

func New(cfg *types.Config, db *gorm.DB, helpers utils.IUtils) (*Controller, error) {
	if cfg.NumberOfWorkers <= 0 {
		cfg.NumberOfWorkers = defaultWorkers
	}
	ctx, cancel := context.WithCancel(context.Background())
	c := &Controller{
		cfg:              cfg,
		ctx:              ctx,
		cancelFunc:       cancel,
		listeners:        make(map[string]types.IListener),
		HandlerABIs:      make(map[string]*abi.ABI),
		utilWrapper:      &utils.Utils{},
		Workers:          make([]*Worker, 0),
		MaxRetry:         100,
		BackOff:          5,
		MaxQueueSize:     defaultMaxQueueSize,
		coolDownDuration: defaultCoolDownDuration,
		store:            stores.NewMainStore(db),
	}
	if helpers != nil {
		c.utilWrapper = helpers
	}
	c.JobChan = make(chan types.IJob, c.MaxQueueSize*cfg.NumberOfWorkers)
	c.PrepareJobChan = make(chan types.IJob, c.MaxQueueSize)
	c.SuccessJobChan = make(chan types.IJob, c.MaxQueueSize)
	c.FailedJobChan = make(chan types.IJob, c.MaxQueueSize)
	c.Queue = make(chan chan types.IJob, c.MaxQueueSize)

	// add listeners from config
	for name, lsConfig := range c.cfg.Listeners {
		if lsConfig.LoadInterval <= 0 {
			continue
		}
		lsConfig.LoadInterval *= time.Second
		lsConfig.Name = name
		// load abi from lsConfig
		if err := c.LoadABIsFromConfig(lsConfig); err != nil {
			return nil, err
		}
		// Invoke init function which is based on listener's name
		initFuncName := fmt.Sprintf("Init%s", name)
		val, err := c.utilWrapper.Invoke(c, initFuncName, c.ctx, lsConfig, c.store)
		if err != nil {
			return nil, err
		}
		l, ok := val.Interface().(types.IListener)
		if !ok {
			return nil, errors.New("cannot cast invoke result to IListener")
		}
		if l == nil {
			return nil, errors.New("listener is nil")
		}
		c.listeners[name] = l
	}

	// init workers
	for i := 0; i < cfg.NumberOfWorkers; i++ {
		c.Workers = append(c.Workers, NewWorker(ctx, i, c.PrepareJobChan, c.FailedJobChan, c.SuccessJobChan, c.Queue, c.MaxQueueSize, c.listeners))
	}
	return c, nil
}

// LoadABIsFromConfig loads all ABIPath and add results to Handler.ABI
func (c *Controller) LoadABIsFromConfig(lsConfig *types.LsConfig) (err error) {
	for _, subscription := range lsConfig.Subscriptions {
		if subscription.Handler.ABIPath == "" {
			continue
		}
		// load abi for handler
		if subscription.Handler.ABI, err = c.LoadAbi(subscription.Handler.ABIPath); err != nil {
			return
		}
	}
	return
}

func (c *Controller) prepareJob(job types.IJob) error {
	if job.GetID() == 0 {
		return job.Save()
	}
	return nil
}

func (c *Controller) Start() error {
	for _, worker := range c.Workers {
		go worker.start()
	}
	// TODO: load jobs from database and reprocess before start listener
	go func() {
		for {
			select {
			case job := <-c.SuccessJobChan:
				log.Info("process job success", "id", job.GetID())
				if err := job.Update(types.STATUS_DONE); err != nil {
					log.Error("[Controller] failed on updating success job", "err", err, "jobType", job.GetType(), "tx", job.GetTransaction().GetHash().Hex())
					// send back job to successJobChan
					c.SuccessJobChan <- job
				}
			case job := <-c.FailedJobChan:
				log.Info("process job failed", "id", job.GetID())
				if err := job.Update(types.STATUS_FAILED); err != nil {
					log.Error("[Controller] failed on updating failed job", "err", err, "jobType", job.GetType(), "tx", job.GetTransaction().GetHash().Hex())
					// send back job to failedJobChan
					c.FailedJobChan <- job
				}
			case job := <-c.PrepareJobChan:
				// add new job to database before processing
				if err := c.prepareJob(job); err != nil {
					log.Error("[Controller] failed on preparing job", "err", err, "jobType", job.GetType(), "tx", job.GetTransaction().GetHash().Hex())
					continue
				}
				c.JobChan <- job
			case job := <-c.JobChan:
				// get 1 workerCh from queue and push job to this channel
				hash := job.Hash()
				if _, ok := c.processedJobs.Load(hash); ok {
					continue
				}
				c.processedJobs.Store(hash, struct{}{})
				log.Info("[Controller] jobChan received a job", "jobId", job.GetID(), "nextTry", job.GetNextTry(), "type", job.GetType())
				workerCh := <-c.Queue
				workerCh <- job
			case <-c.ctx.Done():
				c.lock.Lock()
				jobs := make([]types.IJob, 0)
				for job := range c.PrepareJobChan {
					if err := c.prepareJob(job); err != nil {
						log.Error("[Controller] error while storing all jobs from prepareJobChan to database in closing step", "err", err, "jobType", job.GetType(), "tx", job.GetTransaction().GetHash().Hex())
						continue
					}
					jobs = append(jobs, job)
				}
				close(c.PrepareJobChan)
				c.lock.Unlock()

				// wait until all jobs are handled
				for len(c.JobChan) > 0 {
				}

				// wait until all success jobs are handled
				for len(c.SuccessJobChan) > 0 {
				}

				// wait until all failed jobs are handled
				for len(c.FailedJobChan) > 0 {
				}

				close(c.JobChan)
				close(c.SuccessJobChan)
				close(c.FailedJobChan)
				close(c.Queue)
				break
			}
		}
	}()

	// load all pending jobs from database
	jobs, err := c.store.GetJobStore().GetPendingJobs()
	if err != nil {
		// just log and do nothing.
		log.Error("[Controller] error while getting pending jobs from database", "err", err)
	}
	for _, job := range jobs {
		listener, ok := c.listeners[job.Listener]
		if !ok {
			continue
		}
		j, err := listener.NewJobFromDB(job)
		if err != nil {
			log.Error("[Controller] error while init job from db", "err", err, "jobId", job.ID)
			continue
		}
		// add job to jobChan
		c.JobChan <- j
	}

	// run all events listeners
	for _, listener := range c.listeners {
		go listener.Start()
		go c.startListener(listener, 0)
	}

	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		select {
		case <-sigc:
			c.Close()
		}
	}()
	return nil
}

// startListener starts listening events for a listener, it comes with a tryCount which close this listener if tryCount reaches 10 times
func (c *Controller) startListener(listener types.IListener, tryCount int) {
	// panic when tryCount reaches 10 times panic
	if tryCount >= 10 {
		log.Error("[Controller][startListener] maximum try has been reached, close listener", "listener", listener.GetName())
		listener.Close()
		return
	}

	// check if listener is behind or not
	latestBlock, err := listener.GetLatestBlock()
	if err != nil {
		log.Error("[Controller][startListener] error while get latest block", "err", err, "listener", listener.GetName())
		// otherwise retry startListener
		time.Sleep(time.Duration(tryCount+1) * time.Second)
		go c.startListener(listener, tryCount+1)
		return
	}
	log.Info("[Controller] Latest Block", "height", latestBlock.GetHeight(), "listener", listener.GetName())
	// start processing past blocks
	currentBlock := listener.GetCurrentBlock()
	if currentBlock != nil {
		for currentBlock.GetHeight() < latestBlock.GetHeight() {
			block, err := listener.GetBlock(currentBlock.GetHeight() + 1)
			if err != nil {
				log.Error("[Controller][startListener] error while get latest block", "err", err, "listener", listener.GetName())
				time.Sleep(time.Duration(tryCount+1) * time.Second)
				go c.startListener(listener, tryCount+1)
				return
			}
			c.Process(listener, block)
			currentBlock = listener.GetCurrentBlock()
		}
	}
	// start listening to block's events
	tick := time.NewTicker(listener.Period())
	for {
		select {
		case <-listener.Context().Done():
			return
		case <-tick.C:
			currentBlock = listener.GetCurrentBlock()
			block, err := listener.GetBlock(currentBlock.GetHeight() + 1)
			if err != nil {
				log.Error("[Controller][Process] error while get latest block", "err", err, "listener", listener.GetName())
				continue
			}
			c.Process(listener, block)
		}
	}
}

func (c *Controller) Process(listener types.IListener, latestBlock types.IBlock) {
	log.Info("[Controller][Process] start processing block", "height", latestBlock.GetHeight())
	// if latest block has been processed then do nothing
	if listener.GetCurrentBlock() != nil && latestBlock.GetHeight() == listener.GetCurrentBlock().GetHeight() {
		return
	}
	if err := listener.UpdateCurrentBlock(latestBlock); err != nil {
		log.Error("[Controller][Process] error while updating current block", "err", err, "listener", listener.GetName())
		return
	}
	// get list of transactions in new block
	txs := latestBlock.GetTransactions()
	receipts := latestBlock.GetReceipts()
	if len(txs) != len(receipts) {
		log.Error("[Controller][Process] txs and receipts are mismatched", "txs", len(txs), "receipts", len(receipts))
		return
	}
	for name, subscribe := range listener.GetSubscriptions() {
		if subscribe.Handler == nil {
			continue
		}
		log.Info("[Controller][Process] processing block", "block", latestBlock.GetHeight(), "txs", len(txs), "receipts", len(receipts), "subscribeType", subscribe.Type)
		switch subscribe.Type {
		case types.TxEvent:
			for _, tx := range txs {
				if (subscribe.From != "" && subscribe.From != tx.GetFromAddress()) ||
					(subscribe.To != "" && subscribe.To != tx.GetToAddress()) {
					continue
				}
				receipt := tx.GetReceipt()
				if !receipt.GetStatus() {
					continue
				}
				if job := listener.GetListenHandleJob(name, tx, tx.GetData()); job != nil {
					c.PrepareJobChan <- job
				}
			}
		case types.LogEvent:
			for _, receipt := range receipts {
				log.Info("[Controller][Process] processing log Event", "receipt status", receipt.GetStatus(), "logs", len(receipt.GetLogs()))
				if !receipt.GetStatus() {
					continue
				}
				tx := receipt.GetTransaction()
				if subscribe.To != tx.GetToAddress() {
					continue
				}
				log.Info("[Controller][Process] generating event job", "height", latestBlock.GetHeight())
				for _, logData := range receipt.GetLogs() {
					if job := listener.GetListenHandleJob(name, tx, logData.GetData()); job != nil {
						c.PrepareJobChan <- job
					}
				}
			}
		}
	}
}

func (c *Controller) LoadAbi(path string) (*abi.ABI, error) {
	if _, ok := c.HandlerABIs[path]; ok {
		return c.HandlerABIs[path], nil
	}
	a, err := c.utilWrapper.LoadAbi(path)
	if err != nil {
		return nil, err
	}
	c.HandlerABIs[path] = a
	return a, nil
}

func (c *Controller) Close() {
	c.cancelFunc()
}
