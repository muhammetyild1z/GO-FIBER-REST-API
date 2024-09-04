package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect() {
	var err error
	// &gorm.Config{}
	DB, err = gorm.Open(sqlite.Open("go-ex-db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connection to database")
	}
	log.Println("connected database")
}
