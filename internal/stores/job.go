package stores

import (
	"errors"
	"fmt"
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

func (j *JobStore) hasJob(hash string, jobType int) bool {
	var job = &models.Job{}
	j.Model(&models.Job{}).Select("id").Where("transaction = ? AND type = ?", hash, jobType).First(job)
	return job.ID > 0
}

func (j *JobStore) Save(job *models.Job) error {
	if j.hasJob(job.Transaction, job.Type) {
		return errors.New("job is already existed")
	}
	return j.Create(job).Error
}

func (j *JobStore) Update(job *models.Job) error {
	return j.Model(models.Job{}).Where("id = ?", job.ID).Updates(job).Error
}

func (j *JobStore) GetPendingJobs() ([]*models.Job, error) {
	// query all pending jobs
	var jobs []*models.Job
	err := j.Model(&models.Job{}).Where("status = ?", types.STATUS_PENDING).
		Order(fmt.Sprintf("created_at + POWER(2, retry_count) * 10 ASC")).Find(&jobs).Error
	return jobs, err
}

func (j *JobStore) DeleteJobs(status []string, createdAt uint64) error {
	return j.Where("status in ? AND created_at <= ?", status, createdAt).Delete(&models.Job{}).Error
}

func (j *JobStore) Count() int64 {
	var count int64
	j.Model(&models.Job{}).Select("id").Count(&count)
	return count
}
