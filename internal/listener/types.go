package listener

import (
	"context"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type EthBlock struct {
	block    *ethtypes.Block
	txs      []types.ITransaction
	receipts []types.IReceipt
}

func NewEthBlock(client utils.EthClient, block *ethtypes.Block) (*EthBlock, error) {
	ethBlock := &EthBlock{
		block:    block,
		txs:      make([]types.ITransaction, len(block.Transactions())),
		receipts: make([]types.IReceipt, len(block.Transactions())),
	}
	for i, tx := range block.Transactions() {
		// get receipt from tx Hash
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return nil, err
		}
		ethTx, err := NewEthTransaction(client, tx)
		if err != nil {
			return nil, err
		}
		ethReceipt := &EthReceipt{
			receipt: receipt,
			tx:      ethTx,
		}
		ethBlock.txs[i] = ethTx
		ethBlock.receipts[i] = ethReceipt
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

func NewEthTransaction(client utils.EthClient, tx *ethtypes.Transaction) (*EthTransaction, error) {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
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

type EthListenJob struct {
	*BaseJob
}

func NewEthListenJob(id int32, jobType int, listener types.IListener, subscriptionName string, tx types.ITransaction, data []byte) *EthListenJob {
	chainId, err := listener.GetChainID()
	if err != nil {
		return nil
	}
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

func NewEthCallbackJob(id int32, listener types.IListener, method string, tx types.ITransaction, data []byte, fromChainID *big.Int, helpers utils.IUtils) *EthCallbackJob {
	if helpers == nil {
		helpers = &utils.Utils{}
	}
	return &EthCallbackJob{
		BaseJob: &BaseJob{
			utilsWrapper: helpers,
			id:           id,
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
