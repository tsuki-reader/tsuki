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
	api.Patch("/repositories/:id", RepositoriesUpdate)
	api.Delete("/repositories/:id", RepositoriesDestroy)

	// ========== Providers
	api.Get("/providers", ProvidersIndex)
	api.Post("/providers", ProvidersCreate)
	api.Patch("/providers/:id", ProvidersUpdate)
	api.Delete("/providers/:id", ProvidersDestroy)
}
