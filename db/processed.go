package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func init() {
	modules = append(modules, &processed{})
}

type processed struct {
	URL  string `gorm:"primaryKey"`
	GUID string `gorm:"primaryKey"`

	Feed  string
	Title string
	Link  string

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (d *DB) HasBeenProcessed(url, guid string) (bool, error) {
	var m processed

	if err := d.db.Where("url = ? AND guid = ?", url, guid).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to query item: %w", err)
	}

	return true, nil
}

func (d *DB) Record(url, guid, feed, title, link string) error {
	if err := d.db.Create(&processed{
		URL:   url,
		GUID:  guid,
		Feed:  feed,
		Title: title,
		Link:  link,
	}).Error; err != nil {
		return fmt.Errorf("failed to save item: %w", err)
	}
	return nil
}
