package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	// Register the RPC function.
	initStart := time.Now()

	err := initializer.RegisterRpc("healthcheck", RpcHealthcheck)
	if err != nil {
		return err
	}

	logger.Info("Module loaded in: %v", time.Since(initStart))
	return nil
}
