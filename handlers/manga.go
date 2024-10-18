package handlers

import (
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"

	"github.com/gofiber/fiber/v2"
)

func MangaIndex(c *fiber.Ctx) error {
	account, response := getLocalAccount(c)
	if response != nil {
		return response
	}

	if account.AnilistToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "You are not authorised.",
		})
	}

	varUserName := anilist.GraphQLVariable{
		Key:   "userName",
		Value: account.AnilistName,
	}
	listCollection, err := anilist.BuildAndSendRequest[al_types.ALMediaListCollectionData](
		"media_list_collection",
		varUserName,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "There was an issue when trying to retrieve the list collection.",
		})
	}

	return c.JSON(listCollection.MediaListCollection.Lists)
}
