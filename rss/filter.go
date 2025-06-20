package rss

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

type filter struct {
	all     []string
	title   []string
	content []string
}

func (f filter) filter(items []*gofeed.Item) (filtered []*gofeed.Item) {
OUTER:
	for _, item := range items {
		for _, keyword := range f.all {
			if strings.Contains(item.Title, keyword) ||
				strings.Contains(item.Content, keyword) {
				continue OUTER
			}
		}

		for _, keyword := range f.title {
			if strings.Contains(item.Title, keyword) {
				continue OUTER
			}
		}

		for _, keyword := range f.content {
			if strings.Contains(item.Content, keyword) {
				continue OUTER
			}
		}

		filtered = append(filtered, item)
	}
	return
}
