package types

import (
	"context"
	"crypto/ecdsa"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"math/big"
	"time"
)

const (
	TxEvent = iota
	LogEvent
)

const (
	ListenHandler = iota
	CallbackHandler
)

const (
	VoteStatusPending = iota
	VoteStatusApproved
	VoteStatusExecuted
)

const (
	ACK_WITHDREW_TASK = "acknowledgeWithdrew"
	DEPOSIT_TASK      = "deposit"
	WITHDRAWAL_TASK   = "withdrawal"

	STATUS_PENDING    = "pending"
	STATUS_FAILED     = "failed"
	STATUS_PROCESSING = "processing"
	STATUS_DONE       = "done"

	GATEWAY_CONTRACT     = "Gateway"
	ETH_GATEWAY_CONTRACT = "EthGateway"
	BRIDGEADMIN_CONTRACT = "BridgeAdmin"
)

type IListener interface {
	GetName() string
	GetStore() IMainStore
	Config() *LsConfig

	Period() time.Duration
	GetSafeBlockRange() uint64
	GetCurrentBlock() IBlock
	GetLatestBlock() (IBlock, error)
	GetLatestBlockHeight() (uint64, error)
	GetBlock(height uint64) (IBlock, error)
	GetBlockWithLogs(height uint64) (IBlock, error)
	GetChainID() (*big.Int, error)
	GetReceipt(common.Hash) (*types.Receipt, error)
	Context() context.Context

	GetSubscriptions() map[string]*Subscribe

	UpdateCurrentBlock(block IBlock) error

	SaveCurrentBlockToDB() error
	SaveTransactionsToDB(txs []ITransaction) error

	GetListenHandleJob(subscriptionName string, tx ITransaction, eventId string, data []byte) IJob
	SendCallbackJobs(listeners map[string]IListener, subscriptionName string, tx ITransaction, inputData []byte)

	NewJobFromDB(job *models.Job) (IJob, error)

	Start()
	Close()

	IsDisabled() bool
	SetInitHeight(uint64)
	GetInitHeight() uint64

	GetEthClient() utils.EthClient
	GetTask() ITask

	IsUpTodate() bool
}

type IEthListener interface {
	IListener
	GetValidatorKey() *ecdsa.PrivateKey
	GetRelayerKey() *ecdsa.PrivateKey
}

type ITransaction interface {
	GetHash() common.Hash
	GetFromAddress() string
	GetToAddress() string
	GetData() []byte
	GetValue() *big.Int
}

type ILog interface {
	GetContractAddress() string
	GetTopics() []string
	GetData() []byte
	GetIndex() uint
	GetTxIndex() uint
	GetTransactionHash() string
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
	GetLogs() []ILog
	GetTimestamp() uint64
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

	Process() ([]byte, error)
	Hash() common.Hash

	IncreaseRetryCount()
	UpdateNextTry(int64)

	GetListener() IListener
	GetSubscriptionName() string
	GetTransaction() ITransaction

	FromChainID() *big.Int

	Save() error
	Update(string) error
}

type ITask interface {
	Start()
	Close()
	GetListener() IListener
	SetLimitQuery(limit int)
}

type IJobStore interface {
	Save(job *models.Job) error
	Update(job *models.Job) error
	GetPendingJobs() ([]*models.Job, error)
	DeleteJobs([]string, uint64) error
	Count() int64
}

type IProcessedBlockStore interface {
	Save(chainId string, height int64) error
	GetLatestBlock(chainId string) (int64, error)
}

type ITaskStore interface {
	Save(task *models.Task) error
	Update(task *models.Task) error
	GetTasks(chain, status string, limit, retrySeconds int, before int64, excludeIds []int) ([]*models.Task, error)
	UpdateTasksWithTransactionHash(txs []string, transactionStatus int, status string) error
	DeleteTasks([]string, uint64) error
	Count() int64
	ResetTo(ids []string, status string) error
}

type IDepositStore interface {
	Save(deposit *models.Deposit) error
}

