package middleware

import (
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func skipFilesystemMiddleware(c *fiber.Ctx) bool {
	path := c.Path()
	if strings.HasPrefix(path, "/api") {
		return true
	}

	if shouldAddExtension(path) {
		path = appendHTMLExtension(path)
	}

	// Account for when path is the root path
	if path == "/.html" {
		path = "/index.html"
	}

	c.Path(path)
	return false
}

func shouldAddExtension(path string) bool {
	return !strings.HasSuffix(path, ".html") && filepath.Ext(path) == ""
}

func appendHTMLExtension(path string) string {
	if strings.Contains(path, "?") {
		parts := strings.SplitN(path, "?", 2)
		return parts[0] + ".html?" + parts[1]
	}
	return path + ".html"
}

func filesystemMiddleware(web fs.FS) func(*fiber.Ctx) error {
	var config = filesystem.Config{
		Root:   http.FS(web),
		Browse: true,
		Next:   skipFilesystemMiddleware,
	}

	return filesystem.New(config)
}
