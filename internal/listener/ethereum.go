package listener

import (
	"context"
	"errors"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
	"sync/atomic"
	"time"
)

type EthereumListener struct {
	jobId          int32
	rpcUrl         string
	name           string
	period         time.Duration
	currentBlock   atomic.Value
	ctx            context.Context
	cancelCtx      context.CancelFunc
	subscriptions  map[string]*Subscribe
	safeBlockRange uint64
	fromHeight     uint64
	batches        sync.Map
}

func (c *Controller) InitEthereumListener(ctx context.Context, lsConfig *LsConfig) *EthereumListener {
	newCtx, cancelFunc := context.WithCancel(ctx)
	return &EthereumListener{
		name:          lsConfig.Name,
		period:        lsConfig.LoadInterval,
		currentBlock:  atomic.Value{},
		ctx:           newCtx,
		cancelCtx:     cancelFunc,
		subscriptions: lsConfig.Subscriptions,
		fromHeight:    lsConfig.FromHeight,
	}
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

func (e *EthereumListener) GetCurrentBlock() IBlock {
	if e.currentBlock.Load() != nil {
		var (
			block IBlock
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
	return e.currentBlock.Load().(IBlock)
}

func (e *EthereumListener) GetLatestBlock() (IBlock, error) {
	client, err := ethclient.Dial(e.rpcUrl)
	if err != nil {
		return nil, err
	}
	// TODO: if apply ronin, lib change nil to -1
	block, err := client.BlockByNumber(e.ctx, nil)
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.rpcUrl, block)
}

func (e *EthereumListener) Context() context.Context {
	return e.ctx
}

func (e *EthereumListener) GetSubscriptions() map[string]*Subscribe {
	return e.subscriptions
}

func (e *EthereumListener) UpdateCurrentBlock(block IBlock) error {
	if e.GetCurrentBlock().GetHash().Hex() != block.GetHash().Hex() {
		e.currentBlock.Store(block)
	}
	return nil
}

func (e *EthereumListener) SaveCurrentBlockToDB() error {
	return nil
}

func (e *EthereumListener) SaveTransactionsToDB(txs []ITransaction) error {
	return nil
}

func (e *EthereumListener) GetListenHandleJob(subscriptionName string, tx ITransaction, data []byte) IJob {
	// validate if data contains subscribed name
	subscription, ok := e.GetSubscriptions()[subscriptionName]
	if !ok {
		return nil
	}
	handlerName := subscription.Handler.Name
	if subscription.Type == TxEvent {
		method, ok := subscription.Handler.ABI.Methods[handlerName]
		if !ok {
			return nil
		}
		if method.RawName != common.Bytes2Hex(data[0:4]) {
			return nil
		}
	} else if subscription.Type == LogEvent {
		event, ok := subscription.Handler.ABI.Events[handlerName]
		if !ok {
			return nil
		}
		if event.RawName != common.Bytes2Hex(data[0:4]) {
			return nil
		}
	} else {
		return nil
	}
	return nil
}

func (e *EthereumListener) SendCallbackJobs(listeners map[string]IListener, subscriptionName string, tx ITransaction, result interface{}, jobChan chan<- IJob) {
}

func (e *EthereumListener) GetBlock(height uint64) (IBlock, error) {
	client, err := ethclient.Dial(e.rpcUrl)
	if err != nil {
		return nil, err
	}
	block, err := client.BlockByNumber(e.ctx, big.NewInt(int64(height)))
	if err != nil {
		return nil, err
	}
	return NewEthBlock(e.rpcUrl, block)
}

func (e *EthereumListener) Close() {
	e.cancelCtx()
}

type BaseJob struct {
	utilsWrapper utils.IUtils

	id      int32
	jobType int

	retryCount int
	maxTry     int
	nextTry    int64
	backOff    int

	data []byte
	tx   ITransaction

	subscriptionName string
	listener         IListener
}

func (e *BaseJob) GetID() int32 {
	return e.id
}

func (e *BaseJob) GetType() int {
	return e.jobType
}

func (e *BaseJob) GetRetryCount() int {
	return e.retryCount
}

func (e *BaseJob) GetNextTry() int64 {
	return e.nextTry
}

func (e *BaseJob) GetMaxTry() int {
	return e.maxTry
}

func (e *BaseJob) GetData() []byte {
	return e.data
}

func (e *BaseJob) GetValue() *big.Int {
	return e.tx.GetValue()
}

func (e *BaseJob) GetBackOff() int {
	return e.backOff
}

func (e *BaseJob) Process() (map[string]interface{}, error) {
	return nil, nil
}

func (e *BaseJob) Hash() common.Hash {
	return common.BytesToHash([]byte(fmt.Sprintf("j-%d-%d-%d", e.id, e.retryCount, e.nextTry)))
}

func (e *BaseJob) IncreaseRetryCount() {
	e.retryCount++
}
func (e *BaseJob) UpdateNextTry(nextTry int64) {
	e.nextTry = nextTry
}

func (e *BaseJob) GetListener() IListener {
	return e.listener
}

func (e *BaseJob) GetSubscriptionName() string {
	return e.subscriptionName
}

func (e *BaseJob) GetTransaction() ITransaction {
	return e.tx
}

type EthListenJob struct {
	*BaseJob
}

func NewEthListenJob(id int32, jobType int, listener IListener, subscriptionName string, tx ITransaction, data []byte) *EthListenJob {
	return &EthListenJob{
		&BaseJob{
			id:               id,
			jobType:          jobType,
			retryCount:       0,
			maxTry:           20,
			nextTry:          0,
			backOff:          5,
			data:             data,
			tx:               tx,
			subscriptionName: subscriptionName,
			listener:         listener,
			utilsWrapper:     &utils.Utils{},
		},
	}
}

func (e *EthListenJob) Process() (map[string]interface{}, error) {
	// get subscription from name, unpack data based on data, method and
	subscription, ok := e.GetListener().GetSubscriptions()[e.subscriptionName]
	if !ok {
		return nil, errors.New("subscription not found")
	}
	var args abi.Arguments
	if subscription.Type == TxEvent {
		method, ok := subscription.Handler.ABI.Methods[subscription.Handler.Name]
		if !ok {
			return nil, errors.New("method not found")
		}
		args = method.Inputs
	} else if subscription.Type == LogEvent {
		event, ok := subscription.Handler.ABI.Events[subscription.Handler.Name]
		if !ok {
			return nil, errors.New("eventName not found")
		}
		args = event.Inputs
	} else {
		return nil, errors.New("invalid subscription's type")
	}
	// omit first 4 bytes which is method
	data := e.data[4:]
	// unpack data into map[string]interface{}
	rs := make(map[string]interface{})
	err := args.UnpackIntoMap(rs, data)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

type EthCallbackJob struct {
	*BaseJob
}

func (e *EthCallbackJob) Process() (map[string]interface{}, error) {
	return nil, nil
}
