package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
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

func (t *TaskStore) GetPendingTasks(chain string, limit int) ([]*models.Task, error) {
	// query all pending tasks which chain id
	// also apply exponential, created_time
	var tasks []*models.Task
	err := t.Model(&models.Task{}).
		Where("chain_id = ? AND status = ?", chain, types.STATUS_PENDING).
		Order("created_at + POWER(2, retries) * 10 ASC").
		Limit(limit).Find(&tasks).Error
	return tasks, err
}
