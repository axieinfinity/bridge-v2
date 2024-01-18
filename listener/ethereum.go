package listener

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/axieinfinity/bridge-v2/stats"

	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	bridgeCoreModels "github.com/axieinfinity/bridge-core/models"
	"github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	bridgeCoreUtils "github.com/axieinfinity/bridge-core/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type EthereumListener struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	config *bridgeCore.LsConfig

	chainId *big.Int
	jobId   int32

	rpcUrl                   string
	name                     string
	period                   time.Duration
	currentBlock             atomic.Value
	safeBlockRange           uint64
	fromHeight               uint64
	batches                  sync.Map
	utilsWrapper             utils.Utils
	client                   utils.EthClient
	bridgeOperatorSign       bridgeCoreUtils.ISign
	voterSign                bridgeCoreUtils.ISign
	relayerSign              bridgeCoreUtils.ISign
	legacyBridgeOperatorSign bridgeCoreUtils.ISign
	store                    stores.MainStore
	listeners                map[string]bridgeCore.Listener

	prepareJobChan chan bridgeCore.JobHandler
	tasks          []bridgeCore.TaskHandler
	pool           *bridgeCore.Pool
}

func (e *EthereumListener) AddListeners(m map[string]bridgeCore.Listener) {
	e.listeners = m
}

func (e *EthereumListener) GetListener(s string) bridgeCore.Listener {
	return e.listeners[s]
}

func NewEthereumListener(ctx context.Context, cfg *bridgeCore.LsConfig, helpers utils.Utils, store stores.MainStore, pool *bridgeCore.Pool) (*EthereumListener, error) {
	newCtx, cancelFunc := context.WithCancel(ctx)
	ethListener := &EthereumListener{
		name:           cfg.Name,
		period:         cfg.LoadInterval,
		currentBlock:   atomic.Value{},
		ctx:            newCtx,
		cancelCtx:      cancelFunc,
		fromHeight:     cfg.FromHeight,
		utilsWrapper:   utils.NewUtils(),
		store:          store,
		config:         cfg,
		listeners:      make(map[string]bridgeCore.Listener),
		chainId:        hexutil.MustDecodeBig(cfg.ChainId),
		safeBlockRange: cfg.SafeBlockRange,
		tasks:          make([]bridgeCore.TaskHandler, 0),
		pool:           pool,
	}
	if helpers != nil {
		ethListener.utilsWrapper = helpers
	}
	client, err := ethListener.utilsWrapper.NewEthClient(cfg.RpcUrl)
	if err != nil {
		log.Error(fmt.Sprintf("[New%sListener] error while dialing rpc client", cfg.Name), "err", err, "url", cfg.RpcUrl)
		return nil, err
	}
	ethListener.client = client

	if cfg.Secret.BridgeOperator != nil {
		ethListener.bridgeOperatorSign, err = bridgeCoreUtils.NewSignMethod(cfg.Secret.BridgeOperator)
		if err != nil {
			log.Error(fmt.Sprintf("[New%sListener] error while getting validator key", cfg.Name), "err", err)
			return nil, err
		}

		if ethListener.bridgeOperatorSign == nil {
			log.Warn(fmt.Sprintf("[%sListener] No sign method provided for operator key", cfg.Name))
		} else {
			log.Info(fmt.Sprintf("[%sListener] Operator account", cfg.Name), "address", ethListener.bridgeOperatorSign.GetAddress())
		}
	}
	if cfg.Secret.Voter != nil {
		ethListener.voterSign, err = bridgeCoreUtils.NewSignMethod(cfg.Secret.Voter)
		if err != nil {
			log.Error(fmt.Sprintf("[New%sListener] error while getting voter key", cfg.Name), "err", err)
			return nil, err
		}

		if ethListener.voterSign == nil {
			log.Warn(fmt.Sprintf("[%sListener] No sign method provided for voter key", cfg.Name))
		} else {
			log.Info(fmt.Sprintf("[%sListener] Voter account", cfg.Name), "address", ethListener.voterSign.GetAddress())
		}
	}
	if cfg.Secret.Relayer != nil {
		ethListener.relayerSign, err = bridgeCoreUtils.NewSignMethod(cfg.Secret.Relayer)
		if err != nil {
			log.Error(fmt.Sprintf("[New%sListener] error while getting relayer key", cfg.Name), "err", err)
			return nil, err
		}

		if ethListener.relayerSign == nil {
			log.Warn(fmt.Sprintf("[%sListener] No sign method provided for relay key", cfg.Name))
		} else {
			log.Info(fmt.Sprintf("[%sListener] Relayer account", cfg.Name), "address", ethListener.relayerSign.GetAddress())
		}
	}
	if cfg.Secret.LegacyBridgeOperator != nil {
		ethListener.legacyBridgeOperatorSign, err = bridgeCoreUtils.NewSignMethod(cfg.Secret.LegacyBridgeOperator)
		if err != nil {
			log.Error(fmt.Sprintf("[New%sListener] error while getting legacy bridge operator key", cfg.Name), "err", err)
			return nil, err
		}

		if ethListener.legacyBridgeOperatorSign == nil {
			log.Warn(fmt.Sprintf("[%sListener] No sign method provided for legacy bridge operator key", cfg.Name))
		} else {
			log.Info(fmt.Sprintf("[%sListener] Legacy bridge operator account", cfg.Name), "address", ethListener.legacyBridgeOperatorSign.GetAddress())
		}
	}
	return ethListener, nil
}

