package internal

import (
	"context"
	bridgeCore "github.com/axieinfinity/bridge-core"
	bridgeCoreStores "github.com/axieinfinity/bridge-core/stores"
	"github.com/axieinfinity/bridge-core/utils"
	"github.com/axieinfinity/bridge-v2/internal/listener"
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

func InitEthereum(ctx context.Context, lsConfig *bridgeCore.LsConfig, store bridgeCoreStores.MainStore) bridgeCore.Listener {
	ethListener, err := listener.NewEthereumListener(ctx, lsConfig, utils.NewUtils(), store)
	if err != nil {
		return nil
	}
	return ethListener
}

func InitRonin(ctx context.Context, lsConfig *bridgeCore.LsConfig, store bridgeCoreStores.MainStore) bridgeCore.Listener {
	roninListener, err := listener.NewRoninListener(ctx, lsConfig, utils.NewUtils(), store)
	if err != nil {
		return nil
	}
	return roninListener
}
