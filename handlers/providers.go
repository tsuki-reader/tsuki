package handlers

import (
	"errors"
	"strings"
	"tsuki/extensions"

	"github.com/gofiber/fiber/v2"
	"github.com/tsuki-reader/nisshoku/providers"
)

type params struct {
	RepositoryId string `json:"repository_id"`
	ProviderType string `json:"provider_type"`
	ProviderId   string `json:"provider_id"`
}

// /api/providers = GET = Index

// /api/providers = POST = Create
func ProvidersCreate(c *fiber.Ctx) error {
	// Body will need to include repositoryId and providerType
	body := new(params)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseError{
			Error: "An error occurred. Ensure that you are including the repository id and provider type in the JSON body.",
		})
	}

	// Get the ProviderType
	providerType, err := parseProviderType(body.ProviderType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{
			Error: "Could not install provider: " + err.Error(),
		})
	}

	var repository extensions.Repository
	_, err = extensions.GetRepository(body.RepositoryId, &repository)
	if err != nil || repository.ID == "" {
		return c.Status(fiber.StatusNotFound).JSON(&ResponseError{
			Error: "Repository not found.",
		})
	}

	_, err = extensions.InstallProvider(repository, body.ProviderId, providerType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{
			Error: "An error occurred when installing the provider: " + err.Error(),
		})
	}

	providers := repository.GetProviders(providerType)

	return c.JSON(providers)
}

// /api/providers/1 = GET = Show

// /api/providers/1 = PATCH = Update

// /api/providers/1 = DELETE = Destroy

// Private

func parseProviderType(providerTypeStr string) (providers.ProviderType, error) {
	var providerType providers.ProviderType
	switch strings.ToLower(providerTypeStr) {
	case "comics":
		providerType = providers.Comic
	case "manga":
		providerType = providers.Manga
	default:
		return "", errors.New("invalid provider type")
	}
	return providerType, nil
}
