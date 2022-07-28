package migration

import (
	"strings"

	models20220515 "github.com/axieinfinity/bridge-v2/internal/migration/20220515_init"
	models20220701 "github.com/axieinfinity/bridge-v2/internal/migration/20220701_retry_all_failed_tx"
	models20220703 "github.com/axieinfinity/bridge-v2/internal/migration/20220703_add_timestamp_when_submit_tx"
	models20220707 "github.com/axieinfinity/bridge-v2/internal/migration/20220707_add_processed_receipt_table"
	models20220708 "github.com/axieinfinity/bridge-v2/internal/migration/20220708_update_processed_block"
	models20220728 "github.com/axieinfinity/bridge-v2/internal/migration/20220728_add_withdrawal_created_at"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, cfg *types.Config) error {
	migrations := []*gormigrate.Migration{
		models20220515.Migrate(),
		models20220701.Migrate(),
		models20220703.Migrate(),
		models20220707.Migrate(),
		models20220708.Migrate(),
		models20220728.Migrate(),
	}
	migrate := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	for _, version := range migrations {
		if err := migrate.MigrateTo(version.ID); err != nil {
			if strings.Contains(err.Error(), "Duplicated") {
				continue
			}
			return err
		}
	}
	return nil
}
