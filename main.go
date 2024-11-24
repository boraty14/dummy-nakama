package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	// Register the RPC function.

	if err := initializer.RegisterMatch(MatchModule, createMatch); err != nil {
		return fmt.Errorf("failed to register match: %v", err)
	}

	// Register matchmaker matched callback
	if err := initializer.RegisterMatchmakerMatched(matchmakerMatched); err != nil {
		return err
	}
	return nil

	logger.Error("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	err := initializer.RegisterRpc("healthcheck", RpcHealthcheck)
	if err != nil {
		return err
	}

	return nil
}
