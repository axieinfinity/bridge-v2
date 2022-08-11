package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/axieinfinity/bridge-v2/internal"

	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/adapters"
	bridgeCoreStore "github.com/axieinfinity/bridge-core/stores"
	bridgeCoreUtils "github.com/axieinfinity/bridge-core/utils"
	migration "github.com/axieinfinity/bridge-migrations"
	"github.com/axieinfinity/bridge-v2/cmd/utils"
	"github.com/axieinfinity/bridge-v2/internal/debug"
	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/urfave/cli.v1"

	kms "github.com/axieinfinity/ronin-kms-client"
)

const (
	configPath           = "CONFIG_PATH"
	roninRpc             = "RONIN_RPC"
	roninValidatorKey    = "RONIN_VALIDATOR_KEY"
	roninRelayKey        = "RONIN_RELAYER_KEY"
	ethereumRpc          = "ETHEREUM_RPC"
	ethereumValidatorKey = "ETHEREUM_VALIDATOR_KEY"
	ethereumRelayerKey   = "ETHEREUM_RELAYER_KEY"
	verbosity            = "VERBOSITY"

	roninValidatorKmsKeyTokenPath = "RONIN_VALIDATOR_KMS_KEY_TOKEN_PATH"
	roninValidatorKmsSslCertPath  = "RONIN_VALIDATOR_KMS_SSL_CERT_PATH"
	roninValidatorKmsServerAddr   = "RONIN_VALIDATOR_KMS_SERVER_ADDR"
	roninValidatorKmsSourceAddr   = "RONIN_VALIDATOR_KMS_SOURCE_ADDR"
	roninValidatorKmsSignTimeout  = "RONIN_VALIDATOR_KMS_SIGN_TIMEOUT"

	roninRelayerKmsKeyTokenPath = "RONIN_RELAYER_KMS_KEY_TOKEN_PATH"
	roninRelayerKmsSslCertPath  = "RONIN_RELAYER_KMS_SSL_CERT_PATH"
	roninRelayerKmsServerAddr   = "RONIN_RELAYER_KMS_SERVER_ADDR"
	roninRelayerKmsSourceAddr   = "RONIN_RELAYER_KMS_SOURCE_ADDR"
	roninRelayerKmsSignTimeout  = "RONIN_RELAYER_KMS_SIGN_TIMEOUT"

	ethereumValidatorKmsKeyTokenPath = "ETHEREUM_VALIDATOR_KMS_KEY_TOKEN_PATH"
	ethereumValidatorKmsSslCertPath  = "ETHEREUM_VALIDATOR_KMS_SSL_CERT_PATH"
	ethereumValidatorKmsServerAddr   = "ETHEREUM_VALIDATOR_KMS_SERVER_ADDR"
	ethereumValidatorKmsSourceAddr   = "ETHEREUM_VALIDATOR_KMS_SOURCE_ADDR"
	ethereumValidatorKmsSignTimeout  = "ETHEREUM_VALIDATOR_KMS_SIGN_TIMEOUT"

	ethereumRelayerKmsKeyTokenPath = "ETHEREUM_RELAYER_KMS_KEY_TOKEN_PATH"
	ethereumRelayerKmsSslCertPath  = "ETHEREUM_RELAYER_KMS_SSL_CERT_PATH"
	ethereumRelayerKmsServerAddr   = "ETHEREUM_RELAYER_KMS_SERVER_ADDR"
	ethereumRelayerKmsSourceAddr   = "ETHEREUM_RELAYER_KMS_SOURCE_ADDR"
	ethereumRelayerKmsSignTimeout  = "ETHEREUM_RELAYER_KMS_SIGN_TIMEOUT"

	roninTaskInterval           = "RONIN_TASK_INTERVAL"
	roninTransactionCheckPeriod = "RONIN_TRANSACTION_CHECK_PERIOD"
	roninMaxProcessingTasks     = "RONIN_MAX_PROCESSING_TASKS"
	ethereumGetLogsBatchSize    = "ETHEREUM_GET_LOGS_BATCH_SIZE"

	roninMaxTasksQuery = "RONIN_MAX_TASKS_QUERY"

	dbHost            = "DB_HOST"
	dbPort            = "DB_PORT"
	dbName            = "DB_NAME"
	dbUsername        = "DB_USERNAME"
	dbPassword        = "DB_PASSWORD"
	dbConnMaxLifeTime = "DB_CONN_MAX_LIFE_TIME"
	dbMaxIdleConns    = "DB_MAX_IDLE_CONNS"
	dbMaxOpenConns    = "DB_MAX_OPEN_CONNS"

	numberOfWorkers = "NUMBER_OF_WORKERS"

	EthereumNetwork = "Ethereum"
	RoninNetwork    = "Ronin"
)

