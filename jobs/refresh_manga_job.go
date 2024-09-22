package jobs

import (
	"strings"
	"tsuki/database"
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"
	"tsuki/models"

	"gorm.io/gorm/clause"
)

// TODO: Account for when user removes a manga from their list
func RefreshMangaJob() error {
	account, err := database.GetAccount()
	if err != nil {
		return err
	}

	varUserName := anilist.GraphQLVariable{
		Key:   "userName",
		Value: account.Name,
	}
	listCollection, err := anilist.BuildAndSendRequest[al_types.ALMediaListCollectionData](
		"media_list_collection",
		varUserName,
	)
	if err != nil {
		return err
	}

	for _, list := range listCollection.MediaListCollection.Lists {
		if list.IsCustomList {
			continue
		}

		listName := list.Name
		for _, entry := range list.Entries {
			// Update the Manga if it exists by Anilist ID
			clause := clause.OnConflict{
				Columns:   []clause.Column{{Name: "anilist_id"}},
				UpdateAll: true,
			}

			manga := models.Manga{
				AnilistID:   entry.Media.Id,
				Provider:    "filesystem", // TODO: Providers
				ProviderID:  "",
				Name:        entry.Media.Title.English,
				Romaji:      entry.Media.Title.Romaji,
				Progress:    entry.Progress,
				Chapters:    entry.Media.Chapters, // TODO: Take this from the provider if its a higher number
				Status:      entry.Media.Status,
				List:        listName,
				Description: entry.Media.Description,
				Score:       entry.Score,
				StartYear:   entry.Media.StartDate.Year,
				StartMonth:  entry.Media.StartDate.Month,
				Genres:      strings.Join(entry.Media.Genres, ","),
			}

			if err := database.DATABASE.Clauses(clause).Create(&manga).Error; err != nil {
				// If error, skip this manga
				// TODO: Keep a log of all the manga that fail this
				continue
			}
		}
	}

	return nil
}
