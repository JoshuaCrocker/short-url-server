package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const SHORTCODE_LENGTH int = 6

var db *gorm.DB

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&shortURL{})
}
