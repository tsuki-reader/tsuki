package backend

import (
	"context"
	"tsuki/backend/models"
)

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
}

func (a *App) SignIn(username string, password string) (*models.Account, error) {
	account, err := models.Authenticate(username, password)
	if err != nil {
		return nil, err
	}
	a.CurrentAccount = account
	return account, nil
}
