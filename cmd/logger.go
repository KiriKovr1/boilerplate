package main

import (
	"log/slog"
	"os"

	local "github.com/KiriKovr1/logger/pkg/handler/local"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(local.New(os.Stdout, local.LocalHandlerOptions{SlogOpts: slog.HandlerOptions{Level: slog.LevelDebug}}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
