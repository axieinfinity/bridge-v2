package listener

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	bridgeCoreModels "github.com/axieinfinity/bridge-core/models"
	"github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type EthereumListener struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	config *bridgeCore.LsConfig

	chainId *big.Int
	jobId   int32

	rpcUrl         string
	name           string
	period         time.Duration
	currentBlock   atomic.Value
	safeBlockRange uint64
	fromHeight     uint64
	batches        sync.Map
	utilsWrapper   utils.Utils
	client         utils.EthClient
	validatorKey   *ecdsa.PrivateKey
	relayerKey     *ecdsa.PrivateKey
	store          stores.MainStore

	prepareJobChan chan bridgeCore.JobHandler
}

func NewEthereumListener(ctx context.Context, cfg *bridgeCore.LsConfig, helpers utils.Utils, store stores.MainStore) (*EthereumListener, error) {
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
		chainId:        hexutil.MustDecodeBig(cfg.ChainId),
		safeBlockRange: cfg.SafeBlockRange,
	}
	if helpers != nil {
		ethListener.utilsWrapper = helpers
	}
	client, err := ethListener.utilsWrapper.NewEthClient(cfg.RpcUrl)
	if err != nil {
		log.Error("[NewEthereumListener] error while dialing rpc client", "err", err, "url", cfg.RpcUrl)
		return nil, err
	}
	ethListener.client = client
	if cfg.Secret.Validator != "" {
		ethListener.validatorKey, err = crypto.HexToECDSA(cfg.Secret.Validator)
		if err != nil {
			log.Error("[NewEthereumListener] error while getting validator key", "err", err)
			return nil, err
		}
	}
	if cfg.Secret.Relayer != "" {
		ethListener.relayerKey, err = crypto.HexToECDSA(cfg.Secret.Relayer)
		if err != nil {
			log.Error("[NewEthereumListener] error while getting relayer key", "err", err)
			return nil, err
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

func (e *EthereumListener) GetTask() bridgeCore.TaskHandler {
	return nil
}

func (e *EthereumListener) GetCurrentBlock() bridgeCore.Block {
	if _, ok := e.currentBlock.Load().(bridgeCore.Block); !ok {
		var (
			block bridgeCore.Block
			err   error
		)
		block, err = e.GetProcessedBlock()
		if err != nil {
			log.Error(fmt.Sprintf("[%s] error on getting processed block from database", e.GetName()), "err", err)
			if e.fromHeight > 0 {
				block, err = e.GetBlock(e.fromHeight)
				if err != nil {
					log.Error(fmt.Sprintf("[%s] error on getting block from rpc", e.GetName()), "err", err, "fromHeight", e.fromHeight)
				}
			}
		}
		// if block is still nil, get latest block from rpc
		if block == nil {
			block, err = e.GetLatestBlock()
			if err != nil {
				log.Error(fmt.Sprintf("[%s] error on getting latest block from rpc", e.GetName()), "err", err)
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
		log.Error("[EthereumListener][GetLatestBlock] error while getting chainId", "err", err)
		return nil, err
	}
	encodedChainId := hexutil.EncodeBig(chainId)
	height, err := e.store.GetProcessedBlockStore().GetLatestBlock(encodedChainId)
	if err != nil {
		log.Error("[EthereumListener][GetLatestBlock] error while getting latest height from database", "err", err, "chainId", encodedChainId)
		return nil, err
	}
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(height))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, chainId, block, false)
}

func (e *EthereumListener) GetLatestBlock() (bridgeCore.Block, error) {
	block, err := e.client.BlockByNumber(e.ctx, nil)
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, e.chainId, block, false)
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
		log.Info(fmt.Sprintf("[%s] UpdateCurrentBlock", e.name), "block", block.GetHeight())
		e.currentBlock.Store(block)
		return e.SaveCurrentBlockToDB()
	}
	return nil
}

func (e *EthereumListener) GetValidatorKey() *ecdsa.PrivateKey {
	return e.validatorKey
}

func (e *EthereumListener) GetRelayerKey() *ecdsa.PrivateKey {
	return e.relayerKey
}

func (e *EthereumListener) SaveCurrentBlockToDB() error {
	chainId, err := e.GetChainID()
	if err != nil {
		return err
	}

	if err := e.store.GetProcessedBlockStore().Save(hexutil.EncodeBig(chainId), int64(e.GetCurrentBlock().GetHeight())); err != nil {
		return err
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
	log.Info("[EthereumListener][SendCallbackJobs] Start", "subscriptionName", subscriptionName, "listeners", len(listeners), "fromTx", tx.GetHash().Hex())
	chainId, err := e.GetChainID()
	if err != nil {
		return
	}
	subscription, ok := e.GetSubscriptions()[subscriptionName]
	if !ok {
		log.Warn("[EthereumListener][SendCallbackJobs] cannot find subscription", "subscriptionName", subscriptionName)
		return
	}
	log.Info("[EthereumListener][SendCallbackJobs] subscription found", "subscriptionName", subscriptionName, "numberOfCallbacks", len(subscription.CallBacks))
	for listenerName, methodName := range subscription.CallBacks {
		log.Info("[EthereumListener][SendCallbackJobs] Loop through callbacks", "subscriptionName", subscriptionName, "listenerName", listenerName, "methodName", methodName)
		l := listeners[listenerName]
		job := NewEthCallbackJob(l, methodName, tx, inputData, chainId, e.utilsWrapper)
		if job != nil {
			e.prepareJobChan <- job
		}
	}
}

func (e *EthereumListener) GetBlock(height uint64) (bridgeCore.Block, error) {
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, e.chainId, block, false)
}

func (e *EthereumListener) GetBlockWithLogs(height uint64) (bridgeCore.Block, error) {
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, e.chainId, block, true)
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
