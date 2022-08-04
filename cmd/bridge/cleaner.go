package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	bridgeCore "github.com/axieinfinity/bridge-core"
	bridgeCoreStore "github.com/axieinfinity/bridge-core/stores"
	bridgeCoreUtils "github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/cmd/utils"
	"github.com/axieinfinity/bridge-v2/stores"
	"github.com/axieinfinity/bridge-v2/task"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-co-op/gocron"
	"gopkg.in/urfave/cli.v1"
)

var (
	cleanerCommand = cli.Command{
		Name:      "cleaner",
		ShortName: "cl",
		Aliases:   nil,
		Usage:     "activate cleaner",
		Action:    utils.MigrateFlags(startCleaner),
		Flags: []cli.Flag{
			ConfigFlag,
			LogLvlFlag,
		},
	}
)

type cleanerStore struct {
	stores.BridgeStore
	bridgeCoreStore.MainStore
}

type cleaner struct{}

func startCleaner(ctx *cli.Context) error {
	cfg := prepare(ctx)
	return start(cfg)
}

func start(cfg *bridgeCore.Config) error {
	db, err := bridgeCoreStore.MustConnectDatabase(cfg.DB, false)
	if err != nil {
		return err
	}
	listenHandlerStore := stores.NewBridgeStore(db)
	bridgeStore := bridgeCoreStore.NewMainStore(db)
	store := &cleanerStore{listenHandlerStore, bridgeStore}
	u := bridgeCoreUtils.NewUtils()
	s := gocron.NewScheduler(time.UTC)
	s.TagsUnique()
	// start read the configuration and create cleaning schedule
	for name, cleanerCfg := range cfg.Cleaner {
		val, err := u.Invoke(&cleaner{}, fmt.Sprintf("Exec%s", name), s, store, cleanerCfg)
		if err != nil || !val.IsNil() {
			log.Error("[startCleaner] error while executing a cleaner", "err", err, "returnedErr", val.Interface(), "name", name)
		}
	}
	if s.Len() == 0 {
		panic("no job is running")
	}
	log.Info("start scheduler")
	s.StartAsync()
	sigc := make(chan os.Signal, 1)
	go func() {
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigc)
	}()
	<-sigc
	s.Clear()
	return nil
}

func (c *cleaner) ExecClearSuccessTasks(scheduler *gocron.Scheduler, store *cleanerStore, cfg *bridgeCore.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetTaskStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetTaskStore().DeleteTasks([]string{task.STATUS_DONE}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearSuccessTask] error while deleting done tasks", "err", err)
		}
	})
	return err
}

func (c *cleaner) ExecClearFailedTasks(scheduler *gocron.Scheduler, store *cleanerStore, cfg *bridgeCore.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetTaskStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetTaskStore().DeleteTasks([]string{task.STATUS_FAILED}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearFailedTask] error while deleting failed tasks", "err", err)
		}
	})
	return err
}

func (c *cleaner) ExecClearEvents(scheduler *gocron.Scheduler, store *cleanerStore, cfg *bridgeCore.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetEventStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetEventStore().DeleteEvents(uint64(time.Now().Unix()) - cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearFailedTask] error while deleting failed tasks", "err", err)
		}
	})
	return err
}

func (c *cleaner) ExecClearSuccessJobs(scheduler *gocron.Scheduler, store *cleanerStore, cfg *bridgeCore.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetJobStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetJobStore().DeleteJobs([]string{task.STATUS_DONE}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearSuccessTask] error while deleting done tasks", "err", err)
		}
	})
	return err
}

func (c *cleaner) ExecClearFailedJobs(scheduler *gocron.Scheduler, store *cleanerStore, cfg *bridgeCore.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetJobStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetJobStore().DeleteJobs([]string{task.STATUS_FAILED}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearFailedJobs] error while deleting ExecCleanFailedJobs tasks", "err", err)
		}
	})
	return err
}
