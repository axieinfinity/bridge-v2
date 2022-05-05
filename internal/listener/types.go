package listener

import (
	"context"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"sync"
	"time"
)

type EthBlock struct {
	block    *ethtypes.Block
	txs      []types.ITransaction
	receipts []types.IReceipt
}

func NewEthBlock(client utils.EthClient, chainId *big.Int, block *ethtypes.Block) (*EthBlock, error) {
	ethBlock := &EthBlock{
		block:    block,
		txs:      make([]types.ITransaction, len(block.Transactions())),
		receipts: make([]types.IReceipt, len(block.Transactions())),
	}
	var (
		wg   sync.WaitGroup
		err  error
		lock sync.Mutex
	)
	log.Info("Getting transaction receipts", "block", block.NumberU64(), "txs", len(block.Transactions()))
	wg.Add(len(block.Transactions()))
	for i, tx := range block.Transactions() {
		go func(wg *sync.WaitGroup, index int, tx *ethtypes.Transaction) {
			log.Info("Getting receipt from rpc", "index", index, "tx", tx.Hash().Hex(), "block", block.NumberU64())
			// get receipt from tx Hash
			receipt, e := client.TransactionReceipt(context.Background(), tx.Hash())
			if e != nil {
				wg.Done()
				lock.Lock()
				defer lock.Unlock()
				err = e
				return
			}
			ethTx, e := NewEthTransaction(client, chainId, tx)
			if e != nil {
				wg.Done()
				lock.Lock()
				defer lock.Unlock()
				err = e
				return
			}
			ethReceipt := &EthReceipt{
				receipt: receipt,
				tx:      ethTx,
			}
			lock.Lock()
			defer lock.Unlock()
			ethBlock.txs[index] = ethTx
			ethBlock.receipts[index] = ethReceipt
			wg.Done()
		}(&wg, i, tx)
	}
	wg.Wait()
	if err != nil {
		return nil, err
	}
	return ethBlock, nil
}

func (b *EthBlock) GetHash() common.Hash { return b.block.Hash() }
func (b *EthBlock) GetHeight() uint64    { return b.block.NumberU64() }

func (b *EthBlock) GetTransactions() []types.ITransaction {
	return b.txs
}

func (b *EthBlock) GetReceipts() []types.IReceipt {
	if len(b.receipts) > 0 {
		return b.receipts
	}
	for _, tx := range b.GetTransactions() {
		b.receipts = append(b.receipts, tx.GetReceipt())
	}
	return b.receipts
}

type EthTransaction struct {
	chainId *big.Int
	sender  common.Address
	tx      *ethtypes.Transaction
	receipt types.IReceipt
}

func NewEthTransaction(client utils.EthClient, chainId *big.Int, tx *ethtypes.Transaction) (*EthTransaction, error) {
	sender, err := ethtypes.LatestSignerForChainID(chainId).Sender(tx)
	if err != nil {
		return nil, err
	}
	ethTx := &EthTransaction{
		chainId: chainId,
		sender:  sender,
		tx:      tx,
	}
	r, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, err
	}
	ethTx.receipt = &EthReceipt{
		receipt: r,
		tx:      ethTx,
	}
	return ethTx, nil
}

func (b *EthTransaction) GetHash() common.Hash {
	return b.tx.Hash()
}

func (b *EthTransaction) GetFromAddress() string {
	return b.sender.Hex()
}
func (b *EthTransaction) GetToAddress() string {
	if b.tx.To() != nil {
		return b.tx.To().Hex()
	}
	return ""
}

func (b *EthTransaction) GetData() []byte {
	return b.tx.Data()
}

func (b *EthTransaction) GetValue() *big.Int {
	return b.tx.Value()
}

func (b *EthTransaction) GetStatus() bool {
	return b.receipt.GetStatus()
}

func (b *EthTransaction) GetReceipt() types.IReceipt {
	return b.receipt
}

type EthReceipt struct {
	receipt *ethtypes.Receipt
	tx      *EthTransaction
}

func (r *EthReceipt) GetTransaction() types.ITransaction {
	return r.tx
}

func (r *EthReceipt) GetStatus() bool {
	return r.receipt.Status == 1
}

func (r *EthReceipt) GetLogs() (logs []types.ILog) {
	for _, l := range r.receipt.Logs {
		logs = append(logs, &EthLog{l})
	}
	return
}

type EthLog struct {
	l *ethtypes.Log
}

func (e *EthLog) GetContractAddress() string {
	return e.l.Address.Hex()
}

