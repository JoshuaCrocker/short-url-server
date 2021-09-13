package main

import (
	"math/rand"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const SHORTCODE_LENGTH int = 6

var db *gorm.DB

func main() {
	rand.Seed(time.Now().UnixNano())

	var err error
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&shortURL{})
	if err != nil {
		panic(err)
	}

	db.Save(createShortURL("https://google.com"))
}
