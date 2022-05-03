package models

import (
	"gorm.io/gorm"
)

type Task struct {
	ID        int    `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	Chain     string `json:"chain" gorm:"column:chain;index:idx_job_chain_name;not null"`
	Type      string `json:"type" gorm:"column:task_type;not null"`
	Data      string `json:"data" gorm:"column:data;not null"`
	Retries   int    `json:"retries" gorm:"column:retries;not null"`
	Status    string `json:"status" gorm:"column:status;not null"`
	LastError string `json:"lastError" gorm:"column:last_error"`
	CreatedAt int64  `json:"createdAt" gorm:"column:created_at;type:bigint;index:idx_task_created_at;not null"`
}

func (m Task) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (m Task) TableName() string {
	return "task"
}
