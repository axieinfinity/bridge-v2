package stores

import (
	"github.com/axieinfinity/bridge-v2/models"
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
	CountTasks(chain, status string) (int64, error)
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

type BridgeStore interface {
	GetDepositStore() DepositStore
	GetWithdrawalStore() WithdrawalStore
	GetTaskStore() TaskStore
	GetProcessedReceiptStore() ProcessedReceiptStore
}

type bridgeStore struct {
	*gorm.DB

	DepositStore          DepositStore
	WithdrawalStore       WithdrawalStore
	TaskStore             TaskStore
	ProcessedReceiptStore ProcessedReceiptStore
}

func NewBridgeStore(db *gorm.DB) BridgeStore {
	store := &bridgeStore{
		DB: db,

		TaskStore:             NewTaskStore(db),
		DepositStore:          NewDepositStore(db),
		WithdrawalStore:       NewWithdrawalStore(db),
		ProcessedReceiptStore: NewProcessedReceiptStore(db),
	}
	return store
}

func (m *bridgeStore) RelationalDatabaseCheck() error {
	return m.Raw("SELECT 1").Error
}

func (m *bridgeStore) GetDB() *gorm.DB {
	return m.DB
}

func (m *bridgeStore) GetDepositStore() DepositStore {
	return m.DepositStore
}

func (m *bridgeStore) GetWithdrawalStore() WithdrawalStore {
	return m.WithdrawalStore
}

func (m *bridgeStore) GetTaskStore() TaskStore {
	return m.TaskStore
}

func (m *bridgeStore) GetProcessedReceiptStore() ProcessedReceiptStore {
	return m.ProcessedReceiptStore
}
