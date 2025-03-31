package backend

import (
	"tsuki/backend/anilist/types"
	"tsuki/backend/jobs"
	"tsuki/backend/models"

	"github.com/tsuki-reader/nisshoku/providers"
	"gorm.io/gorm/clause"
)

type MappingAssignResponse struct {
	MediaList types.ALMediaList `json:"media_list"`
	Chapters  []models.Chapter  `json:"chapters"`
}

type MappingChapterPagesResponse struct {
	Pages             []providers.Page         `json:"pages"`
	InstalledProvider models.InstalledProvider `json:"installed_provider"`
}

func (a *App) AssignMapping(id string, installedProviderId uint) (MappingAssignResponse, error) {
	mediaList, err := types.GetMediaList(*a.CurrentAccount, id)
	if err != nil {
		return MappingAssignResponse{}, err
	}

	installedProvider := models.InstalledProvider{}
	if err := models.DATABASE.First(&installedProvider, installedProviderId).Error; err != nil {
		return MappingAssignResponse{}, err
	}

	searchQuery := mediaList.MediaList.Title()
	providerResult, err := installedProvider.GetFirstSearchResult(searchQuery)
	if err != nil {
		return MappingAssignResponse{}, err
	}

	externalId := ""
	if providerResult != nil {
		externalId = providerResult.ID
	}
	chapterCount := mediaList.MediaList.HighestPossibleChapterCount()
	chapterList, _ := installedProvider.GetChapterList(externalId)
	providerChapterCount := len(chapterList)
	if chapterCount < providerChapterCount {
		chapterCount = providerChapterCount
	}

	clause := clause.OnConflict{
		Columns:   []clause.Column{{Name: "anilist_id"}, {Name: "account_id"}},
		UpdateAll: true,
	}
	mapping := models.Mapping{
		AnilistID:           mediaList.MediaList.Media.Id,
		ExternalID:          externalId,
		Progress:            mediaList.MediaList.Progress,
		Chapters:            chapterCount,
		InstalledProviderID: installedProvider.ID,
		AccountID:           a.CurrentAccount.ID,
	}
	if err := models.DATABASE.Clauses(clause).Create(&mapping).Error; err != nil {
		return MappingAssignResponse{}, err
	}
	// Instead of sending request to db, just set the installed provider here
	mapping.InstalledProvider = installedProvider
	mediaList.MediaList.Mapping = &mapping

	chapters, _ := jobs.RetrieveChaptersForMapping(mapping)

	return MappingAssignResponse{
		MediaList: mediaList.MediaList,
		Chapters:  chapters,
	}, nil
}

func (a *App) MangaChapterPages(id int, chapterId string) (MappingChapterPagesResponse, error) {
	var mapping *models.Mapping
	result := models.
		DATABASE.
		Preload("InstalledProvider").
		Preload("Account").
		Where(&models.Mapping{AnilistID: id, AccountID: a.CurrentAccount.ID}).
		First(&mapping)
	if result.Error != nil {
		return MappingChapterPagesResponse{}, result.Error
	}

	installedProvider := mapping.InstalledProvider
	pages, err := installedProvider.GetChapterPages(chapterId)
	if err != nil {
		return MappingChapterPagesResponse{}, err
	}

	return MappingChapterPagesResponse{
		Pages:             pages,
		InstalledProvider: installedProvider,
	}, nil
}
