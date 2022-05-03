package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"gorm.io/gorm"
)

type JobStore struct {
	db *gorm.DB
}

func NewJobStore(db *gorm.DB) *JobStore {
	return &JobStore{db: db}
}

func (j *JobStore) Save(job *models.Job) error             { return nil }
func (j *JobStore) Update(job *models.Job) error           { return nil }
func (j *JobStore) GetPendingJobs() ([]*models.Job, error) { return nil, nil }
