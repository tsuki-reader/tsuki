package handlers

import (
	"tsuki/database"
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
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
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ResponseError{
			Error: "Manga entry not found.",
		})
	}

	// Get the manga mapping for this manga
	var mapping *models.MangaMapping
	result := database.DATABASE.Where(&models.MangaMapping{AnilistID: mediaList.MediaList.Media.Id}).First(&mapping)
	if result.Error != nil {
		mapping = nil
	}
	mediaList.MediaList.Mapping = mapping

	// TODO: Add recommendations and maybe character, staff + relations
	return c.JSON(mediaList.MediaList)
}

func verifyAnilistToken(c *fiber.Ctx, account models.Account) (bool, error) {
	if account.AnilistToken == "" {
		return true, c.Status(fiber.StatusForbidden).JSON(ResponseError{
			Error: "Anilist token not found.",
		})
	}
	return false, nil
}
