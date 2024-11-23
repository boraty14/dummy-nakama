package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	TickRate           = 20
	MinPlayers         = 2
	MaxPlayers         = 2
	MatchmakingTimeout = 10 * time.Second
	MatchModule        = "game"
)

type MatchmakerPresence struct {
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
	Username  string `json:"username"`
	Rating    int    `json:"rating"`
}

type Bot struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Rating     int    `json:"rating"`
	Difficulty string `json:"difficulty"`
}

type MatchState struct {
	TickCount     int64
	WaitStartTime int64
	Players       map[string]runtime.Presence
	Bots          map[string]*Bot
	Started       bool
}

func createMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	state := &MatchState{
		Players: make(map[string]runtime.Presence),
		Bots:    make(map[string]*Bot),
	}

	match := &MatchHandler{
		state: state,
	}

	return match, nil
}

func matchmakerMatched(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, matches []runtime.MatchmakerEntry) (string, error) {
	for _, match := range matches {
		logger.Info("Matched user '%s' named '%s'", match.GetPresence().GetUserId(), match.GetPresence().GetUsername())

		for key, value := range match.GetProperties() {
			logger.Info("Matched on '%s' value '%v'", key, value)
		}
	}

	matchID, err := nk.MatchCreate(ctx, MatchModule, map[string]interface{}{"invited": matches})
	if err != nil {
		logger.Error("Error creating match: %v", err)
		return "", err
	}

	return matchID, nil
}

type MatchHandler struct {
	state *MatchState
}

func (m *MatchHandler) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	// Set the start time for the waiting period

	// Add initial players from matchmaker
	if entries, ok := params["matchmaker_entries"].([]runtime.MatchmakerEntry); ok {
		for _, entry := range entries {
			m.state.Players[entry.GetPresence().GetUserId()] = entry.GetPresence()
		}
	}

	return m.state, TickRate, ""
}

func (m *MatchHandler) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	result := true

	// Custom code to process match join attempt.
	return state, result, ""
}

func (m *MatchHandler) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	return state
}

func (m *MatchHandler) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// Custom code to handle a disconnected/leaving user.
	return state
}

func (m *MatchHandler) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	// Custom code to:
	// - Process the messages received.
	// - Update the match state based on the messages and time elapsed.
	// - Broadcast new data messages to match participants.
	return state
}

func (m *MatchHandler) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	// Custom code to process the termination of match.
	return state
}

func (m *MatchHandler) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, "signal received: " + data
}
