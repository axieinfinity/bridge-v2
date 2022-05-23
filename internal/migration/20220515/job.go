package models

import (
	"gorm.io/gorm"
)

type Job struct {
	ID               int    `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	Listener         string `json:"listener" gorm:"column:listener;index:idx_job_listener_name;not null"`
	SubscriptionName string `json:"subscriptionName" gorm:"column:subscription_name;not null"`
	Type             int    `json:"type" gorm:"column:type;not null"`
	RetryCount       int    `json:"retryCount" gorm:"column:retry_count;not null"`
	Status           string `json:"status" gorm:"column:status;not null"`
	Data             string `json:"data" gorm:"column:data;not null"`
	Transaction      string `json:"transaction" gorm:"column:transaction;index:idx_job_transaction;not null"`
	CreatedAt        int64  `json:"created_at" gorm:"column:created_at;type:bigint;index:idx_job_created_at;not null"`
	FromChainId      string `json:"fromChainId" gorm:"column:from_chain_id;not null"`

	// Method is used to execute function in `callback` job
	Method string `json:"method" gorm:"column:method;not null"`
}

func (m Job) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (m Job) TableName() string {
	return "job"
}
