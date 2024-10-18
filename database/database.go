package database

import (
	"errors"
	"tsuki/core"
	"tsuki/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB
var ACCOUNT *models.Account

func Connect() {
	databaseLocation := core.CONFIG.Files.Database
	database, err := gorm.Open(sqlite.Open(databaseLocation), &gorm.Config{})
	if err != nil {
		core.CONFIG.Logger.Fatal("Failed to connect to the database: ", err)
	}

	DATABASE = database
	core.CONFIG.Logger.Println("Database connected successfully")
}

func Migrate() {
	DATABASE.AutoMigrate(
		&models.Account{},
		&models.MangaMapping{},
		&models.InstalledProvider{},
	)
}

func RecordExists(query interface{}, record interface{}) bool {
	err := DATABASE.Where(query).First(record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
