package models

import (
	"gorm.io/gorm"
)

type ProcessedBlock struct {
	ChainId string `json:"id" gorm:"primary_key:true;column:id;"`
	Block   int64  `json:"block" gorm:"column:block;not null"`
}

func (m ProcessedBlock) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (m ProcessedBlock) TableName() string {
	return "processed_block"
}
