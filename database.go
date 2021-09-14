package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbconn databaseConnection = &database{}

type database struct {
	db *gorm.DB
}

type databaseConnection interface {
	// initialiseDatabase connection.
	initialiseDatabase()

	// GetDB retrieves the active database connection.
	//
	// If no database connection is active, it is initialised.
	GetDB() *gorm.DB
}

func (d *database) initialiseDatabase() {
	var err error
	d.db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = d.db.AutoMigrate(&shortURL{})
	if err != nil {
		panic(err)
	}
}

func (d *database) GetDB() *gorm.DB {
	if d.db == nil {
		log.Println("no active database connection - initialising database")
		d.initialiseDatabase()
	}

	return d.db
}
