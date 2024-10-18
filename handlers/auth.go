package handlers

import (
	"tsuki/database"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
)

type _body struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	account := models.Account{}
	response := parseBody(&account, c, &_body{})
	if response != nil {
		return response
	}

	if err := database.DATABASE.Create(&account).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "Could not create user account. There may already be a user with that username.",
		})
	}

	token, err := account.GenerateJWTToken()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Could not generate token.",
		})
	}

	data := fiber.Map{
		"token": token,
	}
	return c.JSON(data)
}

func Login(c *fiber.Ctx) error {
	body := _body{}
	response := parseBody(&models.Account{}, c, &body)
	if response != nil {
		return response
	}

	actualAccount := models.Account{}
	if err := database.DATABASE.Where(&models.Account{Username: body.Username}).First(&actualAccount).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Incorrect username and/or password.",
		})
	}

	unauthorized := models.ComparePassword(actualAccount.Password, body.Password)
	if unauthorized != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Incorrect username and/or password.",
		})
	}

	token, err := actualAccount.GenerateJWTToken()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Could not generate token.",
		})
	}

	data := fiber.Map{
		"token": token,
	}
	return c.JSON(data)
}

func parseBody(account *models.Account, c *fiber.Ctx, body *_body) error {
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "An error occurred. Ensure that you are including the username and password in the JSON body.",
		})
	}

	if body.Username == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "An error occurred. Ensure that you are including the username and password in the JSON body.",
		})
	}

	username := body.Username
	password, err := models.HashPassword(body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "Could not hash password.",
		})
	}

	account.Username = username
	account.Password = password

	return nil
}
