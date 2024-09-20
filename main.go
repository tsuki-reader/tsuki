package main

import (
	"embed"
	"io/fs"

	"tsuki/core"
	"tsuki/database"
	"tsuki/handlers"
	"tsuki/middleware"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
)

//go:embed web/*
var webFS embed.FS

func main() {
	core.SetupConfig()

	database.Connect()
	database.DATABASE.AutoMigrate(&models.Account{})

	app := fiber.New()

	web, err := fs.Sub(webFS, "web")
	if err != nil {
		core.CONFIG.Logger.Fatal(err)
	}

	middleware.RegisterMiddleware(app, web)
	handlers.RegisterRoutes(app)

	err = app.Listen(core.CONFIG.GetServerAddress())
	core.CONFIG.Logger.Fatal(err)
}
