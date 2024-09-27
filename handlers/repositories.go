package handlers

import (
	"tsuki/extensions"

	"github.com/gofiber/fiber/v2"
)

func RepositoriesIndex(c *fiber.Ctx) error {
	repositories, err := extensions.GetRepositories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "An error occurred when trying to list repositories: " + err.Error(),
		})
	}

	return c.JSON(repositories)
}
