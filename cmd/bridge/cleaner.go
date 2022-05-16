package main

import (
	"fmt"
	"github.com/axieinfinity/bridge-v2/cmd/utils"
	"github.com/axieinfinity/bridge-v2/internal/stores"
	"github.com/axieinfinity/bridge-v2/internal/types"
	utils2 "github.com/axieinfinity/bridge-v2/internal/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-co-op/gocron"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cleanerCommand = cli.Command{
		Name:      "cleaner",
		ShortName: "cl",
		Aliases:   nil,
		Usage:     "",
		Action:    utils.MigrateFlags(startCleaner),
		Flags: []cli.Flag{
			ConfigFlag,
			LogLvlFlag,
		},
	}
)

type cleaner struct{}

func startCleaner(ctx *cli.Context) error {
	cfg := prepare(ctx)
	return start(cfg)
}

func start(cfg *types.Config) error {
	db, err := stores.MustConnectDatabase(cfg)
	if err != nil {
		return err
	}
	mainStore := stores.NewMainStore(db)
	u := &utils2.Utils{}
	s := gocron.NewScheduler(time.UTC)
	s.TagsUnique()
	// start read the configuration and create cleaning schedule
	for name, cleanerCfg := range cfg.Cleaner {
		val, err := u.Invoke(&cleaner{}, fmt.Sprintf("Exec%s", name), s, mainStore, cleanerCfg)
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

func (c *cleaner) ExecClearSuccessTasks(scheduler *gocron.Scheduler, store types.IMainStore, cfg *types.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetTaskStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetTaskStore().DeleteTasks([]string{types.STATUS_DONE}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearSuccessTask] error while deleting done tasks", "err", err)
		}
	})
	return err
}

func (c *cleaner) ExecClearFailedTasks(scheduler *gocron.Scheduler, store types.IMainStore, cfg *types.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetTaskStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetTaskStore().DeleteTasks([]string{types.STATUS_FAILED}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearFailedTask] error while deleting failed tasks", "err", err)
		}
	})
	return err
}

func (c *cleaner) ExecClearEvents(scheduler *gocron.Scheduler, store types.IMainStore, cfg *types.CleanerConfig) error {
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

func (c *cleaner) ExecClearSuccessJobs(scheduler *gocron.Scheduler, store types.IMainStore, cfg *types.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetJobStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetJobStore().DeleteJobs([]string{types.STATUS_DONE}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearSuccessTask] error while deleting done tasks", "err", err)
		}
	})
	return err
}

func (c *cleaner) ExecClearFailedJobs(scheduler *gocron.Scheduler, store types.IMainStore, cfg *types.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		if store.GetJobStore().Count() <= int64(cfg.SkipIfLessThan) {
			return
		}
		if err := store.GetJobStore().DeleteJobs([]string{types.STATUS_FAILED}, uint64(time.Now().Unix())-cfg.RemoveAfter); err != nil {
			log.Error("[ExecClearFailedJobs] error while deleting ExecCleanFailedJobs tasks", "err", err)
		}
	})
	return err
}
