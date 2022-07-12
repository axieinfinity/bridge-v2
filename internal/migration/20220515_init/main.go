package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220515",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Deposit{}, &Event{}, &Job{}, &ProcessedBlock{}, &Task{}, &Withdrawal{})
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
