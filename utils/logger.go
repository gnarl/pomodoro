package utils

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func InitJsonLogger(level slog.Level) {
	opts := &slog.HandlerOptions{Level: level}
	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger = slog.New(handler)
	slog.SetDefault(logger)
}

func GetLogger() *slog.Logger {
	if logger == nil {
		panic("logger not initialized")
	}

	return logger
}
