package rss

import (
	"fmt"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

type Subscription struct {
	URL      string   `json:"url"`
	Alias    string   `json:"alias,omitempty"`
	Category category `json:"category,omitempty"`
}

func (s Subscription) Fetch(span time.Duration) (string, []*gofeed.Item, error) {
	parser := newFeedParser()

	feed, err := parser.ParseURL(s.URL)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get feed: %w", err)
	}

	// Keep only the most recent items
	items := filter(feed.Items, span)

	// Sort items from oldest to newest
	sortItems(items)

	// Process item categories
	s.Category.handle(items)

	return s.getFeedName(feed), items, nil
}

func (s Subscription) getFeedName(feed *gofeed.Feed) string {
	if s.Alias != "" {
		return s.Alias
	}

	return feed.Title
}

func sortItems(items []*gofeed.Item) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].PublishedParsed == nil {
			return true
		}
		if items[j].PublishedParsed == nil {
			return false
		}
		return items[i].PublishedParsed.Before(*items[j].PublishedParsed)
	})
}

func filter(items []*gofeed.Item, span time.Duration) []*gofeed.Item {
	now := time.Now()
	var result []*gofeed.Item

	for _, item := range items {
		if item.PublishedParsed == nil || now.Sub(*item.PublishedParsed) <= span {
			result = append(result, item)
		}
	}

	return result
}
