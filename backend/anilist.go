package backend

import (
	"errors"
	"tsuki/backend/anilist"
	"tsuki/backend/anilist/types"
	"tsuki/backend/models"
)

type AnilistStatus struct {
	Authenticated bool            `json:"authenticated"`
	Viewer        *types.ALViewer `json:"viewer"`
	ClientID      string          `json:"client_id"`
}

type LoginStatus struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (a *App) AnilistLogin(token string) (*LoginStatus, error) {
	viewer, _ := anilist.BuildAndSendRequest[types.ALViewerData]("viewer", token, nil)
	if viewer == nil {
		return nil, errors.New("Anilist account does not exist.")
	}
	a.CurrentAccount.AnilistName = viewer.Viewer.Name
	a.CurrentAccount.AnilistToken = token
	if err := models.DATABASE.Save(&a.CurrentAccount).Error; err != nil {
		return nil, err
	}
	return &LoginStatus{
		Success: true,
		Message: "Anilist token saved successfully.",
	}, nil
}

func (a *App) AnilistStatus() AnilistStatus {
	authenticated := false
	var currentViewer *types.ALViewer
	if a.CurrentAccount.AnilistToken != "" {
		viewer, _ := anilist.BuildAndSendRequest[types.ALViewerData]("viewer", a.CurrentAccount.AnilistToken, nil)
		if viewer != nil {
			authenticated = true
			currentViewer = &viewer.Viewer
		}
	}
	return AnilistStatus{
		Authenticated: authenticated,
		Viewer:        currentViewer,
		ClientID:      "21156",
	}
}
