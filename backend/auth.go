package backend

import (
	"encoding/json"
	"os"
	"tsuki/backend/config"
	"tsuki/backend/models"
)

func (a *App) SignIn(username string, password string) (*models.Account, error) {
	account, err := models.Authenticate(username, password)
	if err != nil {
		return nil, err
	}
	a.CurrentAccount = account
	saveUserID(account.ID)
	return account, nil
}

func (a *App) SignOut() {
	a.CurrentAccount = nil
	saveUserID(0)
}

func saveUserID(userID uint) error {
	data := SessionData{UserID: userID}
	file, err := os.Create(config.CONFIG.Files.Session)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func loadUserID() (uint, error) {
	file, err := os.Open(config.CONFIG.Files.Session)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var data SessionData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return 0, err
	}

	return data.UserID, nil
}
