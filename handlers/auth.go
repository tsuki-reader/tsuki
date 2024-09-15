package handlers

import (
	"time"
	"tsuki/database"
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

	// TODO: Get user data from AniList using Viewer query
	newToken := token.AccessToken
	newName := "crxssed"

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

	var account *models.Account
	if account, _ = database.GetAccount(); account != nil {
		if account.Token != "" {
			authenticated = true
		}
	}

	return c.JSON(fiber.Map{
		"authenticated": authenticated,
	})
}
