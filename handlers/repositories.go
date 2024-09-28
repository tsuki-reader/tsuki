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

func RepositoriesCreate(c *fiber.Ctx) error {
	type _body struct {
		URL string `json:"url"`
	}
	body := new(_body)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "An error occurred. Ensure that you are including the repository url in the JSON body.",
		})
	}

	if err := extensions.InstallRepository(body.URL, false); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "An error occurred when installing repository from url: " + err.Error(),
		})
	}

	repositories, err := extensions.GetRepositories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "An error occurred when trying to list repositories: " + err.Error(),
		})
	}

	return c.JSON(repositories)
}

func RepositoriesDestroy(c *fiber.Ctx) error {
	repositoryId := c.Params("id")

	var repository extensions.Repository
	if _, err := extensions.GetRepository(repositoryId, &repository); err != nil {
		// TODO: What status do we actually want to respond with here?
		return c.Status(fiber.StatusNotFound).JSON(ResponseError{
			Error: "Could not retrieve repository: " + err.Error(),
		})
	}

	if err := repository.Uninstall(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "Could not uninstall repository due to the following: " + err.Error(),
		})
	}

	repositories, err := extensions.GetRepositories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseError{
			Error: "An error occurred when trying to list repositories: " + err.Error(),
		})
	}

	return c.JSON(repositories)
}
