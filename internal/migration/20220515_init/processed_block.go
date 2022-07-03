package models

import (
	"gorm.io/gorm"
)

type ProcessedBlock struct {
	ID      int    `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	ChainId string `json:"chainId" gorm:"column:chain_id;not null"`
	Block   int64  `json:"block" gorm:"column:block;not null"`
}

func (m ProcessedBlock) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (m ProcessedBlock) TableName() string {
	return "processed_block"
}
