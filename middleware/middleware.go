package middleware

import (
	"io/fs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterMiddleware(app *fiber.App, web fs.FS) {
	app.Use(filesystemMiddleware(web))
	app.Use(logger.New())
	app.Use(corsMiddleware())
}
