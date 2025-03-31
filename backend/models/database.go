package models

import (
	"errors"
	"log"
	"tsuki/backend/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open(config.CONFIG.Files.Database), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	DATABASE = database
	migrate()
}

func migrate() {
	DATABASE.AutoMigrate(
		&Account{},
		&InstalledProvider{},
	)
}

func RecordExists(query interface{}, record interface{}) bool {
	err := DATABASE.Where(query).First(record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
