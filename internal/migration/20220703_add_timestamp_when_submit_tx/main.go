package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

func Migrate() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220703",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(&Task{}); err != nil {
				return err
			}
			// update tx_created_at of all processing tasks to now
			columns := map[string]interface{}{
				"tx_created_at": time.Now().Unix(),
			}
			if err := tx.Model(&Task{}).Where("status = 'processing'").Updates(columns).Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&PreviousTask{})
		},
	}
}
