package models

import (
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

func Migrate(cfg *types.Config) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220703",
		Migrate: func(tx *gorm.DB) error {
			if err := tx.AutoMigrate(&Task{}); err != nil {
				return err
			}
			// update tx_created_at of all processing tasks to `block_time * safe_block_range` seconds before now
			ronListener := cfg.Listeners["Ronin"]
			sec := uint64((ronListener.LoadInterval * time.Second).Seconds()) * ronListener.SafeBlockRange
			now := time.Now().Unix()
			columns := map[string]interface{}{
				"tx_created_at": now - int64(sec),
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
