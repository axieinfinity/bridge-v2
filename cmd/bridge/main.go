package main

import (
	"encoding/json"
	"fmt"
	"github.com/axieinfinity/bridge-v2/internal"
	"github.com/axieinfinity/bridge-v2/internal/stores"
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
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

	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbName     = "DB_NAME"
	dbUsername = "DB_USERNAME"
	dbPassword = "DB_PASSWORD"

	numberOfWorkers = "NUMBER_OF_WORKERS"

	separator = "__"
)

var (
	app        = NewApp("", "", "the bridge-v2 command interface")
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
}

func setRpcUrlFromEnv(cfg *types.Config, rpc string) {
	if rpc == "" {
		return
	}
	str := strings.Split(rpc, separator)
	if len(str) == 2 {
		listener := str[0]
		url := str[1]
		if _, ok := cfg.Listeners[listener]; ok {
			cfg.Listeners[listener].RpcUrl = url
		}
	}
}

func setKeyFromEnv(cfg *types.Config, isValidator bool, key string) {
	if key == "" {
		return
	}
	str := strings.Split(key, separator)
	if len(str) == 2 {
		listener := str[0]
		key := str[1]
		if _, ok := cfg.Listeners[listener]; ok {
			// delete prefix 0x or ronin: and lower key
			key = strings.ToLower(strings.Replace(strings.Replace(key, "0x", "", 1), "ronin:", "", 1))
			if isValidator {
				cfg.Listeners[listener].Secret.Validator = key
			} else {
				cfg.Listeners[listener].Secret.Relayer = key
			}
		}
	}
}

func checkEnv(cfg *types.Config) {
	setRpcUrlFromEnv(cfg, os.Getenv(roninRpc))
	setKeyFromEnv(cfg, true, os.Getenv(roninValidatorKey))
	setKeyFromEnv(cfg, false, os.Getenv(roninRelayKey))
	setRpcUrlFromEnv(cfg, os.Getenv(ethereumRpc))
	setKeyFromEnv(cfg, true, os.Getenv(ethereumValidatorKey))
	setKeyFromEnv(cfg, false, os.Getenv(ethereumRelayerKey))

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

	if os.Getenv(numberOfWorkers) != "" {
		cfg.NumberOfWorkers, _ = strconv.Atoi(os.Getenv(numberOfWorkers))
	}

	// clean key
	os.Setenv(roninValidatorKey, "")
	os.Setenv(roninRelayKey, "")
	os.Setenv(ethereumValidatorKey, "")
	os.Setenv(ethereumRelayerKey, "")
}

func createPgDb(cfg *types.Config) {
	db, err := stores.MustConnectDatabaseWithName(cfg, "postgres")
	if err != nil {
		panic(err)
	}
	if db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DB.DBName)).Error != nil {
		log.Error("error while creating database", "err", err, "dbName", cfg.DB.DBName)
	}
}

func bridge(ctx *cli.Context) {
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

	if os.Getenv(configPath) != "" {
		if err := ctx.GlobalSet(ConfigFlag.Name, os.Getenv(configPath)); err != nil {
			panic(err)
		}
	}

	// load config file
	if !ctx.GlobalIsSet(ConfigFlag.Name) {
		panic("config path must be defined")
	}
	log.Info("loading config from file", "path", ctx.GlobalString(ConfigFlag.Name))
	plan, err := ioutil.ReadFile(ctx.GlobalString(ConfigFlag.Name))
	if err != nil {
		panic(err)
	}
	var cfg *types.Config
	if err := json.Unmarshal(plan, &cfg); err != nil {
		panic(err)
	}

	checkEnv(cfg)

	// try creating db if it does not exist
	createPgDb(cfg)

	// init db
	db, err := stores.MustConnectDatabase(cfg)
	if err != nil {
		panic(err)
	}
	controller, err := internal.New(cfg, db, nil)
	if err != nil {
		panic(err)
	}
	if err = controller.Start(); err != nil {
		panic(err)
	}
	go func() {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
		select {
		case <-sigc:
			controller.Close()
		}
	}()
	controller.Wait()
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
