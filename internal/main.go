package internal

import (
	"context"
	"fmt"

	bridgeCore "github.com/axieinfinity/bridge-core"
	"github.com/axieinfinity/bridge-core/metrics"
	bridgeCoreStores "github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/listener"
	roninTask "github.com/axieinfinity/bridge-v2/task"
	"github.com/ethereum/go-ethereum/log"
	"gorm.io/gorm"
)

type BridgeController struct {
	*bridgeCore.Controller
}

func NewBridgeController(cfg *bridgeCore.Config, db *gorm.DB, helpers utils.Utils) (*BridgeController, error) {
	bridgeCore.AddListener("Ethereum", InitEthereum)
	bridgeCore.AddListener("Ronin", InitRonin)
	controller, err := bridgeCore.New(cfg, db, helpers)
	if err != nil {
		return nil, err
	}
	return &BridgeController{controller}, nil
}

func InitEthereum(ctx context.Context, lsConfig *bridgeCore.LsConfig, store bridgeCoreStores.MainStore, helpers utils.Utils, pool *bridgeCore.Pool) bridgeCore.Listener {
	ethListener, err := listener.NewEthereumListener(ctx, lsConfig, helpers, store, pool)
	if err != nil {
		log.Error("[EthereumListener]Error while init new ethereum listener", "err", err)
		return nil
	}
	metrics.Pusher.AddCounter(fmt.Sprintf(metrics.ListenerProcessedBlockMetric, ethListener.GetName()), "count number of processed block in ethereum listener")
	return ethListener
}

func InitRonin(ctx context.Context, lsConfig *bridgeCore.LsConfig, store bridgeCoreStores.MainStore, helpers utils.Utils, pool *bridgeCore.Pool) bridgeCore.Listener {
	roninListener, err := listener.NewRoninListener(ctx, lsConfig, helpers, store, pool)
	if err != nil {
		log.Error("[RoninListener]Error while init new ronin listener", "err", err)
		return nil
	}
	metrics.Pusher.AddCounter(fmt.Sprintf(metrics.ListenerProcessedBlockMetric, roninListener.GetName()), "count number of processed block in ethereum listener")

	task, err := roninTask.NewRoninTask(roninListener, store.GetDB(), helpers)
	if err != nil {
		log.Error("[RoninListener][InitRonin] Error while adding new task", "err", err)
		return nil
	}
	roninListener.AddTask(task)
	return roninListener
}
