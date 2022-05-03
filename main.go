package main

import (
	"github.com/axieinfinity/bridge-v2/internal/stores"
	"github.com/axieinfinity/bridge-v2/internal/types"
)

func main() {
	cfg := &types.Config{
		DB: types.Database{
			Host:     "localhost",
			User:     "postgres",
			Password: "example",
			DBName:   "bridge",
			Port:     5432,
		}}
	db, err := stores.MustConnectDatabase(cfg)
	if err != nil {
		panic(err)
	}
	stores.NewMainStore(db)
}
