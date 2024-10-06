package database

import (
	"errors"
	"tsuki/core"
	"tsuki/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func UpdateAccount(account *models.Account) (*models.Account, error) {
	// Create/update the account regardless of whether it exists or not.
	clause := clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}

	if err := DATABASE.Clauses(clause).Create(account).Error; err != nil {
		return nil, err
	}

	ACCOUNT = account

	return account, nil
}

func GetAccount() (*models.Account, error) {
	if ACCOUNT != nil {
		return ACCOUNT, nil
	}

	var account models.Account
	if err := DATABASE.First(&account).Error; err != nil {
		return nil, err
	}

	ACCOUNT = &account

	return &account, nil
}

func RecordExists(query interface{}, record interface{}) bool {
	err := DATABASE.Where(query).First(record).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
