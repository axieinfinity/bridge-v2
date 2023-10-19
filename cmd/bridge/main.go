package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/axieinfinity/bridge-v2/internal"
	"github.com/axieinfinity/bridge-v2/stats"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/adapters"
	bridgeCoreStore "github.com/axieinfinity/bridge-core/stores"
	bridgeCoreUtils "github.com/axieinfinity/bridge-core/utils"
	migration "github.com/axieinfinity/bridge-migrations"
	"github.com/axieinfinity/bridge-v2/cmd/utils"
	"github.com/axieinfinity/bridge-v2/internal/debug"
	"github.com/ethereum/go-ethereum/log"
	"gopkg.in/urfave/cli.v1"
)

const (
	configPath   = "CONFIG_PATH"
	verbosity    = "VERBOSITY"
	RoninNetwork = "ronin"
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
	adapters.New()
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

	var path string
	if os.Getenv(configPath) != "" {
		if err := ctx.GlobalSet(ConfigFlag.Name, os.Getenv(configPath)); err != nil {
			panic(err)
		}
	}
	// load config file
	if ctx.GlobalIsSet(ConfigFlag.Name) {
		log.Info("loading config from file", "path", ctx.GlobalString(ConfigFlag.Name))
		path = ctx.GlobalString(ConfigFlag.Name)
	}
	if path == "" {
		panic("config path is empty")
	}
	// load config from path and environment
	cfg := loadConfig(path)
	// try creating db if it does not exist
	createPgDb(cfg)
	return cfg
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

func setupStats(cfg *bridgeCore.Config, networkName string, db *gorm.DB) {
	if network, ok := cfg.Listeners[networkName]; ok && network.Stats != nil {
		var operator string
		if network.Secret != nil && network.Secret.BridgeOperator != nil {
			signMethod, err := bridgeCoreUtils.NewSignMethod(network.Secret.BridgeOperator)
			if err != nil {
				panic(err)
			}
			operator = signMethod.GetAddress().Hex()
		}
		stats.NewService(network.Stats.Node, network.ChainId, operator, network.Stats.Host, network.Stats.Secret, db)
		go stats.BridgeStats.Start()
	}
}

func bridge(ctx *cli.Context) {
	cfg := prepare(ctx)
	// init db
	db, err := bridgeCoreStore.MustConnectDatabase(cfg.DB, false)
	if err != nil {
		panic(err)
	}

	// setup stats
	setupStats(cfg, RoninNetwork, db)

	// start migration
	if err = migration.Migrate(db, cfg); err != nil {
		panic(err)
	}

	//init controller
	log.Info("starting bridge controller")
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
		if stats.BridgeStats != nil {
			stats.BridgeStats.Stop()
		}
	}
}

func loadConfig(path string) *bridgeCore.Config {
	// get type from path's extension, default is json
	configType := "json"
	exts := strings.Split(path, ".")
	if len(exts) > 0 {
		ext := strings.ToLower(exts[len(exts)-1])
		if ext == "yaml" || ext == "yml" {
			configType = ext
		}
	}
	cfg := &bridgeCore.Config{}
	viper.SetConfigType(configType)
	if path != "" {
		plan, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		err = viper.ReadConfig(bytes.NewBuffer(plan))
		if err != nil {
			panic(err)
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	err := viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
