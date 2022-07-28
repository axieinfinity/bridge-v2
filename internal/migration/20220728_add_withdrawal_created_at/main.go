package models

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate() *gormigrate.Migration {
	log.Printf("hi mom")

	return &gormigrate.Migration{
		ID: "20220728",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Withdrawal{})
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
