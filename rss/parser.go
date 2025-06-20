package rss

import (
	"github.com/mmcdole/gofeed"
)

func newFeedParser() *gofeed.Parser {
	parser := gofeed.NewParser()
	parser.UserAgent = "telegram-bot-rss"
	return parser
}