func (e *EthereumListener) GetStore() stores.MainStore {
	return e.store
}

func (e *EthereumListener) Config() *bridgeCore.LsConfig {
	return e.config
}

func (e *EthereumListener) Start() {
	for _, task := range e.tasks {
		go task.Start()
	}
}

func (e *EthereumListener) GetName() string {
	return e.name
}

func (e *EthereumListener) Period() time.Duration {
	return e.period
}

func (e *EthereumListener) GetSafeBlockRange() uint64 {
	return e.safeBlockRange
}

func (e *EthereumListener) IsDisabled() bool {
	return e.config.Disabled
}

func (e *EthereumListener) SetInitHeight(height uint64) {
	e.fromHeight = height
}

func (e *EthereumListener) GetInitHeight() uint64 {
	return e.fromHeight
}

func (e *EthereumListener) GetTask(index int) bridgeCore.TaskHandler {
	return e.tasks[index]
}

func (e *EthereumListener) GetTasks() []bridgeCore.TaskHandler {
	return e.tasks
}

func (e *EthereumListener) AddTask(task bridgeCore.TaskHandler) {
	e.tasks = append(e.tasks, task)
}

func (e *EthereumListener) GetCurrentBlock() bridgeCore.Block {
	if _, ok := e.currentBlock.Load().(bridgeCore.Block); !ok {
		var (
			block bridgeCore.Block
			err   error
		)
		block, err = e.GetProcessedBlock()
		if err != nil {
			log.Error(fmt.Sprintf("[%sListener] error on getting processed block from database", e.GetName()), "err", err)
			if e.fromHeight > 0 {
				block, err = e.GetBlock(e.fromHeight)
				if err != nil {
					log.Error(fmt.Sprintf("[%sListener] error on getting block from rpc", e.GetName()), "err", err, "fromHeight", e.fromHeight)
				}
			}
		}
		// if block is still nil, get latest block from rpc
		if block == nil {
			block, err = e.GetLatestBlock()
			if err != nil {
				log.Error(fmt.Sprintf("[%sListener] error on getting latest block from rpc", e.GetName()), "err", err)
				return nil
			}
		}
		e.currentBlock.Store(block)
		return block
	}
	return e.currentBlock.Load().(bridgeCore.Block)
}

func (e *EthereumListener) IsUpTodate() bool {
	return true
}

