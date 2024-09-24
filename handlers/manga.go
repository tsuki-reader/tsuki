package handlers

import (
	"tsuki/database"
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"

	"github.com/gofiber/fiber/v2"
)

func MangaIndex(c *fiber.Ctx) error {
	account, err := database.GetAccount()
	if err != nil || account.Token == "" || account.Name == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "You are not authorised.",
		})
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
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "There was an issue when trying to retrieve the list collection.",
		})
	}

	return c.JSON(listCollection.MediaListCollection.Lists)
}
