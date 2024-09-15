package main

import (
	"embed"
	"io/fs"
	"log"

	"tsuki/database"
	"tsuki/handlers"
	"tsuki/middleware"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
)

//go:embed web/*
var webFS embed.FS

func main() {
	database.Connect()
	database.DATABASE.AutoMigrate(&models.Account{})

	app := fiber.New()

	web, err := fs.Sub(webFS, "web")
	if err != nil {
		log.Fatal(err)
	}

	middleware.RegisterMiddleware(app, web)
	handlers.RegisterRoutes(app)

	log.Fatal(app.Listen(":1337"))
}
