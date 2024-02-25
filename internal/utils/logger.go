package utils

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func init() {

	InitJsonLogger(slog.LevelDebug)
}

func InitJsonLogger(level slog.Level) {
	opts := &slog.HandlerOptions{Level: level}
	handler := slog.NewJSONHandler(os.Stdout, opts)

	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}
