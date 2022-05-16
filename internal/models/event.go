package models

import "gorm.io/gorm"

type Event struct {
	ID              int    `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	EventName       string `json:"eventName" gorm:"column:event_name;index:idx_event_event_name;not null"`
	TransactionHash string `json:"transactionHash" gorm:"column:transaction_hash;index:idx_event_transaction_hash;not null"`
	FromChainId     string `json:"fromChainId" gorm:"column:from_chain_id;not null"`
	CreatedAt       int64  `json:"created_at" gorm:"column:created_at;type:bigint;index:idx_job_created_at;not null"`
}

func (e Event) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (e Event) TableName() string {
	return "event"
}