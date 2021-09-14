package main

import (
	"math/rand"
	"time"
)

const SHORTCODE_LENGTH int = 6

func main() {
	rand.Seed(time.Now().UnixNano())

	db := dbconn.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}
