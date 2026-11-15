package app

import (
	"log/slog"
	"os"

	"example.com/open-gateway/internal/config"
)

// App holds process-wide dependencies (logger, otel, redis, kms placeholders).
type App struct {
	Config *config.Config
	Logger *slog.Logger
}

// New wires the application container from configuration.
func New(cfg *config.Config) *App {
	level := slog.LevelInfo
	if cfg.Log.Level == "debug" {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))

	return &App{
		Config: cfg,
		Logger: logger,
	}
}