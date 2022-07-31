package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type TaskStore interface {
	Save(task *models.Task) error
	Update(task *models.Task) error
	GetTasks(chain, status string, limit, retrySeconds int, before int64, excludeIds []int) ([]*models.Task, error)
	UpdateTasksWithTransactionHash(txs []string, transactionStatus int, status string) error
	DeleteTasks([]string, uint64) error
	Count() int64
	ResetTo(ids []string, status string) error
}

type DepositStore interface {
	Save(deposit *models.Deposit) error
}

type ProcessedReceiptStore interface {
	Save(taskId int, receiptId int64) error
}

type WithdrawalStore interface {
	Save(withdraw *models.Withdrawal) error
	Update(withdraw *models.Withdrawal) error
	GetWithdrawalById(withdrawalId int64) (*models.Withdrawal, error)
}

type ListenHandlerStore interface {
	GetDepositStore() DepositStore
	GetWithdrawalStore() WithdrawalStore
	GetTaskStore() TaskStore
	GetProcessedReceiptStore() ProcessedReceiptStore
}

type listenHandlerStore struct {
	*gorm.DB

	DepositStore          DepositStore
	WithdrawalStore       WithdrawalStore
	TaskStore             TaskStore
	ProcessedReceiptStore ProcessedReceiptStore
}

func NewListenHandlerStore(db *gorm.DB) ListenHandlerStore {
	store := &listenHandlerStore{
		DB: db,

		TaskStore:             NewTaskStore(db),
		DepositStore:          NewDepositStore(db),
		WithdrawalStore:       NewWithdrawalStore(db),
		ProcessedReceiptStore: NewProcessedReceiptStore(db),
	}
	return store
}

func (m *listenHandlerStore) RelationalDatabaseCheck() error {
	return m.Raw("SELECT 1").Error
}

func (m *listenHandlerStore) GetDB() *gorm.DB {
	return m.DB
}

func (m *listenHandlerStore) GetDepositStore() DepositStore {
	return m.DepositStore
}

func (m *listenHandlerStore) GetWithdrawalStore() WithdrawalStore {
	return m.WithdrawalStore
}

func (m *listenHandlerStore) GetTaskStore() TaskStore {
	return m.TaskStore
}

func (m *listenHandlerStore) GetProcessedReceiptStore() ProcessedReceiptStore {
	return m.ProcessedReceiptStore
}
