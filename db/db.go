package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var modules []any

func Open(path string) (*DB, error) {
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	err = conn.AutoMigrate(modules...)
	if err != nil {
		return nil, fmt.Errorf("failed to init tables: %w", err)
	}

	return &DB{db: conn}, nil
}

type DB struct {
	db *gorm.DB
}
