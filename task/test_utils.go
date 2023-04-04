package task

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	internal "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RoninChainId    = "0x539"
	EthereumChainId = "0x5"
)

func newMockDB() (*gorm.DB, *sql.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("an error '%s' was not expected when opening a stub database connection", err))
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, nil, errors.New(fmt.Sprintf("an error '%s' was not expected when opening a gorm", err))
	}
	return gormDB, sqlDB, mock, nil
}

func newMockConfig(privateKey string, contracts map[string]string) *internal.LsConfig {
	return &internal.LsConfig{
		Secret: &internal.Secret{
			BridgeOperator: &utils.SignMethodConfig{
				PlainPrivateKey: privateKey,
				KmsConfig:       nil,
			},
			Relayer: nil,
		},
		Contracts: contracts,
		ChainId:   RoninChainId,
	}
}
