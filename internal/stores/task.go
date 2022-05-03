package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type TaskStore struct {
	db *gorm.DB
}

func NewTaskStore(db *gorm.DB) *TaskStore {
	return &TaskStore{db: db}
}

func (t *TaskStore) Save(task *models.Task) error                                    { return nil }
func (t *TaskStore) Update(task *models.Task) error                                  { return nil }
func (t *TaskStore) GetPendingTasks(chain string, limit int) ([]*models.Task, error) { return nil, nil }
