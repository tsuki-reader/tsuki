package handlers

import (
	"tsuki/middleware"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
)

type ResponseError struct {
	Error string `json:"error"`
}

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Use(middleware.AuthMiddleware)

	auth := app.Group("/auth")

	// ========== Auth
	auth.Post("/register", Register)
	auth.Post("/login", Login)

	// ========== Anilist
	api.Get("/anilist/status", AnilistStatus)
	api.Post("/anilist/login", AnilistLogin)

	// ========== Manga
	api.Get("/manga", MangaIndex)

	// ========== Repositories
	api.Get("/repositories", RepositoriesIndex)
	api.Post("/repositories", RepositoriesCreate)
	api.Patch("/repositories/:id", RepositoriesUpdate)
	api.Delete("/repositories/:id", RepositoriesDestroy)

	// ========== Providers
	api.Get("/providers", ProvidersIndex)
	api.Post("/providers", ProvidersCreate)
	api.Patch("/providers/:id", ProvidersUpdate)
	api.Delete("/providers/:id", ProvidersDestroy)
}

func getLocalAccount(c *fiber.Ctx) (*models.Account, error) {
	local := c.Locals("account")
	account, ok := local.(*models.Account)
	if !ok {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(ResponseError{
			Error: "Token invalid",
		})
	}

	return account, nil
}
