package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	// CGO gerektirmeyen modernc.org/sqlite sürücüsü ile bağlantı
	DB, err = gorm.Open(sqlite.Open("file:data/go-ex-db.db?cache=shared&mode=memory"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Connected to database")
}
