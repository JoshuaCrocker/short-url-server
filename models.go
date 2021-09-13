package main

import (
	"gorm.io/gorm"
)

type shortURL struct {
	gorm.Model
	ShortCode string
	URL       string
}
