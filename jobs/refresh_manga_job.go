package jobs

import (
	"errors"
	"tsuki/core"
	"tsuki/database"
	"tsuki/models"

	"gorm.io/gorm"
)

func RefreshMangaJob() error {
	mappings := []models.MangaMapping{}
	result := database.DATABASE.Preload("InstalledProvider").Find(&mappings)
	if result.Error != nil {
		return result.Error
	}

	for _, mangaMapping := range mappings {
		if _, err := RetrieveChaptersForMapping(mangaMapping); err != nil {
			// TODO: Log an error and notify the user
			continue
		}
	}
	return nil
}

func RetrieveChaptersForMapping(mangaMapping models.MangaMapping) ([]models.Chapter, error) {
	chapters, err := mangaMapping.InstalledProvider.GetChapterList(mangaMapping.ExternalID)
	results := []models.Chapter{}
	if err != nil {
		core.CONFIG.Logger.Println(err)
		return results, err
	}

	for _, chapter := range chapters {
		var existingChapter models.Chapter

		// Try to find an existing chapter by ExternalID and AbsoluteNumber
		result := database.DATABASE.Where("external_id = ? AND absolute_number = ?", chapter.ID, chapter.AbsoluteNumber).First(&existingChapter)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// No record found, create a new one
				newChapter := models.Chapter{
					ExternalID:        chapter.ID,
					Title:             chapter.Title,
					Provider:          chapter.Provider,
					AbsoluteNumber:    chapter.AbsoluteNumber,
					InstalledProvider: mangaMapping.InstalledProvider,
				}
				if err := database.DATABASE.Create(&newChapter).Error; err != nil {
					return results, err
				}
				results = append(results, newChapter)
			} else {
				// Some other error occurred
				return results, result.Error
			}
		} else {
			// Record exists, update it
			existingChapter.Title = chapter.Title
			existingChapter.Provider = chapter.Provider
			existingChapter.InstalledProvider = mangaMapping.InstalledProvider

			if err := database.DATABASE.Save(&existingChapter).Error; err != nil {
				return results, err
			}

			results = append(results, existingChapter)
		}
	}

	return results, nil
}