func (e *EthereumListener) GetProcessedBlock() (bridgeCore.Block, error) {
	chainId, err := e.GetChainID()
	if err != nil {
		log.Error(fmt.Sprintf("[%sListener][GetLatestBlock] error while getting chainId", e.GetName()), "err", err)
		return nil, err
	}
	encodedChainId := hexutil.EncodeBig(chainId)
	height, err := e.store.GetProcessedBlockStore().GetLatestBlock(encodedChainId)
	if err != nil {
		log.Error(fmt.Sprintf("[%sListener][GetLatestBlock] error while getting latest height from database", e.GetName()), "err", err, "chainId", encodedChainId)
		return nil, err
	}
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(height))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, block, false)
}

func (e *EthereumListener) GetLatestBlock() (bridgeCore.Block, error) {
	block, err := e.client.BlockByNumber(e.ctx, nil)
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, block, false)
}

func (e *EthereumListener) GetLatestBlockHeight() (uint64, error) {
	return e.client.BlockNumber(e.ctx)
}

func (e *EthereumListener) GetChainID() (*big.Int, error) {
	if e.chainId != nil {
		return e.chainId, nil
	}
	return e.client.ChainID(e.ctx)
}

func (e *EthereumListener) Context() context.Context {
	return e.ctx
}

func (e *EthereumListener) GetSubscriptions() map[string]*bridgeCore.Subscribe {
	return e.config.Subscriptions
}

func (e *EthereumListener) UpdateCurrentBlock(block bridgeCore.Block) error {
	if block != nil && e.GetCurrentBlock().GetHeight() < block.GetHeight() {
		log.Info(fmt.Sprintf("[%sListener] UpdateCurrentBlock", e.name), "block", block.GetHeight())
		e.currentBlock.Store(block)
		return e.SaveCurrentBlockToDB()
	}
	return nil
}

func (e *EthereumListener) SaveCurrentBlockToDB() error {
	chainId, err := e.GetChainID()
	if err != nil {
		return err
	}

	if err := e.store.GetProcessedBlockStore().Save(hexutil.EncodeBig(chainId), int64(e.GetCurrentBlock().GetHeight())); err != nil {
		return err
	}
	if stats.BridgeStats != nil {
		stats.BridgeStats.SendProcessedBlock(e.GetName(), e.GetCurrentBlock().GetHeight())
	}
	metrics.Pusher.IncrCounter(fmt.Sprintf(metrics.ListenerProcessedBlockMetric, e.GetName()), 1)
	return nil
}

func (e *EthereumListener) SaveTransactionsToDB(txs []bridgeCore.Transaction) error {
	return nil
}

func (e *EthereumListener) SetPrepareJobChan(jobChan chan bridgeCore.JobHandler) {
	e.prepareJobChan = jobChan
}

func (e *EthereumListener) GetEthClient() utils.EthClient {
	return e.client
}

func (e *EthereumListener) GetListenHandleJob(subscriptionName string, tx bridgeCore.Transaction, eventId string, data []byte) bridgeCore.JobHandler {
	// validate if data contains subscribed name
	subscription, ok := e.GetSubscriptions()[subscriptionName]
	if !ok {
		return nil
	}
	handlerName := subscription.Handler.Name
	if subscription.Type == bridgeCore.TxEvent {
		method, ok := subscription.Handler.ABI.Methods[handlerName]
		if !ok {
			return nil
		}
		if method.RawName != common.Bytes2Hex(data[0:4]) {
			return nil
		}
	} else if subscription.Type == bridgeCore.LogEvent {
		event, ok := subscription.Handler.ABI.Events[handlerName]
		if !ok {
			return nil
		}
		if hexutil.Encode(event.ID.Bytes()) != eventId {
			return nil
		}
	} else {
		return nil
	}
	return NewEthListenJob(bridgeCore.ListenHandler, e, subscriptionName, tx, data)
}

