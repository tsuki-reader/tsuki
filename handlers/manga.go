package handlers

import (
	"tsuki/database"
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
	"github.com/tsuki-reader/nisshoku/providers"
	"gorm.io/gorm/clause"
)

func MangaIndex(c *fiber.Ctx) error {
	account := getLocalAccount(c)
	if account == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Token invalid",
		})
	}

	performed, err := verifyAnilistToken(c, *account)
	if performed {
		return err
	}

	varUserName := anilist.GraphQLVariable{
		Key:   "userName",
		Value: account.AnilistName,
	}
	listCollection, err := anilist.BuildAndSendRequest[al_types.ALMediaListCollectionData](
		"media_list_collection",
		account.AnilistToken,
		nil,
		varUserName,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "There was an issue when trying to retrieve the list collection.",
		})
	}

	for _, groupList := range listCollection.MediaListCollection.Lists {
		for _, list := range groupList.Entries {
			list.SetMangaMapping(*account)
		}
	}

	return c.JSON(listCollection.MediaListCollection.Lists)
}

func MangaShow(c *fiber.Ctx) error {
	account := getLocalAccount(c)
	if account == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Token invalid",
		})
	}

	performed, err := verifyAnilistToken(c, *account)
	if performed {
		return err
	}

	mangaId := c.Params("id")

	// Get the manga from Anilist
	mediaList, err := getMediaList(*account, mangaId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseError{
			Error: "Manga entry not found.",
		})
	}

	// Get the manga mapping for this manga
	mediaList.MediaList.SetMangaMapping(*account)

	// Get the chapter list and send it in the response
	chapterList := []providers.Chapter{}
	if mediaList.MediaList.Mapping != nil {
		chapterList, _ = mediaList.MediaList.Mapping.InstalledProvider.GetChapterList(mediaList.MediaList.Mapping.ExternalID)
	}

	data := fiber.Map{
		"mediaList": mediaList.MediaList,
		"chapters": chapterList
	}

	// TODO: Add recommendations and maybe character, staff + relations
	return c.JSON(data)
}

func MangaAssignMapping(c *fiber.Ctx) error {
	account := getLocalAccount(c)
	if account == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Token invalid",
		})
	}

	mangaId := c.Params("id")

	mediaList, err := getMediaList(*account, mangaId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseError{
			Error: "Manga entry not found.",
		})
	}

	type params struct {
		InstalledProviderID uint `json:"installed_provider_id"`
	}
	body := new(params)
	if err := c.BodyParser(body); err != nil || body.InstalledProviderID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{
			Error: "An error occurred. Ensure that you are including the installed provider id in the JSON body.",
		})
	}

	installedProvider := models.InstalledProvider{}
	if err := database.DATABASE.First(&installedProvider, body.InstalledProviderID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&ResponseError{
			Error: "Could not find installed provider.",
		})
	}

	searchQuery := mediaList.MediaList.Title()
	providerResult, err := installedProvider.GetFirstSearchResult(searchQuery)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{
			Error: "Could not load provider: " + err.Error(),
		})
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
	mapping := models.MangaMapping{
		AnilistID:           mediaList.MediaList.Media.Id,
		ExternalID:          externalId,
		Progress:            mediaList.MediaList.Progress,
		Chapters:            chapterCount,
		InstalledProviderID: body.InstalledProviderID,
		AccountID:           account.ID,
	}
	if err := database.DATABASE.Clauses(clause).Create(&mapping).Error; err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(ResponseError{
			Error: "Could not assign manga: " + err.Error(),
		})
	}
	// Instead of sending request to db, just set the installed provider here
	mapping.InstalledProvider = installedProvider

	mediaList.MediaList.Mapping = &mapping

	data := fiber.Map{
		"mediaList": mediaList.MediaList,
		"chapters":  chapterList,
	}

	return c.JSON(data)
}

func verifyAnilistToken(c *fiber.Ctx, account models.Account) (bool, error) {
	if account.AnilistToken == "" {
		return true, c.Status(fiber.StatusForbidden).JSON(ResponseError{
			Error: "Anilist token not found.",
		})
	}
	return false, nil
}

func getMediaList(account models.Account, mangaId string) (*al_types.ALMediaListData, error) {
	varUserName := anilist.GraphQLVariable{
		Key:   "userName",
		Value: account.AnilistName,
	}
	varMediaId := anilist.GraphQLVariable{
		Key:   "mediaId",
		Value: mangaId,
	}
	mediaList, err := anilist.BuildAndSendRequest[al_types.ALMediaListData](
		"media_list",
		account.AnilistToken,
		nil,
		varUserName,
		varMediaId,
	)
	return mediaList, err
}
