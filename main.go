package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-telegram/bot"
)

func main() {
	logLevel := flag.String("log.level", "info", "Log level in [debug, info, warn, error]")
	flag.Parse()

	logger, err := newLogger(*logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to new logger: %w", err))
	}

	telegramBot, err := newTelegramBot(os.Getenv("TELEGRAM_BOT_KEY"))
	if err != nil {
		logger.Error("failed to new telegram bot", "err", err)
		os.Exit(1)
	}
}

func newLogger(level string) (*slog.Logger, error) {
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
		return nil, fmt.Errorf("invalid level: %s", level))
	}

	handler := slog.NewTextHandler(os.Stdout, options)

	return slog.New(handler), nil
}

func newTelegramBot(token string) (*bot.Bot, error) {
	return bot.New(token)
}
