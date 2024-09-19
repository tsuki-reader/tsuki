package database

import (
	"tsuki/interfaces"
	"tsuki/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DATABASE *gorm.DB
var ACCOUNT *models.Account
var LOGGER interfaces.LoggerInterface = interfaces.NewLogger()

// TODO: In tests, this creates a new tsuki.db file
func Connect() {
	database, err := gorm.Open(sqlite.Open("tsuki.db"), &gorm.Config{})
	if err != nil {
		LOGGER.Fatal("Failed to connect to the database: ", err)
	}

	DATABASE = database
	LOGGER.Println("Database connected successfully")
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
