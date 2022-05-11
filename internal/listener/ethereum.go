package listener

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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

	prepareJobChan chan types.IJob
}

func NewEthereumListener(ctx context.Context, cfg *types.LsConfig, helpers utils.IUtils, store types.IMainStore, prepareJobChan chan types.IJob) (*EthereumListener, error) {
	newCtx, cancelFunc := context.WithCancel(ctx)
	ethListener := &EthereumListener{
		name:           cfg.Name,
		period:         cfg.LoadInterval,
		currentBlock:   atomic.Value{},
		ctx:            newCtx,
		cancelCtx:      cancelFunc,
		fromHeight:     cfg.FromHeight,
		utilsWrapper:   &utils.Utils{},
		store:          store,
		config:         cfg,
		chainId:        hexutil.MustDecodeBig(cfg.ChainId),
		prepareJobChan: prepareJobChan,
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

func (e *EthereumListener) IsDisabled() bool {
	return e.config.Disabled
}

func (e *EthereumListener) SetInitHeight(height uint64) {
	e.fromHeight = height
}

func (e *EthereumListener) GetInitHeight() uint64 {
	return e.fromHeight
}

func (e *EthereumListener) GetCurrentBlock() types.IBlock {
	if _, ok := e.currentBlock.Load().(types.IBlock); !ok {
		var (
			block types.IBlock
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
	return e.currentBlock.Load().(types.IBlock)
}

func (e *EthereumListener) GetProcessedBlock() (types.IBlock, error) {
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

func (e *EthereumListener) GetLatestBlock() (types.IBlock, error) {
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

func (e *EthereumListener) GetSubscriptions() map[string]*types.Subscribe {
	return e.config.Subscriptions
}

func (e *EthereumListener) UpdateCurrentBlock(block types.IBlock) error {
	if e.GetCurrentBlock().GetHash().Hex() != block.GetHash().Hex() {
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
	return e.store.GetProcessedBlockStore().Save(hexutil.EncodeBig(chainId), int64(e.GetCurrentBlock().GetHeight()))
}

func (e *EthereumListener) SaveTransactionsToDB(txs []types.ITransaction) error {
	return nil
}

func (e *EthereumListener) GetEthClient() utils.EthClient {
	return e.client
}

func (e *EthereumListener) GetListenHandleJob(subscriptionName string, tx types.ITransaction, eventId string, data []byte) types.IJob {
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
		if hexutil.Encode(event.ID.Bytes()) != eventId {
			return nil
		}
	} else {
		return nil
	}
	return NewEthListenJob(types.ListenHandler, e, subscriptionName, tx, data)
}

func (e *EthereumListener) SendCallbackJobs(listeners map[string]types.IListener, subscriptionName string, tx types.ITransaction, inputData []byte) {
	log.Info("[EthereumListener][SendCallbackJobs] Start", "subscriptionName", subscriptionName, "listeners", len(listeners))
	chainId, err := e.GetChainID()
	if err != nil {
		return
	}
	subscription, ok := e.GetSubscriptions()[subscriptionName]
	if !ok {
		log.Info("[EthereumListener][SendCallbackJobs] cannot find subscription", "subscriptionName", subscriptionName)
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

func (e *EthereumListener) SendTransactionCheckerJob(chainId *big.Int, ids []int, tx *ethtypes.Transaction) {
	log.Info("[EthereumListener][SendTransactionCheckerJob] Start", "tx", tx.Hash().Hex())
	job := NewEthCheckTransactionStatusJob(e, ids, tx, chainId, e.utilsWrapper)
	e.prepareJobChan <- job
}

func (e *EthereumListener) GetBlock(height uint64) (types.IBlock, error) {
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, e.chainId, block, false)
}

func (e *EthereumListener) GetBlockWithLogs(height uint64) (types.IBlock, error) {
	block, err := e.client.BlockByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.client, e.chainId, block, true)
}

func (e *EthereumListener) GetReceipt(txHash common.Hash) (*ethtypes.Receipt, error) {
	return e.client.TransactionReceipt(e.ctx, txHash)
}

func (e *EthereumListener) NewJobFromDB(job *models.Job) (types.IJob, error) {
	chainId, err := hexutil.DecodeBig(job.FromChainId)
	if err != nil {
		return nil, err
	}
	// get transaction from hash
	tx, _, err := e.client.TransactionByHash(e.ctx, common.HexToHash(job.Transaction))
	if err != nil {
		return nil, err
	}
	transaction, err := NewEthTransaction(e.chainId, tx)
	if err != nil {
		return nil, err
	}

	switch job.Type {
	case types.ListenHandler:
		return &EthListenJob{
			&BaseJob{
				jobType:          types.ListenHandler,
				retryCount:       job.RetryCount,
				maxTry:           20,
				nextTry:          time.Now().Unix() + int64(job.RetryCount*5),
				backOff:          5,
				data:             common.Hex2Bytes(job.Data),
				tx:               transaction,
				subscriptionName: job.SubscriptionName,
				listener:         e,
				utilsWrapper:     e.utilsWrapper,
				fromChainID:      chainId,
				id:               int32(job.ID),
			},
		}, nil
	case types.CallbackHandler:
		if job.Method != "" {
			return nil, nil
		}
		return &EthCallbackJob{
			BaseJob: &BaseJob{
				utilsWrapper: e.utilsWrapper,
				jobType:      types.CallbackHandler,
				retryCount:   job.RetryCount,
				maxTry:       20,
				nextTry:      time.Now().Unix() + int64(job.RetryCount*5),
				backOff:      5,
				data:         common.Hex2Bytes(job.Data),
				tx:           transaction,
				listener:     e,
				fromChainID:  chainId,
				id:           int32(job.ID),
			},
			method: job.Method,
		}, nil
	case types.TransactionChecker:
		var ids []int
		if err := json.Unmarshal(common.Hex2Bytes(job.Data), &ids); err != nil {
			return nil, err
		}
		return &EthCheckTransactionStatusJob{
			BaseJob: &BaseJob{
				utilsWrapper: e.utilsWrapper,
				id:           int32(job.ID),
				jobType:      types.TransactionChecker,
				retryCount:   job.RetryCount,
				maxTry:       20,
				nextTry:      time.Now().Unix() + int64(job.RetryCount*5),
				backOff:      5,
				data:         common.Hex2Bytes(job.Data),
				tx:           transaction,
				listener:     e,
				fromChainID:  chainId,
			},
			ids: ids,
		}, nil
	}
	return nil, errors.New("jobType does not match")
}

func (e *EthereumListener) Close() {
	e.client.Close()
	e.cancelCtx()
}
