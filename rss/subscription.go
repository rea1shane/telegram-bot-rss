package rss

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func newSubscription(url, alias string) subscription {
	return subscription{
		url:   url,
		alias: alias,
	}
}

type subscription struct {
	url    string
	alias  string
	filter filter
}

func (s subscription) Fetch() (string, []*gofeed.Item, error) {
	parser := newFeedParser()

	feed, err := parser.ParseURL(s.url)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get feed: %w", err)
	}

	return s.getFeedName(feed), feed.Items, nil
}

func (s subscription) getFeedName(feed *gofeed.Feed) string {
	if s.alias != "" {
		return s.alias
	}

	return feed.Title
}
