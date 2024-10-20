package al_types

import (
	"tsuki/database"
	"tsuki/models"
)

type ALMediaList struct {
	Progress    int     `json:"progress"`
	CompletedAt ALDate  `json:"completedAt"`
	StartedAt   ALDate  `json:"startedAt"`
	Notes       string  `json:"notes"`
	Score       int     `json:"score"`
	Status      string  `json:"status"`
	Media       ALManga `json:"media"`

	// Doesn't actually come from Anilist and is set programatically
	Mapping *models.MangaMapping `json:"mapping"`
}

type ALMediaListData struct {
	MediaList ALMediaList `json:"MediaList"`
}

func (ml *ALMediaList) SetMangaMapping() {
	var mapping *models.MangaMapping
	result := database.DATABASE.Where(&models.MangaMapping{AnilistID: ml.Media.Id}).First(&mapping)
	if result.Error != nil {
		mapping = nil
	}
	ml.Mapping = mapping
}
