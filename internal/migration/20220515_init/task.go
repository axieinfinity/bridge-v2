package models

import (
	"gorm.io/gorm"
)

type Task struct {
	ID                int    `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	ChainId           string `json:"chainId" gorm:"column:chain_id;index:idx_job_chain_id;not null"`
	FromChainId       string `json:"fromChainId" gorm:"column:from_chain_id;not null"`
	Type              string `json:"type" gorm:"column:task_type;not null"`
	Data              string `json:"data" gorm:"column:data;not null"`
	Retries           int    `json:"retries" gorm:"column:retries;not null"`
	Status            string `json:"status" gorm:"column:status;not null"`
	LastError         string `json:"lastError" gorm:"column:last_error"`
	TransactionHash   string `json:"transactionHash" gorm:"transaction_hash;index:idx_task_transaction_hash;not null"`
	FromTransaction   string `json:"fromTransaction" gorm:"from_transaction;index:idx_task_from_transaction;not null"`
	TransactionStatus int    `json:"transactionStatus" gorm:"transaction_status;index:idx_task_transaction_status;not null"`
	CreatedAt         int64  `json:"createdAt" gorm:"column:created_at;type:bigint;index:idx_task_created_at;not null"`
}

func (m Task) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (m Task) TableName() string {
	return "task"
}
