package middleware

import (
	"fmt"
	"log"
	"strings"
	"tsuki/core"
	"tsuki/database"
	"tsuki/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token invalid",
		})
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(core.CONFIG.Server.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token invalid",
		})
	}

	account, err := getUserFromToken(token)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token invalid",
		})
	}
	c.Locals("account", account)

	return c.Next()
}

func getUserFromToken(token *jwt.Token) (*models.Account, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	id, idOk := claims["id"].(float64)
	username, nameOk := claims["username"].(string)

	if !idOk || !nameOk {
		return nil, fmt.Errorf("claims do not contain the required fields")
	}

	account := models.Account{}
	if err := database.DATABASE.Where(&models.Account{ID: uint(id), Username: username}).First(&account).Error; err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	return &account, nil
}
