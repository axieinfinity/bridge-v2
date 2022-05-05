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
	"syscall"
)

var (
	app        = NewApp("", "", "the bridge-v2 command interface")
	ConfigFlag = cli.StringFlag{
		Name:  "config",
		Usage: "path to config file",
	}
	LogLvlFlag = cli.IntFlag{
		Name:  "logLvl",
		Usage: "log level",
	}
)

func init() {
	app.Action = bridge
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2022 The Sky Mavis Authors"
	app.Flags = append(app.Flags, ConfigFlag, LogLvlFlag)
}

func bridge(ctx *cli.Context) {
	// load log level
	logLvl := log.LvlInfo
	if ctx.GlobalIsSet(LogLvlFlag.Name) {
		logLvl = log.Lvl(ctx.GlobalInt(LogLvlFlag.Name))
	}
	log.Root().SetHandler(log.LvlFilterHandler(logLvl, log.StreamHandler(os.Stdout, log.TerminalFormat(true))))

	// load config file
	if !ctx.GlobalIsSet(ConfigFlag.Name) {
		panic("config path must be defined")
	}
	plan, _ := ioutil.ReadFile(ctx.GlobalString(ConfigFlag.Name))
	var cfg *types.Config
	if err := json.Unmarshal(plan, &cfg); err != nil {
		panic(err)
	}

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
