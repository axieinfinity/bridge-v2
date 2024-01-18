package listener

import (
	"math/big"
	"time"

	"github.com/axieinfinity/bridge-v2/stats"

	bridgeCore "github.com/axieinfinity/bridge-core"
	bridgeCoreModels "github.com/axieinfinity/bridge-core/models"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type EthBlock struct {
	block *ethtypes.Block
}

func NewEthBlock(block *ethtypes.Block) *EthBlock {
	return &EthBlock{
		block: block,
	}
}

func (b *EthBlock) GetHash() common.Hash { return b.block.Hash() }
func (b *EthBlock) GetHeight() uint64    { return b.block.NumberU64() }

func (b *EthBlock) GetTransactions() []bridgeCore.Transaction {
	return nil
}

func (b *EthBlock) GetLogs() []bridgeCore.Log {
	return nil
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

type EthListenJob struct {
	*bridgeCore.BaseJob
}

func NewEthListenJob(jobType int, listener bridgeCore.Listener, subscriptionName string, tx bridgeCore.Transaction, data []byte) *EthListenJob {
	chainId, err := listener.GetChainID()
	if err != nil {
		return nil
	}
	job := &bridgeCoreModels.Job{
		ID:               0,
		SubscriptionName: subscriptionName,
		Type:             jobType,
		RetryCount:       0,
		Data:             common.Bytes2Hex(data),
		FromChainId:      hexutil.EncodeBig(chainId),
	}
	baseJob, err := bridgeCore.NewBaseJob(listener, job, tx)
	if err != nil {
		return nil
	}
	return &EthListenJob{
		baseJob,
	}
}

func (e *EthListenJob) Process() ([]byte, error) {
	data, err := e.BaseJob.Process()
	stats.SendErrorToStats(e.GetListener(), err)
	return data, err
}

type EthCallbackJob struct {
	*bridgeCore.BaseJob
	result interface{}
	method string
}

func NewEthCallbackJob(listener bridgeCore.Listener, method string, tx bridgeCore.Transaction, data []byte, fromChainID *big.Int, helpers utils.Utils) *EthCallbackJob {
	if helpers == nil {
		helpers = utils.NewUtils()
	}
	job := &bridgeCoreModels.Job{
		ID:          0,
		Type:        bridgeCore.CallbackHandler,
		RetryCount:  0,
		Data:        common.Bytes2Hex(data),
		FromChainId: hexutil.EncodeBig(fromChainID),
	}
	baseJob, err := bridgeCore.NewBaseJob(listener, job, tx)
	if err != nil {
		return nil
	}

	return &EthCallbackJob{
		BaseJob: baseJob,
		method:  method,
	}
}

func (e *EthCallbackJob) Process() ([]byte, error) {
	log.Info("[EthCallbackJob] Start Process", "method", e.method, "jobId", e.GetID())
	val, err := e.Utils().Invoke(e.GetListener(), e.method, e.FromChainID(), e.GetTransaction(), e.GetData())
	if err != nil {
		stats.SendErrorToStats(e.GetListener(), err)
		return nil, err
	}
	invokeErr, ok := val.Interface().(error)
	if ok {
		stats.SendErrorToStats(e.GetListener(), invokeErr)
		return nil, invokeErr
	}
	return nil, nil
}

func (e *EthCallbackJob) Update(status string) error {
	job := &bridgeCoreModels.Job{
		ID:               int(e.GetID()),
		Listener:         e.GetListener().GetName(),
		SubscriptionName: e.GetSubscriptionName(),
		Type:             e.GetType(),
		RetryCount:       e.GetRetryCount(),
		Status:           status,
		Data:             common.Bytes2Hex(e.GetData()),
		Transaction:      e.GetTransaction().GetHash().Hex(),
		CreatedAt:        time.Now().Unix(),
		FromChainId:      hexutil.EncodeBig(e.FromChainID()),
		Method:           e.method,
	}
	if err := e.GetListener().GetStore().GetJobStore().Update(job); err != nil {
		return err
	}
	return nil
}

func (e *EthCallbackJob) Save(status string) error {
	job := &bridgeCoreModels.Job{
		Listener:         e.GetListener().GetName(),
		SubscriptionName: e.GetSubscriptionName(),
		Type:             e.GetType(),
		RetryCount:       e.GetRetryCount(),
		Status:           status,
		Data:             common.Bytes2Hex(e.GetData()),
		Transaction:      e.GetTransaction().GetHash().Hex(),
		CreatedAt:        time.Now().Unix(),
		FromChainId:      hexutil.EncodeBig(e.FromChainID()),
		Method:           e.method,
	}
	if err := e.GetListener().GetStore().GetJobStore().Save(job); err != nil {
		return err
	}
	e.SetID(int32(job.ID))
	return nil
}
