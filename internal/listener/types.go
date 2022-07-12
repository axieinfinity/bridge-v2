package listener

import (
	"context"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

type EthBlock struct {
	block *ethtypes.Block
	txs   []types.ITransaction
	logs  []types.ILog
}

func NewEthBlock(client utils.EthClient, chainId *big.Int, block *ethtypes.Block, getLogs bool) (*EthBlock, error) {
	ethBlock := &EthBlock{
		block: block,
	}
	// convert txs into ILog
	for _, tx := range block.Transactions() {
		transaction, err := NewEthTransaction(chainId, tx)
		if err != nil {
			log.Error("[NewEthBlock] error while init new Eth Transaction", "err", err, "tx", tx.Hash().Hex())
			return nil, err
		}
		ethBlock.txs = append(ethBlock.txs, transaction)
	}
	if getLogs {
		log.Info("Getting logs from block hash", "block", block.NumberU64(), "hash", block.Hash().Hex())
		blockHash := block.Hash()
		logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{BlockHash: &blockHash})
		if err != nil {
			log.Error("[NewEthBlock] error while getting logs", "err", err, "block", block.NumberU64(), "hash", block.Hash().Hex())
			return nil, err
		}
		// convert logs to ILog
		for _, l := range logs {
			ethLog := EthLog(l)
			ethBlock.logs = append(ethBlock.logs, &ethLog)
		}
	}
	log.Info("[NewEthBlock] Finish getting eth block", "block", ethBlock.block.NumberU64(), "txs", len(ethBlock.txs), "logs", len(ethBlock.logs))
	return ethBlock, nil
}

func (b *EthBlock) GetHash() common.Hash { return b.block.Hash() }
func (b *EthBlock) GetHeight() uint64    { return b.block.NumberU64() }

func (b *EthBlock) GetTransactions() []types.ITransaction {
	return b.txs
}

func (b *EthBlock) GetLogs() []types.ILog {
	return b.logs
}

func (b *EthBlock) GetTimestamp() uint64 {
	return b.block.Time()
}

type EthTransaction struct {
	chainId *big.Int
	sender  common.Address
	tx      *ethtypes.Transaction
}

func NewEthTransactionWithoutError(chainId *big.Int, tx *ethtypes.Transaction) *EthTransaction {
	ethTx := &EthTransaction{
		chainId: chainId,
		tx:      tx,
	}
	sender, err := ethtypes.LatestSignerForChainID(chainId).Sender(tx)
	if err == nil {
		ethTx.sender = sender
	}
	return ethTx
}

func NewEthTransaction(chainId *big.Int, tx *ethtypes.Transaction) (*EthTransaction, error) {
	sender, err := ethtypes.LatestSignerForChainID(chainId).Sender(tx)
	if err != nil {
		return nil, err
	}
	return &EthTransaction{
		chainId: chainId,
		sender:  sender,
		tx:      tx,
	}, nil
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

type EmptyTransaction struct {
	chainId  *big.Int
	hash     common.Hash
	from, to *common.Address
	data     []byte
}

func NewEmptyTransaction(chainId *big.Int, tx common.Hash, data []byte, from, to *common.Address) *EmptyTransaction {
	return &EmptyTransaction{
		chainId: chainId,
		hash:    tx,
		from:    from,
		to:      to,
		data:    data,
	}
}

func (b *EmptyTransaction) GetHash() common.Hash {
	return b.hash
}

func (b *EmptyTransaction) GetFromAddress() string {
	if b.from != nil {
		return b.from.Hex()
	}
	return ""
}
func (b *EmptyTransaction) GetToAddress() string {
	if b.to != nil {
		return b.to.Hex()
	}
	return ""
}

func (b *EmptyTransaction) GetData() []byte {
	return b.data
}

func (b *EmptyTransaction) GetValue() *big.Int {
	return nil
}

type EthLog ethtypes.Log

func (e *EthLog) GetContractAddress() string {
	return e.Address.Hex()
}

func (e *EthLog) GetTopics() (topics []string) {
	for _, topic := range e.Topics {
		topics = append(topics, topic.Hex())
	}
	return
}

func (e *EthLog) GetData() []byte {
	return e.Data
}

func (e *EthLog) GetIndex() uint {
	return e.Index
}

func (e *EthLog) GetTxIndex() uint {
	return e.TxIndex
}

func (e *EthLog) GetTransactionHash() string {
	return e.TxHash.Hex()
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

func (e *BaseJob) Process() ([]byte, error) {
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
	// save event data to database
	subscription, ok := e.listener.GetSubscriptions()[e.subscriptionName]
	if ok {
		if err := e.listener.GetStore().GetEventStore().Save(&models.Event{
			EventName:       subscription.Handler.Name,
			TransactionHash: e.tx.GetHash().Hex(),
			FromChainId:     hexutil.EncodeBig(e.FromChainID()),
			CreatedAt:       time.Now().Unix(),
		}); err != nil {
			log.Error("[EthListenJob][Process] error while storing event to database", "err", err)
		}
	}
	return e.data, nil
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
	val, err := e.utilsWrapper.Invoke(e.listener, e.method, e.fromChainID, e.tx, e.data)
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

func (e *EthCallbackJob) Save() error {
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
		Method:           e.method,
	}
	if err := e.listener.GetStore().GetJobStore().Save(job); err != nil {
		return err
	}
	e.id = int32(job.ID)
	return nil
}