func (e *EthLog) GetTopics() (topics []string) {
	for _, topic := range e.l.Topics {
		topics = append(topics, topic.Hex())
	}
	return
}

func (e *EthLog) GetData() []byte {
	return e.l.Data
}

func (e *EthLog) GetIndex() uint {
	return e.l.Index
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
	tx   types.ITransaction

	subscriptionName string
	listener         types.IListener

	fromChainID *big.Int
}

func (e *BaseJob) FromChainID() *big.Int {
	return e.fromChainID
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

func (e *BaseJob) GetListener() types.IListener {
	return e.listener
}

func (e *BaseJob) GetSubscriptionName() string {
	return e.subscriptionName
}

func (e *BaseJob) GetTransaction() types.ITransaction {
	return e.tx
}

func (e *BaseJob) Save() error {
	job := &models.Job{
		Listener:         e.listener.GetName(),
		SubscriptionName: e.subscriptionName,
		Type:             e.jobType,
		RetryCount:       e.retryCount,
		Status:           types.STATUS_PENDING,
		Data:             common.Bytes2Hex(e.data),
		Transaction:      e.tx.GetHash().Hex(),
		CreatedAt:        time.Now().Unix(),
		FromChainId:      hexutil.EncodeBig(e.fromChainID),
	}
	if err := e.listener.GetStore().GetJobStore().Save(job); err != nil {
		return err
	}
	e.id = int32(job.ID)
	return nil
}

func (e *BaseJob) Update(status string) error {
	job := &models.Job{
		ID:               int(e.id),
		Listener:         e.listener.GetName(),
		SubscriptionName: e.subscriptionName,
		Type:             e.jobType,
		RetryCount:       e.retryCount,
		Status:           status,
		Data:             common.Bytes2Hex(e.data),
		Transaction:      e.tx.GetHash().Hex(),
		CreatedAt:        time.Now().Unix(),
		FromChainId:      hexutil.EncodeBig(e.fromChainID),
	}
	if err := e.listener.GetStore().GetJobStore().Update(job); err != nil {
		return err
	}
	return nil
}

type EthListenJob struct {
	*BaseJob
}

func NewEthListenJob(jobType int, listener types.IListener, subscriptionName string, tx types.ITransaction, data []byte) *EthListenJob {
	chainId, err := listener.GetChainID()
	if err != nil {
		return nil
	}
	return &EthListenJob{
		&BaseJob{
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
			fromChainID:      chainId,
		},
	}
}

func (e *EthListenJob) Process() ([]byte, error) {
	// TODO: implement handleMethod, if it is defined then process it and return its result.
	// omit first 4 bytes which is method
	return e.data[4:], nil
}

type EthCallbackJob struct {
	*BaseJob
	result interface{}
	method string
}

func NewEthCallbackJob(listener types.IListener, method string, tx types.ITransaction, data []byte, fromChainID *big.Int, helpers utils.IUtils) *EthCallbackJob {
	if helpers == nil {
		helpers = &utils.Utils{}
	}
	return &EthCallbackJob{
		BaseJob: &BaseJob{
			utilsWrapper: helpers,
			jobType:      types.CallbackHandler,
			retryCount:   0,
			maxTry:       20,
			nextTry:      0,
			backOff:      5,
			data:         data,
			tx:           tx,
			listener:     listener,
			fromChainID:  fromChainID,
		},
		method: method,
	}
}

func (e *EthCallbackJob) Process() ([]byte, error) {
	log.Info("[EthCallbackJob] Start Process", "method", e.method, "jobId", e.id)
	val, err := e.utilsWrapper.Invoke(e.listener, e.method, e.tx, e.data)
	if err != nil {
		return nil, err
	}
	invokeErr, ok := val.Interface().(error)
	if ok {
		return nil, invokeErr
	}
	return nil, nil
}

func (e *EthCallbackJob) Update(status string) error {
	job := &models.Job{
		ID:               int(e.id),
		Listener:         e.listener.GetName(),
		SubscriptionName: e.subscriptionName,
		Type:             e.jobType,
		RetryCount:       e.retryCount,
		Status:           status,
		Data:             common.Bytes2Hex(e.data),
		Transaction:      e.tx.GetHash().Hex(),
		CreatedAt:        time.Now().Unix(),
		FromChainId:      hexutil.EncodeBig(e.fromChainID),
		Method:           e.method,
	}
	if err := e.listener.GetStore().GetJobStore().Update(job); err != nil {
		return err
	}
	return nil
}
