package rss

import (
	"fmt"
	"testing"
	"time"
)

var (
	url   = "https://feeds.feedburner.com/ruanyifeng"
	alias = ""
)

func TestSubscription_Fetch(t *testing.T) {
	name, items, err := Subscription{
		URL:   url,
		Alias: alias,
	}.Fetch(7 * 24 * time.Hour)
	if err != nil {
		t.Fatalf("failed to fetch: %v", err)
	}

	fmt.Printf("Feed name: %s\n", name)

	for _, item := range items {
		fmt.Println()
		fmt.Printf("Item GUID:             %s\n", item.GUID)
		fmt.Printf("Item name:             %s\n", item.Title)
		fmt.Printf("Item link:             %s\n", item.Link)
		fmt.Printf("Item image:            %s\n", item.Image)
		fmt.Printf("Item categories:       %s\n", item.Categories)
		fmt.Printf("Item published:        %s\n", item.Published)
		fmt.Printf("Item published parsed: %s\n", item.PublishedParsed)
		fmt.Printf("Item updated:          %s\n", item.Updated)
		fmt.Printf("Item updated parsed:   %s\n", item.UpdatedParsed)
	}
}
