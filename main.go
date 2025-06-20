package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/rea1shane/telegram-bot-rss/bot"
	"github.com/rea1shane/telegram-bot-rss/db"
	"github.com/rea1shane/telegram-bot-rss/log"
	"github.com/rea1shane/telegram-bot-rss/rss"
)

func main() {
	// Flags
	logLevel := flag.String("log.level", "info", "Log level in [debug, info, warn, error]")
	sqlitePath := flag.String("sqlite.path", "telegram-bot-rss.db", "SQLite database file path")
	subscriptionsPath := flag.String("subscriptions.path", "subscriptions.json", "JSON format subscription list file path")
	flag.Parse()

	// Logger
	logger, err := log.NewLogger(*logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to new logger: %w", err))
	}

	// DB
	d, err := db.Open(*sqlitePath)
	if err != nil {
		fatal(logger, "failed to open db", err)
	}

	// Bot
	botToken := os.Getenv("TELEGRAM_BOT_KEY")
	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	b, err := bot.New(botToken, chatId)
	if err != nil {
		fatal(logger, "failed to new bot", err)
	}

	// Subscriptions
	subscriptions, err := loadSubscriptions(*subscriptionsPath)
	if err != nil {
		fatal(logger, "failed to load subscriptions", err)
	}

	// Process
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for {
		process(logger, d, b, subscriptions, 7*24*time.Hour)
		<-ticker.C
	}
}

func loadSubscriptions(path string) ([]rss.Subscription, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var subscriptions []rss.Subscription
	if err := json.Unmarshal(data, &subscriptions); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return subscriptions, nil
}

func fatal(logger *slog.Logger, msg string, err error) {
	logger.Error(msg, "error", err.Error())
	os.Exit(1)
}
