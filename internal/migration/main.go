package migration

import (
	models20220515 "github.com/axieinfinity/bridge-v2/internal/migration/20220515"
	models20220701 "github.com/axieinfinity/bridge-v2/internal/migration/20220701"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"strings"
)

func Migrate(db *gorm.DB, cfg *types.Config) error {
	migrations := []*gormigrate.Migration{
		models20220515.Migrate(),
		models20220701.Migrate(),
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
