package stores

import (
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal/models"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type MainStore struct {
	*gorm.DB

	DepositStore        types.IDepositStore
	WithdrawalStore     types.IWithdrawalStore
	JobStore            types.IJobStore
	TaskStore           types.ITaskStore
	ProcessedBlockStore types.IProcessedBlockStore
	EventStore          types.IEventStore
}

func NewMainStore(db *gorm.DB) *MainStore {
	cl := &MainStore{
		DB: db,

		JobStore:            NewJobStore(db),
		TaskStore:           NewTaskStore(db),
		ProcessedBlockStore: NewProcessedBlockStore(db),
		DepositStore:        NewDepositStore(db),
		WithdrawalStore:     NewWithdrawalStore(db),
		EventStore:          NewEventStore(db),
	}
	m := []interface{}{
		&models.Deposit{},
		&models.Job{},
		&models.ProcessedBlock{},
		&models.Task{},
		&models.Withdrawal{},
		&models.Event{},
	}
	if err := cl.AutoMigrate(m...); err != nil {
		panic(err)
	}
	return cl
}

func (m *MainStore) RelationalDatabaseCheck() error {
	return m.Raw("SELECT 1").Error
}

func (m *MainStore) GetDB() *gorm.DB {
	return m.DB
}

func (m *MainStore) GetDepositStore() types.IDepositStore {
	return m.DepositStore
}

func (m *MainStore) GetWithdrawalStore() types.IWithdrawalStore {
	return m.WithdrawalStore
}

func (m *MainStore) GetTaskStore() types.ITaskStore {
	return m.TaskStore
}

func (m *MainStore) GetJobStore() types.IJobStore {
	return m.JobStore
}

func (m *MainStore) GetProcessedBlockStore() types.IProcessedBlockStore {
	return m.ProcessedBlockStore
}

func (m *MainStore) GetEventStore() types.IEventStore {
	return m.EventStore
}

func MustConnectDatabase(cfg *types.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)
	dialect := postgres.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	pgDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	pgDB.SetConnMaxLifetime(time.Duration(cfg.DB.ConnMaxLifetime) * time.Hour)
	pgDB.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	pgDB.SetMaxOpenConns(cfg.DB.MaxOpenConns)

	err = db.Raw("SELECT 1").Error
	if err != nil {
		log.Error("error querying SELECT 1", "err", err)
		panic(err)
	}
	return db, err
}

func MustConnectDatabaseWithName(cfg *types.Config, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, dbName, cfg.DB.Port)
	dialect := postgres.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	pgDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	pgDB.SetConnMaxLifetime(time.Duration(cfg.DB.ConnMaxLifetime) * time.Hour)
	pgDB.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	pgDB.SetMaxOpenConns(cfg.DB.MaxOpenConns)

	err = db.Raw("SELECT 1").Error
	if err != nil {
		log.Error("error querying SELECT 1", "err", err)
		panic(err)
	}
	return db, err
}
