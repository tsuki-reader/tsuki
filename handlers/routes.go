package handlers

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// ========== Auth
	api.Get("/auth/status", Status)
	api.Post("/auth/login", Login)
	api.Post("/auth/logout", Logout)
}
