package handlers

import (
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
)

func MangaIndex(c *fiber.Ctx) error {
	account, response := getLocalAccount(c)
	if response != nil {
		return response
	}

	if err := verifyAnilistToken(c, *account); err != nil {
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
	account, response := getLocalAccount(c)
	if response != nil {
		return response
	}

	if err := verifyAnilistToken(c, *account); err != nil {
		return err
	}

	mangaId := c.Params("id")

	// Get the manga from Anilist
	// TODO: Also retrieve the manga mapping if one exists
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

	// TODO: Add recommendations and maybe character, staff + relations
	return c.JSON(mediaList.MediaList)
}

func verifyAnilistToken(c *fiber.Ctx, account models.Account) error {
	if account.AnilistToken == "" {
		return c.Status(fiber.StatusForbidden).JSON(ResponseError{
			Error: "Anilist token not found.",
		})
	}
	return nil
}
