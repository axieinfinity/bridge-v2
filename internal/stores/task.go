package stores

import (
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type TaskStore struct {
	*gorm.DB
}

func NewTaskStore(db *gorm.DB) *TaskStore {
	return &TaskStore{db}
}

func (t *TaskStore) Save(task *models.Task) error {
	return t.Create(task).Error
}

func (t *TaskStore) Update(task *models.Task) error {
	return t.Updates(task).Error
}

func (t *TaskStore) GetTasks(chain, status string, limit, retrySeconds int) ([]*models.Task, error) {
	// query all tasks with status and chain id
	// also apply exponential at created_time
	var tasks []*models.Task
	err := t.Model(&models.Task{}).
		Where("chain_id = ? AND status = ?", chain, status).
		Order(fmt.Sprintf("created_at + POWER(2, retries) * %d ASC", retrySeconds)).
		Limit(limit).Find(&tasks).Error
	return tasks, err
}

func (t *TaskStore) UpdateTaskWithIds(ids []int, transactionStatus int, status string) error {
	return t.Model(&models.Task{}).Where("id in ?", ids).Updates(map[string]interface{}{"status": status, "transaction_status": transactionStatus}).Error
}

func (t *TaskStore) UpdateTasksWithTransactionHash(txs []string, transactionStatus int, status string) error {
	return t.Model(&models.Task{}).Where("transaction_hash in ?", txs).Updates(map[string]interface{}{
		"status":             status,
		"transaction_status": transactionStatus,
	}).Error
}

func (t *TaskStore) IncrementRetries(ids []int) error {
	return t.Model(&models.Task{}).Where("id in ?", ids).Update("retries", gorm.Expr("retries + 1")).Error
}
