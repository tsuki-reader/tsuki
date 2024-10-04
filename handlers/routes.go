package handlers

import "github.com/gofiber/fiber/v2"

type ResponseError struct {
	Error string `json:"error"`
}

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// ========== Auth
	api.Get("/auth/status", Status)
	api.Post("/auth/login", Login)
	api.Post("/auth/logout", Logout)

	// ========== Manga
	api.Get("/manga", MangaIndex)

	// ========== Repositories
	api.Get("/repositories", RepositoriesIndex)
	api.Post("/repositories", RepositoriesCreate)
	api.Delete("/repositories/:id", RepositoriesDestroy)
	api.Patch("/repositories/:id", RepositoriesUpdate)

	// ========== Providers
	api.Post("/providers", ProvidersCreate)
}
