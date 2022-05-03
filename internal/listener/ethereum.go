package listener

import (
	"context"
	"crypto/ecdsa"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

type EthereumListener struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	config *types.LsConfig

	chainId *big.Int
	jobId   int32

	rpcUrl         string
	name           string
	period         time.Duration
	currentBlock   atomic.Value
	safeBlockRange uint64
	fromHeight     uint64
	batches        sync.Map
	utilsWrapper   utils.IUtils
	client         utils.EthClient
	validatorKey   *ecdsa.PrivateKey
	relayerKey     *ecdsa.PrivateKey
	store          types.IMainStore
}

func NewEthereumListener(ctx context.Context, cfg *types.LsConfig, helpers utils.IUtils, store types.IMainStore) (*EthereumListener, error) {
	newCtx, cancelFunc := context.WithCancel(ctx)
	ethListener := &EthereumListener{
		name:         cfg.Name,
		period:       cfg.LoadInterval,
		currentBlock: atomic.Value{},
		ctx:          newCtx,
		cancelCtx:    cancelFunc,
		fromHeight:   cfg.FromHeight,
		utilsWrapper: &utils.Utils{},
		store:        store,
		config:       cfg,
	}
	if helpers != nil {
		ethListener.utilsWrapper = helpers
	}
	client, err := ethListener.utilsWrapper.NewEthClient(cfg.RpcUrl)
	if err != nil {
		return nil, err
	}
	ethListener.client = client
	ethListener.validatorKey, err = crypto.HexToECDSA(cfg.Secret.Validator)
	if err != nil {
		return nil, err
	}
	ethListener.relayerKey, err = crypto.HexToECDSA(cfg.Secret.Relayer)
	if err != nil {
		return nil, err
	}
	return ethListener, nil
}

func (e *EthereumListener) GetStore() types.IMainStore {
	return e.store
}

func (e *EthereumListener) Config() *types.LsConfig {
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

func (e *EthereumListener) GetCurrentBlock() types.IBlock {
	if _, ok := e.currentBlock.Load().(types.IBlock); !ok {
		var (
			block types.IBlock
			err   error
		)
		if e.fromHeight > 0 {
			block, err = e.GetBlock(e.fromHeight)
		} else {
			block, err = e.GetLatestBlock()
		}
		if err != nil {
			return nil
		}
		e.currentBlock.Store(block)
		return block
	}
	return e.currentBlock.Load().(types.IBlock)
}

func (e *EthereumListener) GetLatestBlock() (types.IBlock, error) {
	// TODO: if apply ronin, lib change nil to -1
	block, err := e.client.BlockByNumber(e.ctx, nil)
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, block)
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

func (e *EthereumListener) GetSubscriptions() map[string]*types.Subscribe {
	return e.config.Subscriptions
}

func (e *EthereumListener) UpdateCurrentBlock(block types.IBlock) error {
	if e.GetCurrentBlock().GetHash().Hex() != block.GetHash().Hex() {
		e.currentBlock.Store(block)
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
	return e.store.GetProcessedBlockStore().Save(int64(e.GetCurrentBlock().GetHeight()))
}

func (e *EthereumListener) SaveTransactionsToDB(txs []types.ITransaction) error {
	return nil
}

func (e *EthereumListener) GetListenHandleJob(subscriptionName string, tx types.ITransaction, data []byte) types.IJob {
	// validate if data contains subscribed name
	subscription, ok := e.GetSubscriptions()[subscriptionName]
	if !ok {
		return nil
	}
	handlerName := subscription.Handler.Name
	if subscription.Type == types.TxEvent {
		method, ok := subscription.Handler.ABI.Methods[handlerName]
		if !ok {
			return nil
		}
		if method.RawName != common.Bytes2Hex(data[0:4]) {
			return nil
		}
	} else if subscription.Type == types.LogEvent {
		event, ok := subscription.Handler.ABI.Events[handlerName]
		if !ok {
			return nil
		}
		if common.Bytes2Hex(event.ID.Bytes()[0:4]) != common.Bytes2Hex(data[0:4]) {
			return nil
		}
	} else {
		return nil
	}
	jobId := atomic.AddInt32(&e.jobId, 1)
	return NewEthListenJob(jobId, types.ListenHandler, e, subscriptionName, tx, data)
}

func (e *EthereumListener) SendCallbackJobs(listeners map[string]types.IListener, subscriptionName string, tx types.ITransaction, inputData []byte, jobChan chan<- types.IJob) {
	log.Info("[EthereumListener][SendCallbackJobs] Start", "subscriptionName", subscriptionName, "listeners", len(listeners))
	chainId, err := e.GetChainID()
	if err != nil {
		return
	}
	subscription, ok := e.GetSubscriptions()[subscriptionName]
	if !ok {
		return
	}
	for listenerName, methodName := range subscription.CallBacks {
		log.Info("[EthereumListener][SendCallbackJobs] Loop through callbacks", "subscriptionName", subscriptionName, "listenerName", listenerName, "methodName", methodName)
		l := listeners[listenerName]
		jobId := atomic.AddInt32(&e.jobId, 1)
		job := NewEthCallbackJob(jobId, l, methodName, tx, inputData, chainId, e.utilsWrapper)
		if job != nil {
			jobChan <- job
		}
	}
}

func (e *EthereumListener) GetBlock(height uint64) (types.IBlock, error) {
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, block)
}

func (e *EthereumListener) Close() {
	e.client.Close()
	e.cancelCtx()
}
