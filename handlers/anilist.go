package handlers

import (
	"tsuki/core"
	"tsuki/database"
	"tsuki/external/anilist"
	"tsuki/external/anilist/al_types"

	"github.com/gofiber/fiber/v2"
)

type _token struct {
	AccessToken string `json:"access_token"`
}

func AnilistLogin(c *fiber.Ctx) error {
	account := getLocalAccount(c)
	if account == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Token invalid",
		})
	}

	token := new(_token)
	if err := c.BodyParser(token); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "An error occurred. Ensure that you are including the access token in the JSON body.",
		})
	}

	newToken := token.AccessToken

	viewer, _ := anilist.BuildAndSendRequest[al_types.ALViewerData]("viewer", newToken, nil)
	if viewer == nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "Anilist account does not exist.",
		})
	}
	newName := viewer.Viewer.Name

	account.AnilistName = newName
	account.AnilistToken = newToken

	if err := database.DATABASE.Save(&account).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "Could not save Anilist token.",
		})
	}

	data := fiber.Map{
		"success": true,
		"message": "Logged in successfully",
	}
	return c.JSON(data)
}

func AnilistStatus(c *fiber.Ctx) error {
	account := getLocalAccount(c)
	if account == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Token invalid",
		})
	}

	authenticated := false

	var currentViewer al_types.ALViewer
	if account.AnilistToken != "" {
		viewer, _ := anilist.BuildAndSendRequest[al_types.ALViewerData]("viewer", account.AnilistToken, nil)
		if viewer != nil {
			authenticated = true
			currentViewer = viewer.Viewer
		}
	}

	return c.JSON(fiber.Map{
		"authenticated": authenticated,
		"viewer":        currentViewer,
		"clientId":      core.CONFIG.Anilist.ClientID,
	})
}
