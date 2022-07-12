package models

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20220708",
		Migrate: func(tx *gorm.DB) error {
			type Result struct {
				ChainId string `gorm:"column:chain_id"`
				Block   int64  `gorm:"column:block"`
			}
			var results []*Result
			// get latest data from processedBlock
			tx.Raw("select chain_id, max(block) as block from processed_block group by chain_id").Scan(&results)

			// remove all data
			if err := tx.Exec("drop table processed_block").Error; err != nil {
				return err
			}
			// update 2 data to new table
			if err := tx.AutoMigrate(&ProcessedBlock{}); err != nil {
				return err
			}

			// insert result back to processedBlock with value - 500
			// to make the bridge re-process past block,
			// because the previous version of this bridge could not handle graceful shutdown nicely
			for _, result := range results {
				if err := tx.Create(&ProcessedBlock{
					ChainId: result.ChainId,
					Block:   result.Block - 500,
				}).Error; err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return nil
		},
	}
}