type IWithdrawalStore interface {
	Save(withdraw *models.Withdrawal) error
	Update(withdraw *models.Withdrawal) error
	GetWithdrawalById(withdrawalId int64) (*models.Withdrawal, error)
}

type IEventStore interface {
	Save(event *models.Event) error
	DeleteEvents(uint64) error
	Count() int64
}

type IMainStore interface {
	GetDB() *gorm.DB
	GetDepositStore() IDepositStore
	GetWithdrawalStore() IWithdrawalStore
	GetProcessedReceiptStore() IProcessedReceiptStore
	GetTaskStore() ITaskStore
	GetJobStore() IJobStore
	GetProcessedBlockStore() IProcessedBlockStore
	GetEventStore() IEventStore
}

type IProcessedReceiptStore interface {
	Save(taskId int, receiptId int64) error
}

type Config struct {
	Listeners       map[string]*LsConfig `json:"listeners"`
	NumberOfWorkers int                  `json:"numberOfWorkers"`
	DB              *Database            `json:"database"`
	Cleaner         Cleaner              `json:"cleaner"`

	// this field is used for testing purpose
	Testing bool
}

type Database struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`
	Port     int    `json:"port"`

	ConnMaxLifetime int `json:"connMaxLifeTime"`
	MaxIdleConns    int `json:"maxIdleConns"`
	MaxOpenConns    int `json:"maxOpenConns"`
}

type Cleaner map[string]*CleanerConfig

type CleanerConfig struct {
	Cron           string `json:"cron"`
	RemoveAfter    uint64 `json:"removeAfter"`
	SkipIfLessThan uint64 `json:"SkipIfLessThan"`
	Description    string `json:"description"`
}

type LsConfig struct {
	ChainId        string        `json:"chainId"`
	Name           string        `json:"-"`
	RpcUrl         string        `json:"rpcUrl"`
	LoadInterval   time.Duration `json:"blockTime"`
	SafeBlockRange uint64        `json:"safeBlockRange"`
	FromHeight     uint64        `json:"fromHeight"`
	TaskInterval   time.Duration `json:"taskInterval"`
	Disabled       bool          `json:"disabled"`

	// TODO: apply more ways to get privatekey. such as: PLAINTEXT, KMS, etc.
	Secret                 *Secret               `json:"secret"`
	Subscriptions          map[string]*Subscribe `json:"subscriptions"`
	TransactionCheckPeriod time.Duration         `json:"transactionCheckPeriod"`
	Contracts              map[string]string     `json:"contracts"`
	ProcessWithinBlocks    uint64                `json:"processWithinBlocks"`

	MaxTasksQuery int `json:"maxTasksQuery"`
	MinTasksQuery int `json:"minTasksQuery"`

	// GetLogsBatchSize is used at batch size when calling processBatchLogs
	GetLogsBatchSize int `json:"getLogsBatchSize"`

	// MaxProcessingTasks is used to specify max processing tasks allowed while processing tasks
	// if number of tasks reaches this number, it waits until this number decrease
	MaxProcessingTasks int `json:"maxProcessingTasks"`
}

type Secret struct {
	Validator string `json:"validator"`
	Relayer   string `json:"relayer"`
}

type Subscribe struct {
	From string `json:"from"`
	To   string `json:"to"`

	// Type can be either TxEvent or LogEvent
	Type int `json:"type"`

	Handler   *Handler          `json:"handler"`
	CallBacks map[string]string `json:"callbacks"`
}

type Handler struct {
	// ABIPath specifies path to abi file
	ABIPath string `json:"abi"`

	// Name is method/event name
	Name string `json:"name"`

	// ContractAddress is used in callback case
	ContractAddress string `json:"contractAddress"`

	// Listener who triggers callback event
	Listener string `json:"listener"`

	ABI *abi.ABI `json:"-"`

	// HandleMethod is used when processing listened job, do nothing if it is empty
	HandleMethod string `json:"handleMethod"`
}
