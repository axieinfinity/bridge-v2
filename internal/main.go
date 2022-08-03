package internal

import (
	"context"
	bridgeCore "github.com/axieinfinity/bridge-core"
	bridgeCoreStores "github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/internal/listener"
	roninTask "github.com/axieinfinity/bridge-v2/internal/task"
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

func InitEthereum(ctx context.Context, lsConfig *bridgeCore.LsConfig, store bridgeCoreStores.MainStore, helpers utils.Utils) bridgeCore.Listener {
	ethListener, err := listener.NewEthereumListener(ctx, lsConfig, helpers, store)
	if err != nil {
		log.Error("[EthereumListener]Error while init new ethereum listener", "err", err)
		return nil
	}
	return ethListener
}

func InitRonin(ctx context.Context, lsConfig *bridgeCore.LsConfig, store bridgeCoreStores.MainStore, helpers utils.Utils) bridgeCore.Listener {
	roninListener, err := listener.NewRoninListener(ctx, lsConfig, helpers, store)
	if err != nil {
		log.Error("[RoninListener]Error while init new ronin listener", "err", err)
		return nil
	}
	task, err := roninTask.NewRoninTask(roninListener, store.GetDB(), helpers)
	if err != nil {
		log.Error("[RoninListener][InitRonin] Error while adding new task", "err", err)
		return nil
	}
	roninListener.AddTask(task)
	return roninListener
}
