package main

import (
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/rea1shane/telegram-bot-rss/bot"
	"github.com/rea1shane/telegram-bot-rss/db"
	"github.com/rea1shane/telegram-bot-rss/rss"
)

func process(logger *slog.Logger, d *db.DB, b *bot.Bot, subscriptions []rss.Subscription, span time.Duration) {
	var wg sync.WaitGroup

	for _, subscription := range subscriptions {
		wg.Add(1)

		go func(logger *slog.Logger, d *db.DB, b *bot.Bot, subscription rss.Subscription) {
			defer wg.Done()

			if err := processSubscription(logger, d, b, subscription, span); err != nil {
				logger.Error("failed to process", "error", err.Error())
			}
		}(logger.With("subscription", subscription.URL), d, b, subscription)
	}

	wg.Wait()
}

func processSubscription(logger *slog.Logger, d *db.DB, b *bot.Bot, subscription rss.Subscription, span time.Duration) error {
	feed, items, err := subscription.Fetch(span)
	if err != nil {
		return fmt.Errorf("failed to fetch: %w", err)
	}

	for _, item := range items {
		if err := processItem(logger, d, b, subscription.URL, feed, item); err != nil && !isProcessedError(err) {
			return fmt.Errorf("failed to process item: %w", err)
		}
	}
	return nil
}

func processItem(logger *slog.Logger, d *db.DB, b *bot.Bot, url, feed string, item *gofeed.Item) error {
	// Check status
	if processed, err := d.HasBeenProcessed(url, item.GUID); err != nil {
		return fmt.Errorf("failed to check item status: %w", err)
	} else if processed {
		return errProcessed
	}

	// Parse fields
	var imageURL string
	if item.Image != nil {
		imageURL = item.Image.URL
	}

	// Send notification
	if err := b.Send(feed, item.Title, item.Link, imageURL, item.Categories); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	// Record item
	if err := d.Record(url, item.GUID, feed, item.Title, item.Link); err != nil {
		return fmt.Errorf("failed to record item: %w", err)
	}

	return nil
}

var errProcessed = errors.New("item processed")

func isProcessedError(err error) bool {
	return errors.Is(err, errProcessed)
}
