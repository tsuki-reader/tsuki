package handlers

import (
	"errors"
	"strings"
	"tsuki/database"
	"tsuki/extensions"
	"tsuki/models"

	"github.com/gofiber/fiber/v2"
	"github.com/tsuki-reader/nisshoku/providers"
)

type params struct {
	RepositoryId string `json:"repository_id"`
	ProviderType string `json:"provider_type"`
	ProviderId   string `json:"provider_id"`
}

// /api/providers = GET = Index
func ProvidersIndex(c *fiber.Ctx) error {
	providerTypeStr := c.Query("provider_type", "comics")
	repositoryId := c.Query("repository_id")

	providerType, err := parseProviderType(providerTypeStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{Error: "Could not retrieve providers: " + err.Error()})
	}

	query := models.InstalledProvider{ProviderType: string(providerType)}
	if repositoryId != "" {
		query.RepositoryId = repositoryId
	}

	var providers []models.InstalledProvider
	if err := database.DATABASE.Where(&query).Find(&providers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&ResponseError{Error: "Could not retrieve providers: " + err.Error()})
	}

	return c.JSON(providers)
}

// /api/providers = POST = Create
func ProvidersCreate(c *fiber.Ctx) error {
	body := new(params)

	if err := c.BodyParser(body); err != nil || body.RepositoryId == "" || body.ProviderId == "" || body.ProviderType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{Error: "An error occurred. Ensure that you are including the repository id and provider type in the JSON body."})
	}

	return installOrUpdateProvider(*body, c, body.ProviderId)
}

// /api/providers/providerId = GET = Show

// /api/providers/providerId = PATCH = Update
func ProvidersUpdate(c *fiber.Ctx) error {
	providerId := c.Params("id")
	body := new(params)

	if err := c.BodyParser(body); err != nil || body.RepositoryId == "" || body.ProviderType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{Error: "An error occurred. Ensure that you are including the repository id and provider type in the JSON body."})
	}

	return installOrUpdateProvider(*body, c, providerId)
}

// /api/providers/providerId = DELETE = Destroy
func ProvidersDestroy(c *fiber.Ctx) error {
	providerId := c.Params("id")
	body := new(params)

	if err := c.BodyParser(body); err != nil || body.RepositoryId == "" || body.ProviderType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{Error: "An error occurred. Ensure that you are including the repository id and provider type in the JSON body."})
	}

	providerType, err := parseProviderType(body.ProviderType)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&ResponseError{
			Error: "Could not uninstall provider: " + err.Error(),
		})
	}

	repository, err := getRepository(*body)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&ResponseError{
			Error: "Repository not found.",
		})
	}

	provider := models.InstalledProvider{}
	err = database.DATABASE.
		Where(&models.InstalledProvider{ProviderId: providerId, RepositoryId: body.RepositoryId, ProviderType: string(providerType)}).
		First(&provider).
		Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&ResponseError{
			Error: "Provider not found.",
		})
	}

	err = extensions.UninstallProvider(provider)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&ResponseError{
			Error: "Could not uninstall provider: " + err.Error(),
		})
	}

	providers := repository.GetProviders(providerType)
	return c.JSON(providers)
}

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

	repository, err := getRepository(body)
	if err != nil {
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

func getRepository(body params) (extensions.Repository, error) {
	repository := extensions.Repository{}
	_, err := extensions.GetRepository(body.RepositoryId, &repository)
	if err != nil || repository.ID == "" {
		return repository, errors.New("repository not found")
	}

	return repository, nil
}
