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
	body := new(params)

	if err := c.BodyParser(body); err != nil || body.RepositoryId == "" || body.ProviderId == "" || body.ProviderType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{Error: "An error occurred. Ensure that you are including the repository id and provider type in the JSON body."})
	}

	return installOrUpdateProvider(*body, c, body.ProviderId)
}

// /api/providers/provider-id = GET = Show

// /api/providers/provider-id = PATCH = Update
func ProvidersUpdate(c *fiber.Ctx) error {
	providerId := c.Params("id")
	body := new(params)

	if err := c.BodyParser(body); err != nil || body.RepositoryId == "" || body.ProviderType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{Error: "An error occurred. Ensure that you are including the repository id and provider type in the JSON body."})
	}

	return installOrUpdateProvider(*body, c, providerId)
}

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

func installOrUpdateProvider(body params, c *fiber.Ctx, providerId string) error {
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

	_, err = extensions.InstallProvider(repository, providerId, providerType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{
			Error: "An error occurred when installing the provider: " + err.Error(),
		})
	}

	providers := repository.GetProviders(providerType)
	return c.JSON(providers)
}