var (
	app        = utils.NewApp("", "", "the bridge-v2 command interface")
	ConfigFlag = cli.StringFlag{
		Name:  "config",
		Usage: "path to config file",
	}
	LogLvlFlag = cli.IntFlag{
		Name:  "verbosity",
		Usage: "log level",
	}
)

func init() {
	app.Action = bridge
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2022 The Sky Mavis Authors"
	app.Flags = append(app.Flags, ConfigFlag, LogLvlFlag)
	app.Flags = append(app.Flags, debug.Flags...)
	app.Before = func(ctx *cli.Context) error {
		return debug.Setup(ctx)
	}
	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		return nil
	}
	app.Commands = []cli.Command{
		cleanerCommand,
	}

	adapters.New()
}

func setRpcUrlFromEnv(cfg *bridgeCore.Config, rpc, network string) {
	if rpc == "" {
		return
	}
	if _, ok := cfg.Listeners[network]; ok {
		cfg.Listeners[network].RpcUrl = rpc
	}
}

func setKeyFromEnv(cfg *bridgeCore.Config, isValidator bool, key, network string) {
	if key == "" {
		return
	}
	if _, ok := cfg.Listeners[network]; ok {
		// delete prefix 0x or ronin: and lower key
		key = strings.ToLower(strings.Replace(strings.Replace(key, "0x", "", 1), "ronin:", "", 1))
		if isValidator {
			cfg.Listeners[network].Secret.Validator = &bridgeCoreUtils.SignMethodConfig{
				PlainPrivateKey: key,
			}
		} else {
			cfg.Listeners[network].Secret.Relayer = &bridgeCoreUtils.SignMethodConfig{
				PlainPrivateKey: key,
			}
		}
	}
}

func setKmsFromEnv(cfg *bridgeCore.Config, isValidator bool, config *kms.KmsConfig, network string) {
	if _, ok := cfg.Listeners[network]; ok {
		if isValidator {
			cfg.Listeners[network].Secret.Validator = &bridgeCoreUtils.SignMethodConfig{
				KmsConfig: config,
			}
		} else {
			cfg.Listeners[network].Secret.Relayer = &bridgeCoreUtils.SignMethodConfig{
				KmsConfig: config,
			}
		}
	}

}

