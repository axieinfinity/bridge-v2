package stores

import (
	"fmt"
	"time"

	"github.com/axieinfinity/bridge-v2/models"
	"gorm.io/gorm"
)

type taskStore struct {
	*gorm.DB
}

func NewTaskStore(db *gorm.DB) TaskStore {
	return &taskStore{db}
}

func (t *taskStore) Save(task *models.Task) error {
	return t.Create(task).Error
}

func (t *taskStore) Update(task *models.Task) error {
	columns := map[string]interface{}{
		"status":  task.Status,
		"retries": task.Retries,
	}
	if task.TransactionHash != "" {
		columns["transaction_hash"] = task.TransactionHash
	}
	if task.TxCreatedAt > 0 {
		columns["tx_created_at"] = task.TxCreatedAt
	}
	if task.LastError != "" {
		columns["last_error"] = task.LastError
	}
	return t.Model(&models.Task{}).Where("id = ?", task.ID).Updates(columns).Error
}

func (t *taskStore) GetTasks(chain, status string, limit, retrySeconds int, before int64, excludeIds []int) ([]*models.Task, error) {
	// query all tasks with status and chain id and tx created time must be before specified time
	// also apply exponential at created_time
	var tasks []*models.Task
	db := t.Model(&models.Task{}).Where("chain_id = ? AND status = ?", chain, status)
	db = db.Where("created_at + POWER(2, retries) * ? <= ?", retrySeconds, time.Now().Unix())
	if before > 0 {
		db = db.Where("tx_created_at <= ?", before)
	}
	if len(excludeIds) > 0 {
		db = db.Where("id not in ?", excludeIds)
	}
	err := db.Order(fmt.Sprintf("created_at + POWER(2, retries) * %d ASC", retrySeconds)).
		Limit(limit).Find(&tasks).Error
	return tasks, err
}

func (t *taskStore) UpdateTasksWithTransactionHash(txs []string, transactionStatus int, status string) error {
	return t.Model(&models.Task{}).Where("transaction_hash in ?", txs).Updates(map[string]interface{}{
		"status":             status,
		"transaction_status": transactionStatus,
	}).Error
}

func (t *taskStore) ResetTo(txs []string, status string) error {
	columns := map[string]interface{}{
		"status":  status,
		"retries": gorm.Expr("retries + 1"),
	}
	return t.Model(&models.Task{}).Where("transaction_hash in ?", txs).Updates(columns).Error
}

func (t *taskStore) DeleteTasks(status []string, fromTime uint64) error {
	return t.Where("status in ? AND created_at <= ?", status, fromTime).Delete(&models.Task{}).Error
}

func (t *taskStore) Count() int64 {
	var count int64
	t.Model(&models.Task{}).Select("id").Count(&count)
	return count
}
