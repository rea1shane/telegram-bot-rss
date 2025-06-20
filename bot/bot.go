package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	bot    *bot.Bot
	chatId string
}

func New(token, chatID string) (*Bot, error) {
	b, err := bot.New(token)
	if err != nil {
		return nil, fmt.Errorf("failed to new telegram bot: %w", err)
	}

	return &Bot{
		bot:    b,
		chatId: chatID,
	}, nil
}

func (b *Bot) Send(feed, title, link, imageURL string, categories []string) (err error) {
	mode, content := generateContent(feed, title, link, categories)
	ctx := context.Background()

	switch imageURL {
	case "":
		disabledLinkPreview := true
		p := &bot.SendMessageParams{
			ChatID:    b.chatId,
			ParseMode: mode,
			Text:      content,
			LinkPreviewOptions: &models.LinkPreviewOptions{
				IsDisabled: &disabledLinkPreview,
			},
		}
		_, err = b.bot.SendMessage(ctx, p)

	default:
		p := &bot.SendPhotoParams{
			ChatID: b.chatId,
			Photo: &models.InputFileString{
				Data: imageURL,
			},
			ParseMode: mode,
			Caption:   content,
		}
		_, err = b.bot.SendPhoto(ctx, p)
	}

	return
}

func generateContent(feed, title, link string, categories []string) (models.ParseMode, string) {
	mode := models.ParseModeMarkdown

	var builder strings.Builder

	// Feed
	builder.WriteString(fmt.Sprintf("*%s*\n", bot.EscapeMarkdown(feed)))

	// Title and link
	builder.WriteString(fmt.Sprintf("[*%s*](%s)\n", bot.EscapeMarkdown(title), bot.EscapeMarkdown(link)))

	// Blank line
	builder.WriteString("\n")

	// Categories
	for _, category := range categories {
		builder.WriteString(bot.EscapeMarkdown(fmt.Sprintf("#%s ", category)))
	}

	return mode, builder.String()
}
