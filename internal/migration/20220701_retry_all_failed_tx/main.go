package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220701",
		Migrate: func(tx *gorm.DB) error {
			columns := map[string]interface{}{
				"status":     "pending",
				"retries":    0,
				"last_error": "",
			}
			if err := tx.Model(&Task{}).Where("status = 'failed'").Updates(columns).Error; err != nil {
				return err
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
