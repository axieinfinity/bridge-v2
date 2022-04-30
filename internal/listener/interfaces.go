package listener

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type IListener interface {
	GetName() string

	Period() time.Duration
	GetSafeBlockRange() uint64
	GetCurrentBlock() IBlock
	GetLatestBlock() (IBlock, error)
	GetBlock(height uint64) (IBlock, error)
	Context() context.Context

	GetSubscriptions() map[string]*Subscribe

	UpdateCurrentBlock(block IBlock) error

	SaveCurrentBlockToDB() error
	SaveTransactionsToDB(txs []ITransaction) error

	GetListenHandleJob(subscriptionName string, tx ITransaction, data []byte) IJob
	SendCallbackJobs(listeners map[string]IListener, subscriptionName string, tx ITransaction, result interface{}, jobChan chan<- IJob)

	Start()
	Close()
}

type ITransaction interface {
	GetHash() common.Hash
	GetFromAddress() string
	GetToAddress() string
	GetData() []byte
	GetValue() *big.Int
	GetStatus() bool
	GetReceipt() IReceipt
}

type ILog interface {
	GetContractAddress() string
	GetTopics() []string
	GetData() []byte
	GetIndex() uint
}

type IReceipt interface {
	GetTransaction() ITransaction
	GetStatus() bool
	GetLogs() []ILog
}

type IBlock interface {
	GetHash() common.Hash
	GetHeight() uint64
	GetTransactions() []ITransaction
	GetReceipts() []IReceipt
}

type IJob interface {
	GetID() int32
	GetType() int
	GetRetryCount() int
	GetNextTry() int64
	GetMaxTry() int
	GetData() []byte
	GetValue() *big.Int
	GetBackOff() int

	Process() (map[string]interface{}, error)
	Hash() common.Hash

	IncreaseRetryCount()
	UpdateNextTry(int64)

	GetListener() IListener
	GetSubscriptionName() string
	GetTransaction() ITransaction
}
