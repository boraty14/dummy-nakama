package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

type HealthCheckResponse struct {
	Success bool `json:"success"`
}

func RpcHealthcheck(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("Healtcheck RPC called")
	response := &HealthCheckResponse{Success: true}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error("Error marshalling response: %v", err)
		return "", err
	}

	return string(responseBytes), nil
}