func prepare(ctx *cli.Context) *bridgeCore.Config {
	// load log level
	logLvl := log.LvlInfo
	if os.Getenv(verbosity) != "" {
		if err := ctx.GlobalSet(LogLvlFlag.Name, os.Getenv("VERBOSITY")); err != nil {
			fmt.Println("cannot set verbosity from environment")
		}
	}
	if ctx.GlobalIsSet(LogLvlFlag.Name) {
		logLvl = log.Lvl(ctx.GlobalInt(LogLvlFlag.Name))
	}
	log.Root().SetHandler(log.LvlFilterHandler(logLvl, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	cfg := &bridgeCore.Config{}

	if os.Getenv(configPath) != "" {
		if err := ctx.GlobalSet(ConfigFlag.Name, os.Getenv(configPath)); err != nil {
			panic(err)
		}
	}

	// load config file
	if ctx.GlobalIsSet(ConfigFlag.Name) {
		log.Info("loading config from file", "path", ctx.GlobalString(ConfigFlag.Name))
		plan, err := ioutil.ReadFile(ctx.GlobalString(ConfigFlag.Name))
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(plan, &cfg); err != nil {
			panic(err)
		}
	}

	checkEnv(cfg)

	// try creating db if it does not exist
	createPgDb(cfg)

	return cfg
}

func checkEnv(cfg *bridgeCore.Config) {
	if cfg.DB == nil {
		cfg.DB = &bridgeCoreStore.Database{}
	}

	if os.Getenv(dbHost) != "" {
		log.Info("load db hostname from env", "path", os.Getenv(dbHost))
		cfg.DB.Host = os.Getenv(dbHost)
	}

	if os.Getenv(dbPort) != "" {
		port, err := strconv.Atoi(os.Getenv(dbPort))
		if err != nil {
			panic(err)
		}
		cfg.DB.Port = port
	}

	if os.Getenv(dbName) != "" {
		cfg.DB.DBName = os.Getenv(dbName)
	}

	if os.Getenv(dbUsername) != "" {
		cfg.DB.User = os.Getenv(dbUsername)
	}

	if os.Getenv(dbPassword) != "" {
		cfg.DB.Password = os.Getenv(dbPassword)
	}

	if os.Getenv(dbConnMaxLifeTime) != "" {
		cfg.DB.ConnMaxLifetime, _ = strconv.Atoi(os.Getenv(dbConnMaxLifeTime))
	}

	if os.Getenv(dbMaxIdleConns) != "" {
		cfg.DB.MaxIdleConns, _ = strconv.Atoi(os.Getenv(dbMaxIdleConns))
	}

	if os.Getenv(dbMaxOpenConns) != "" {
		cfg.DB.MaxOpenConns, _ = strconv.Atoi(os.Getenv(dbMaxOpenConns))
	}

	if os.Getenv(numberOfWorkers) != "" {
		cfg.NumberOfWorkers, _ = strconv.Atoi(os.Getenv(numberOfWorkers))
	}

	if cfg.Listeners[RoninNetwork] != nil {

		if os.Getenv(roninTaskInterval) != "" {
			taskInterval, _ := strconv.Atoi(os.Getenv(roninTaskInterval))
			if taskInterval > 0 {
				cfg.Listeners[RoninNetwork].TaskInterval = time.Duration(int64(taskInterval)) * time.Second
				log.Info("setting TaskInterval", "value", cfg.Listeners[RoninNetwork].TaskInterval)
			}
		}

		if os.Getenv(roninTransactionCheckPeriod) != "" {
			txPeriod, _ := strconv.Atoi(os.Getenv(roninTransactionCheckPeriod))
			if txPeriod > 0 {
				cfg.Listeners[RoninNetwork].TransactionCheckPeriod = time.Duration(int64(txPeriod)) * time.Second
				log.Info("setting transactionCheckPeriod", "value", cfg.Listeners[RoninNetwork].TransactionCheckPeriod)
			}
		}

		if os.Getenv(roninMaxProcessingTasks) != "" {
			tasks, _ := strconv.Atoi(os.Getenv(roninMaxProcessingTasks))
			if tasks > 0 {
				cfg.Listeners[RoninNetwork].MaxProcessingTasks = tasks
				log.Info("setting MaxProcessingTasks", "value", tasks)
			}
		}

		if os.Getenv(roninMaxTasksQuery) != "" {
			limit, _ := strconv.Atoi(os.Getenv(roninMaxTasksQuery))
			if limit > 0 {
				cfg.Listeners[RoninNetwork].MaxTasksQuery = limit
			}
		}

		setRpcUrlFromEnv(cfg, os.Getenv(roninRpc), RoninNetwork)

		if os.Getenv(roninValidatorKey) != "" {
			setKeyFromEnv(cfg, true, os.Getenv(roninValidatorKey), RoninNetwork)
		} else if os.Getenv(roninValidatorKmsKeyTokenPath) != "" {
			signTimeout, _ := strconv.Atoi(os.Getenv(roninValidatorKmsSignTimeout))
			config := &kms.KmsConfig{
				KeyTokenPath:  os.Getenv(roninValidatorKmsKeyTokenPath),
				SslCertPath:   os.Getenv(roninValidatorKmsSslCertPath),
				KmsServerAddr: os.Getenv(roninValidatorKmsServerAddr),
				KmsSourceAddr: os.Getenv(roninValidatorKmsSourceAddr),
				SignTimeout:   int64(signTimeout),
			}
			setKmsFromEnv(cfg, true, config, RoninNetwork)
		}

		if os.Getenv(roninRelayKey) != "" {
			setKeyFromEnv(cfg, false, os.Getenv(roninRelayKey), RoninNetwork)
		} else if os.Getenv(roninRelayerKmsKeyTokenPath) != "" {
			signTimeout, _ := strconv.Atoi(os.Getenv(roninRelayerKmsSignTimeout))
			config := &kms.KmsConfig{
				KeyTokenPath:  os.Getenv(roninRelayerKmsKeyTokenPath),
				SslCertPath:   os.Getenv(roninRelayerKmsSslCertPath),
				KmsServerAddr: os.Getenv(roninRelayerKmsServerAddr),
				KmsSourceAddr: os.Getenv(roninRelayerKmsSourceAddr),
				SignTimeout:   int64(signTimeout),
			}
			setKmsFromEnv(cfg, false, config, RoninNetwork)
		}
	}

	if cfg.Listeners[EthereumNetwork] != nil {
		if os.Getenv(ethereumGetLogsBatchSize) != "" {
			batchSize, _ := strconv.Atoi(os.Getenv(ethereumGetLogsBatchSize))
			if batchSize > 0 {
				cfg.Listeners[EthereumNetwork].GetLogsBatchSize = batchSize
			}
		}

		setRpcUrlFromEnv(cfg, os.Getenv(ethereumRpc), EthereumNetwork)

		if os.Getenv(ethereumValidatorKey) != "" {
			setKeyFromEnv(cfg, true, os.Getenv(ethereumValidatorKey), EthereumNetwork)
		} else if os.Getenv(ethereumValidatorKmsKeyTokenPath) != "" {
			signTimeout, _ := strconv.Atoi(os.Getenv(ethereumValidatorKmsSignTimeout))
			config := &kms.KmsConfig{
				KeyTokenPath:  os.Getenv(ethereumValidatorKmsKeyTokenPath),
				SslCertPath:   os.Getenv(ethereumValidatorKmsSslCertPath),
				KmsServerAddr: os.Getenv(ethereumValidatorKmsServerAddr),
				KmsSourceAddr: os.Getenv(ethereumValidatorKmsSourceAddr),
				SignTimeout:   int64(signTimeout),
			}
			setKmsFromEnv(cfg, true, config, EthereumNetwork)
		}

		if os.Getenv(ethereumRelayerKey) != "" {
			setKeyFromEnv(cfg, false, os.Getenv(ethereumRelayerKey), EthereumNetwork)
		} else if os.Getenv(ethereumRelayerKmsKeyTokenPath) != "" {
			signTimeout, _ := strconv.Atoi(os.Getenv(ethereumRelayerKmsSignTimeout))
			config := &kms.KmsConfig{
				KeyTokenPath:  os.Getenv(ethereumRelayerKmsKeyTokenPath),
				SslCertPath:   os.Getenv(ethereumRelayerKmsSslCertPath),
				KmsServerAddr: os.Getenv(ethereumRelayerKmsServerAddr),
				KmsSourceAddr: os.Getenv(ethereumRelayerKmsSourceAddr),
				SignTimeout:   int64(signTimeout),
			}
			setKmsFromEnv(cfg, false, config, EthereumNetwork)
		}
	}

	// clean key
	os.Setenv(roninValidatorKey, "")
	os.Setenv(roninRelayKey, "")
	os.Setenv(ethereumValidatorKey, "")
	os.Setenv(ethereumRelayerKey, "")
}

func createPgDb(cfg *bridgeCore.Config) {
	db, err := bridgeCoreStore.MustConnectDatabaseWithName(cfg.DB, "postgres", false)
	if err != nil {
		panic(err)
	}
	if db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DB.DBName)).Error != nil {
		log.Error("error while creating database", "err", err, "dbName", cfg.DB.DBName)
	}
}

func bridge(ctx *cli.Context) {
	cfg := prepare(ctx)
	// init db
	db, err := bridgeCoreStore.MustConnectDatabase(cfg.DB, false)
	if err != nil {
		panic(err)
	}
	// start migration
	if err = migration.Migrate(db, cfg); err != nil {
		panic(err)
	}
	//init controller
	controller, err := internal.NewBridgeController(cfg, db, nil)
	if err != nil {
		panic(err)
	}
	if err = controller.Start(); err != nil {
		panic(err)
	}
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
	select {
	case <-sigc:
		controller.Close()
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
