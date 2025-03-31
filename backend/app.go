package backend

import (
	"context"
	"tsuki/backend/config"
	"tsuki/backend/models"
)

type SessionData struct {
	UserID uint `json:"user_id"`
}

// App struct
type App struct {
	CurrentAccount *models.Account
	Config         *config.Config
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
