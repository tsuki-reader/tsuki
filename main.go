package main

import (
	"embed"
	"io/fs"

	"tsuki/core"
	"tsuki/database"
	"tsuki/handlers"
	"tsuki/jobs"
	"tsuki/middleware"

	"github.com/gofiber/fiber/v2"
)

//go:embed web/*
var webFS embed.FS

func main() {
	core.SetupConfig()

	database.Connect()
	database.Migrate()

	app := fiber.New()

	web, err := fs.Sub(webFS, "web")
	if err != nil {
		core.CONFIG.Logger.Fatal(err)
	}

	jobs.Run()

	middleware.RegisterMiddleware(app, web)
	handlers.RegisterRoutes(app)

	err = app.Listen(core.CONFIG.GetServerAddress())
	core.CONFIG.Logger.Fatal(err)
}
