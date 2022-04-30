package listener

import (
	"context"
	"errors"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	TxEvent = iota
	LogEvent
)

const (
	defaultWorkers          = 2048
	defaultMaxQueueSize     = 4096
	defaultCoolDownDuration = 1
)

type Controller struct {
	once       sync.Once
	ctx        context.Context
	cancelFunc context.CancelFunc

	listeners   map[string]IListener
	HandlerABIs map[string]*abi.ABI
	utilWrapper utils.IUtils

	Workers []*Worker

	// message backoff
	MaxRetry int32
	BackOff  int32

	// coolDownDuration is used to sleep for a while when a channel reaches its size
	coolDownDuration int

	// Queue holds a list of worker
	Queue chan chan IJob

	// JobChan receives new job
	JobChan chan IJob

	PrepareJobChan chan IJob

	jobId         int32
	processedJobs sync.Map

	MaxQueueSize int
}

type Config struct {
	Listeners       map[string]*LsConfig `json:"listeners"`
	NumberOfWorkers int                  `json:"numberOfWorkers"`
}

type LsConfig struct {
	Name           string        `json:"-"`
	RpcUrl         string        `json:"rpcUrl"`
	LoadInterval   time.Duration `json:"blockTime"`
	SafeBlockRange uint64        `json:"safeBlockRange"`
	FromHeight     uint64        `json:"fromHeight"`

	// TODO: apply more ways to get privatekey. such as: PLAINTEXT, KMS, etc.
	Secret        Secret                `json:"secret"`
	Subscriptions map[string]*Subscribe `json:"subscriptions"`
}

type Secret struct {
	Validator string `json:"validator"`
	Relayer   string `json:"relayer"`
}

type Subscribe struct {
	From string `json:"from"`
	To   string `json:"to"`

	// Type can be either TxEvent or LogEvent
	Type int `json:"type"`

	Handler   *Handler              `json:"handler"`
	CallBacks map[string][]*Handler `json:"callbacks"`
}

type Handler struct {
	// ABIPath specifies path to abi file
	ABIPath string `json:"abi"`

	// Name is method/event name
	Name string `json:"name"`

	// ContractAddress is used in callback case
	ContractAddress string `json:"contractAddress"`

	// Listener who triggers callback event
	Listener string `json:"listener"`

	ABI *abi.ABI `json:"-"`
}

func New(cfg *Config) (*Controller, error) {
	if cfg.NumberOfWorkers <= 0 {
		cfg.NumberOfWorkers = defaultWorkers
	}
	ctx, cancel := context.WithCancel(context.Background())
	c := &Controller{
		ctx:              ctx,
		cancelFunc:       cancel,
		listeners:        make(map[string]IListener),
		HandlerABIs:      make(map[string]*abi.ABI),
		utilWrapper:      &utils.Utils{},
		Workers:          make([]*Worker, 0),
		MaxRetry:         100,
		BackOff:          5,
		MaxQueueSize:     defaultMaxQueueSize,
		coolDownDuration: defaultCoolDownDuration,
	}
	c.JobChan = make(chan IJob, c.MaxQueueSize*cfg.NumberOfWorkers)
	c.PrepareJobChan = make(chan IJob, c.MaxQueueSize)
	c.Queue = make(chan chan IJob, c.MaxQueueSize)

	// add listeners from config
	for name, lsConfig := range cfg.Listeners {
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
		val, err := c.utilWrapper.Invoke(c, initFuncName, c.ctx, lsConfig)
		if err != nil {
			return nil, err
		}
		listener, ok := val.Interface().(IListener)
		if !ok {
			return nil, errors.New("cannot cast invoke result to IListener")
		}
		c.listeners[name] = listener
	}

	// init workers
	for i := 0; i < cfg.NumberOfWorkers; i++ {
		c.Workers = append(c.Workers, NewWorker(ctx, i, c.JobChan, c.Queue, c.MaxQueueSize, c.listeners))
	}
	return c, nil
}

// LoadABIsFromConfig loads all ABIPath and add results to Handler.ABI
func (c *Controller) LoadABIsFromConfig(lsConfig *LsConfig) (err error) {
	for _, subscription := range lsConfig.Subscriptions {
		if subscription.Handler.ABIPath == "" {
			continue
		}
		// load abi for handler
		if subscription.Handler.ABI, err = c.LoadAbi(subscription.Handler.ABIPath); err != nil {
			return
		}
		// load abi for callbacks handler
		for _, cbs := range subscription.CallBacks {
			for _, cb := range cbs {
				if cb.ABI, err = c.LoadAbi(cb.ABIPath); err != nil {
					return
				}
			}
		}
	}
	return
}

func (c *Controller) Start() error {
	for _, worker := range c.Workers {
		go worker.start()
	}
	for _, listener := range c.listeners {
		go listener.Start()
		go c.startListener(listener, 0)
	}
	// run all events listeners
	go func() {
		for {
			select {
			case job := <-c.PrepareJobChan:
				// TODO: add job into database before sending job to jobChan
				c.JobChan <- job
			case job := <-c.JobChan:
				// get 1 workerCh from queue and push job to this channel
				hash := job.Hash()
				if _, ok := c.processedJobs.Load(hash); ok {
					continue
				}
				c.processedJobs.Store(hash, struct{}{})
				log.Info("jobChan received a job", "jobId", job.GetID(), "nextTry", job.GetNextTry(), "type", job.GetType())
				workerCh := <-c.Queue
				workerCh <- job
			case <-c.ctx.Done():
				c.once.Do(func() {
					close(c.Queue)
					close(c.JobChan)
				})
				break
			}
		}
	}()
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

func (c *Controller) startListener(listener IListener, tryCount int) {
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

	// start processing past blocks
	currentBlock := listener.GetCurrentBlock()
	if currentBlock != nil {
		for currentBlock.GetHeight() <= latestBlock.GetHeight() {
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

	tick := time.NewTicker(listener.Period())
	for {
		select {
		case <-listener.Context().Done():
			return
		case <-tick.C:
			block, err := listener.GetLatestBlock()
			if err != nil {
				log.Error("[Controller][Process] error while get latest block", "err", err, "listener", listener.GetName())
				continue
			}
			c.Process(listener, block)
		}
	}
}

func (c *Controller) Process(listener IListener, latestBlock IBlock) {
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
		switch subscribe.Type {
		case TxEvent:
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
					c.JobChan <- job
				}
			}
		case LogEvent:
			for _, receipt := range receipts {
				if !receipt.GetStatus() {
					continue
				}
				tx := receipt.GetTransaction()
				if subscribe.To != "" && subscribe.To != tx.GetToAddress() {
					continue
				}
				for _, logData := range receipt.GetLogs() {
					if job := listener.GetListenHandleJob(name, tx, logData.GetData()); job != nil {
						c.JobChan <- job
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
