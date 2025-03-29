package backend

import (
	"context"
	"encoding/json"
	"os"
	"tsuki/backend/models"
)

type SessionData struct {
	UserID uint `json:"user_id"`
}

// App struct
type App struct {
	CurrentAccount *models.Account
	ctx            context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	userID, _ := loadUserID()

	if userID == 0 {
		return
	}

	account := models.Account{}
	err := (&models.Account{}).Find(userID, &account)
	if err != nil {
		return
	}
	a.CurrentAccount = &account
}

// domReady is called after front-end resources have been loaded
func (a App) DomReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
	if a.CurrentAccount != nil {
		saveUserID(a.CurrentAccount.ID)
	}
}

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
	file, err := os.Create("session.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}

func loadUserID() (uint, error) {
	file, err := os.Open("session.json")
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
