package rss

import (
	"github.com/mmcdole/gofeed"
)

type categoryMode string

const (
	categoryModeAppend   categoryMode = "APPEND"
	categoryModeOverride categoryMode = "OVERRIDE"
	categoryModeEmpty    categoryMode = "EMPTY"
)

type category struct {
	Mode    categoryMode `json:"mode"`
	Content []string     `json:"content"`
}

func (c category) handle(items []*gofeed.Item) {
	for _, item := range items {
		switch c.Mode {
		case categoryModeAppend:
			item.Categories = append(item.Categories, c.Content...)
		case categoryModeOverride:
			item.Categories = c.Content
		case categoryModeEmpty:
			item.Categories = nil
		}

		item.Categories = formatCategories(item.Categories)
	}
}

func formatCategories(input []string) []string {
	seen := make(map[string]bool)
	var unique []string

	for _, val := range input {
		if !seen[val] {
			seen[val] = true
			unique = append(unique, val)
		}
	}

	return unique
}
