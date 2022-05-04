package stores

import (
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"gorm.io/gorm"
)

type JobStore struct {
	*gorm.DB
}

func NewJobStore(db *gorm.DB) *JobStore {
	return &JobStore{db}
}

func (j *JobStore) Save(job *models.Job) error {
	return j.Create(job).Error
}

func (j *JobStore) Update(job *models.Job) error {
	return j.Updates(job).Error
}

func (j *JobStore) GetPendingJobs() ([]*models.Job, error) {
	// query all pending jobs
	var jobs []*models.Job
	err := j.Model(&models.Job{}).Where("status = ?", types.STATUS_PENDING).Find(jobs).Error
	return jobs, err
}
