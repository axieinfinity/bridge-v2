package models

import (
	"gorm.io/gorm"
)

type Deposit struct {
	ID             int    `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	Owner          string `json:"owner" gorm:"column:owner;index:idx_deposit_owner;not null"`
	TokenAddress   string `json:"tokenAddress" gorm:"column:token_address;index:idx_deposit_token_address;not null"`
	TokenNumber    int    `json:"tokenNumber" gorm:"column:token_number;not null"`
	BlockNumber    int64  `json:"blockNumber" gorm:"column:block_number;index:idx_deposit_block_number;not null"`
	SidechainToken string `json:"sidechainToken" gorm:"column:sidechain_token;not null"`
	Standard       int    `json:"standard" gorm:"column:standard;not null"`
}

func (m Deposit) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (m Deposit) TableName() string {
	return "deposit"
}
