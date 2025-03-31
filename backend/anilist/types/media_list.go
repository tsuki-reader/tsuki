package types

import (
	"tsuki/backend/anilist"
	"tsuki/backend/models"
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
	Mapping *models.Mapping `json:"mapping"`
}

type ALMediaListData struct {
	MediaList ALMediaList `json:"MediaList"`
}

func (ml *ALMediaList) SetMapping(account models.Account) {
	var mapping *models.Mapping
	result := models.
		DATABASE.
		Preload("InstalledProvider").
		Preload("Account").
		Where(&models.Mapping{AnilistID: ml.Media.Id, AccountID: account.ID}).
		First(&mapping)
	if result.Error != nil {
		mapping = nil
	}
	ml.Mapping = mapping
}

func (ml *ALMediaList) HighestPossibleChapterCount() int {
	if ml.Progress > ml.Media.Chapters {
		return ml.Progress
	}
	return ml.Media.Chapters
}

func (ml *ALMediaList) Title() string {
	title := ml.Media.Title.English
	if title == "" {
		return ml.Media.Title.Romaji
	}
	return title
}

func GetMediaList(account models.Account, mangaId string) (*ALMediaListData, error) {
	varUserName := anilist.GraphQLVariable{
		Key:   "userName",
		Value: account.AnilistName,
	}
	varMediaId := anilist.GraphQLVariable{
		Key:   "mediaId",
		Value: mangaId,
	}
	mediaList, err := anilist.BuildAndSendRequest[ALMediaListData](
		"media_list",
		account.AnilistToken,
		nil,
		varUserName,
		varMediaId,
	)
	return mediaList, err
}
