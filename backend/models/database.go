package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

func Connect() {
	// databaseLocation := core.CONFIG.Files.Database
	database, err := gorm.Open(sqlite.Open("tsuki.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	DATABASE = database
	migrate()
}

func migrate() {
	DATABASE.AutoMigrate(
		&Account{},
	)
}

// func RecordExists(query interface{}, record interface{}) bool {
// 	err := DATABASE.Where(query).First(record).Error
// 	return !errors.Is(err, gorm.ErrRecordNotFound)
// }
