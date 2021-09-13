package main

import (
	"log"
	"math/rand"
	"strings"

	"gorm.io/gorm"
)

type shortURL struct {
	gorm.Model
	ShortCode string `gorm:"uniqueIndex"`
	URL       string
}

func generateUniqueShortCode() string {
	chars := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabsdefghijklmnopqrstuvwxyz0123456789", "")
	var shortcode string
	valid := false

	for !valid {
		index := rand.Intn(len(chars))
		shortcode += chars[index]

		if len(shortcode) > SHORTCODE_LENGTH {
			valid = true
		}
	}

	if !qryIsShortcodeUnique(shortcode) {
		log.Printf("Shortcode %s is not unique, retrying...", shortcode)
		shortcode = generateUniqueShortCode()
	}

	return shortcode
}

func createShortURL(url string) *shortURL {
	return &shortURL{
		ShortCode: generateUniqueShortCode(),
		URL:       url,
	}
}
