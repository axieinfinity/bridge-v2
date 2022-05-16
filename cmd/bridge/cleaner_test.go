package main

import (
	"github.com/axieinfinity/bridge-v2/internal/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-co-op/gocron"
	"os"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestCleaner(t *testing.T) {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	cfg := &types.Config{
		Cleaner: types.Cleaner{
			"TestCleaner": &types.CleanerConfig{
				Cron:           "0/1 * * * *",
				RemoveAfter:    100,
				SkipIfLessThan: 10,
				Description:    "this is test cleaner which will be triggered every 1 minute",
			},
		},
		Testing: true,
	}

	wg.Add(1)
	go func() {
		if err := start(cfg); err != nil {
			t.Fatal(err)
		}
	}()
	wg.Wait()
}

func (c *cleaner) ExecTestCleaner(scheduler *gocron.Scheduler, store types.IMainStore, cfg *types.CleanerConfig) error {
	_, err := scheduler.Cron(cfg.Cron).Do(func() {
		println("test cleaner has been reached, cancelling...")
		wg.Done()
	})
	return err
}
