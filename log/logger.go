package log

import (
	"fmt"
	"log/slog"
	"os"
)

func NewLogger(level string) (*slog.Logger, error) {
	options := &slog.HandlerOptions{}

	switch level {
	case "debug":
		options.Level = slog.LevelDebug

	case "info":
		options.Level = slog.LevelInfo

	case "warn":
		options.Level = slog.LevelWarn

	case "error":
		options.Level = slog.LevelError

	default:
		return nil, fmt.Errorf("invalid level: %s", level)
	}

	handler := slog.NewTextHandler(os.Stdout, options)

	return slog.New(handler), nil
}
