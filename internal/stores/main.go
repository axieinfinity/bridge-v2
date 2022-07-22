package stores

import (
	"fmt"
	"time"

	"github.com/axieinfinity/bridge-v2/configs"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormprometheus "gorm.io/plugin/prometheus"
)

var ()

type MainStore struct {
	*gorm.DB

	DepositStore          types.IDepositStore
	WithdrawalStore       types.IWithdrawalStore
	JobStore              types.IJobStore
	TaskStore             types.ITaskStore
	ProcessedBlockStore   types.IProcessedBlockStore
	EventStore            types.IEventStore
	ProcessedReceiptStore types.IProcessedReceiptStore
}

func NewMainStore(db *gorm.DB) *MainStore {
	cl := &MainStore{
		DB: db,

		JobStore:              NewJobStore(db),
		TaskStore:             NewTaskStore(db),
		ProcessedBlockStore:   NewProcessedBlockStore(db),
		DepositStore:          NewDepositStore(db),
		WithdrawalStore:       NewWithdrawalStore(db),
		EventStore:            NewEventStore(db),
		ProcessedReceiptStore: NewProcessedReceiptStore(db),
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

func (m *MainStore) GetProcessedReceiptStore() types.IProcessedReceiptStore {
	return m.ProcessedReceiptStore
}

func (m *MainStore) GetEventStore() types.IEventStore {
	return m.EventStore
}

func MustConnectDatabase(cfg *types.Config) (*gorm.DB, error) {
	// load sqlite db for testing purpose
	if cfg.Testing {
		return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

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
	if err := db.Use(gormprometheus.New(gormprometheus.Config{
		DBName:          cfg.DB.DBName,                        // use `DBName` as metrics label
		RefreshInterval: 15,                                   // Refresh metrics interval (default 15 seconds)
		PushAddr:        configs.AppConfig.Prometheus.PushURL, // push metrics if `PushAddr` configured
		StartServer:     true,                                 // start http server to expose metrics
		HTTPServerPort:  8082,                                 // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
		MetricsCollector: []gormprometheus.MetricsCollector{
			&gormprometheus.Postgres{
				VariableNames: []string{"Threads_running"},
			},
		}, // user defined metrics
	})); err != nil {
		panic(err)
	}
	err = db.Raw("SELECT 1").Error
	if err != nil {
		log.Error("error querying SELECT 1", "err", err)
		panic(err)
	}
	return db, err
}

func MustConnectDatabaseWithName(cfg *types.Config, dbName string) (*gorm.DB, error) {
	var (
		err error
		db  *gorm.DB
	)
	// load sqlite db for testing purpose
	if cfg.Testing {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, dbName, cfg.DB.Port)
		dialect := postgres.Open(dsn)
		db, err = gorm.Open(dialect, &gorm.Config{})
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
	}

	if err := db.Use(gormprometheus.New(gormprometheus.Config{
		DBName:          cfg.DB.DBName,    // use `DBName` as metrics label
		RefreshInterval: 15,               // Refresh metrics interval (default 15 seconds)
		PushAddr:        "localhost:9091", // push metrics if `PushAddr` configured
		StartServer:     true,             // start http server to expose metrics
		HTTPServerPort:  8082,             // configure http server port, default port 8080 (if you have configured multiple instances, only the first `HTTPServerPort` will be used to start server)
		MetricsCollector: []gormprometheus.MetricsCollector{
			&gormprometheus.Postgres{
				VariableNames: []string{"Threads_running"},
			},
		}, // user defined metrics
	})); err != nil {
		panic(err)
	}

	err = db.Raw("SELECT 1").Error
	if err != nil {
		log.Error("error querying SELECT 1", "err", err)
		panic(err)
	}
	return db, err
}
