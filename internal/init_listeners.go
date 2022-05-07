package internal

import (
	"context"
	"github.com/axieinfinity/bridge-v2/internal/listener"
	"github.com/axieinfinity/bridge-v2/internal/types"
)

func (c *Controller) InitEthereum(ctx context.Context, lsConfig *types.LsConfig, store types.IMainStore) *listener.EthereumListener {
	ethListener, err := listener.NewEthereumListener(ctx, lsConfig, c.utilWrapper, store)
	if err != nil {
		return nil
	}
	return ethListener
}

func (c *Controller) InitRonin(ctx context.Context, lsConfig *types.LsConfig, store types.IMainStore) *listener.RoninListener {
	roninListener, err := listener.NewRoninListener(ctx, lsConfig, c.utilWrapper, store)
	if err != nil {
		return nil
	}
	return roninListener
}