func (e *EthereumListener) SendCallbackJobs(listeners map[string]bridgeCore.Listener, subscriptionName string, tx bridgeCore.Transaction, inputData []byte) {
	log.Info(fmt.Sprintf("[%sListener][SendCallbackJobs] Start", e.GetName()), "subscriptionName", subscriptionName, "listeners", len(listeners), "fromTx", tx.GetHash().Hex())
	chainId, err := e.GetChainID()
	if err != nil {
		return
	}
	subscription, ok := e.GetSubscriptions()[subscriptionName]
	if !ok {
		log.Warn(fmt.Sprintf("[%sListener][SendCallbackJobs] cannot find subscription", e.GetName()), "subscriptionName", subscriptionName)
		return
	}
	log.Info(fmt.Sprintf("[%sListener][SendCallbackJobs] subscription found", e.GetName()), "subscriptionName", subscriptionName, "numberOfCallbacks", len(subscription.CallBacks))
	for listenerName, methodName := range subscription.CallBacks {
		log.Info(fmt.Sprintf("[%sListener][SendCallbackJobs] Loop through callbacks", e.GetName()), "subscriptionName", subscriptionName, "listenerName", listenerName, "methodName", methodName)
		l := listeners[listenerName]
		job := NewEthCallbackJob(l, methodName, tx, inputData, chainId, e.utilsWrapper)
		if job != nil {
			e.pool.Enqueue(job)
		}
	}
}

func (e *EthereumListener) GetBlock(height uint64) (bridgeCore.Block, error) {
	header, err := e.client.HeaderByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, ethtypes.NewBlockWithHeader(header), false)
}

func (e *EthereumListener) GetBlockWithLogs(height uint64) (bridgeCore.Block, error) {
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, block, true)
}

func (e *EthereumListener) GetReceipt(txHash common.Hash) (*ethtypes.Receipt, error) {
	return e.client.TransactionReceipt(e.ctx, txHash)
}

func (e *EthereumListener) NewJobFromDB(job *bridgeCoreModels.Job) (bridgeCore.JobHandler, error) {
	return newJobFromDB(e, job)
}

func newJobFromDB(listener bridgeCore.Listener, job *bridgeCoreModels.Job) (bridgeCore.JobHandler, error) {
	chainId, err := hexutil.DecodeBig(job.FromChainId)
	if err != nil {
		return nil, err
	}
	// get transaction from hash
	tx, _, err := listener.GetEthClient().TransactionByHash(context.Background(), common.HexToHash(job.Transaction))
	if err != nil {
		return nil, err
	}
	transaction, err := NewEthTransaction(chainId, tx)
	if err != nil {
		return nil, err
	}
	baseJob, err := bridgeCore.NewBaseJob(listener, job, transaction)
	if err != nil {
		return nil, err
	}
	switch job.Type {
	case bridgeCore.ListenHandler:
		return &EthListenJob{
			BaseJob: baseJob,
		}, nil
	case bridgeCore.CallbackHandler:
		if job.Method == "" {
			return nil, nil
		}
		return &EthCallbackJob{
			BaseJob: baseJob,
			method:  job.Method,
		}, nil
	}
	return nil, errors.New("jobType does not match")
}

func (e *EthereumListener) Close() {
	e.client.Close()
	e.cancelCtx()
}

func (e *EthereumListener) GetBridgeOperatorSign() bridgeCoreUtils.ISign {
	return e.bridgeOperatorSign
}

func (e *EthereumListener) GetVoterSign() bridgeCoreUtils.ISign {
	return e.voterSign
}

func (e *EthereumListener) GetRelayerSign() bridgeCoreUtils.ISign {
	return e.relayerSign
}

func (e *EthereumListener) GetLegacyBridgeOperatorSign() bridgeCoreUtils.ISign {
	return e.legacyBridgeOperatorSign
}

func (e *EthereumListener) CacheBlocks(blockNumbers map[uint64]struct{}) {
}
