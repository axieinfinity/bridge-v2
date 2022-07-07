package models

import (
	"gorm.io/gorm"
)

type ProcessedReceipt struct {
	ID        int   `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	TaskId    int   `json:"taskId" gorm:"column:task_id;index:idx_processedReceipt_taskId;not null"`
	ReceiptId int64 `json:"receiptId" gorm:"column:receipt_id;index:idx_processedReceipt_receiptId;not null"`
}

func (m ProcessedReceipt) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (m ProcessedReceipt) TableName() string {
	return "processed_receipt"
}
