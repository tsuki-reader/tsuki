package handlers

import (
	"time"
	"tsuki/core"
	"tsuki/database"
	"tsuki/external/anilist"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Token struct {
	AccessToken string `json:"access_token"`
}

func Login(c *fiber.Ctx) error {
	token := new(Token)

	if err := c.BodyParser(token); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "An error occurred. Ensure that you are including the access token in the JSON body.",
		})
	}

	newToken := token.AccessToken
	anilist.SetupClient(newToken)

	viewer, _ := anilist.BuildAndSendRequest[anilist.ViewerData]("viewer")
	if viewer == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Anilist account does not exist.",
		})
	}
	newName := viewer.Viewer.Name

	account := models.Account{
		Model: gorm.Model{
			ID:        1,
			UpdatedAt: time.Now(),
		},
		Token: newToken,
		Name:  newName,
	}

	if _, err := database.UpdateAccount(&account); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An error occurred while trying to save the token.",
		})
	}

	data := fiber.Map{
		"success": true,
		"message": "Logged in successfully",
	}

	return c.JSON(data)
}

func Logout(c *fiber.Ctx) error {
	account := models.Account{
		Model: gorm.Model{
			ID:        1,
			UpdatedAt: time.Now(),
		},
		Token: "",
	}

	if _, err := database.UpdateAccount(&account); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "An error occurred while trying to logout.",
		})
	}

	data := fiber.Map{
		"success": true,
		"message": "Logged out successfully",
	}

	return c.JSON(data)
}

func Status(c *fiber.Ctx) error {
	authenticated := false

	var currentViewer anilist.Viewer
	if account, _ := database.GetAccount(); account != nil {
		if account.Token != "" {
			anilist.SetupClient(account.Token)

			viewer, _ := anilist.BuildAndSendRequest[anilist.ViewerData]("viewer")
			if viewer != nil {
				authenticated = true
				currentViewer = viewer.Viewer
			}
		}
	}

	return c.JSON(fiber.Map{
		"authenticated": authenticated,
		"viewer":        currentViewer,
		"clientId":      core.CONFIG.Anilist.ClientID,
	})
}
