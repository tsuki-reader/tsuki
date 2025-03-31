package backend

import (
	"errors"
	"strings"
	"tsuki/backend/extensions"
	"tsuki/backend/models"

	"github.com/tsuki-reader/nisshoku/providers"
)

func (a *App) ProvidersIndex(repositoryId string, providerType string) ([]models.InstalledProvider, error) {
	pType, err := parseProviderType(providerType)
	if err != nil {
		return []models.InstalledProvider{}, err
	}

	query := models.InstalledProvider{ProviderType: string(pType)}
	if repositoryId != "" { // Repository ID is optional, it will only return providers for the specified repository if provided
		query.RepositoryId = repositoryId
	}

	var providers []models.InstalledProvider
	if err := models.DATABASE.Where(&query).Find(&providers).Error; err != nil {
		return []models.InstalledProvider{}, err
	}

	return providers, nil
}

func (a *App) ProvidersCreateOrUpdate(repositoryId string, providerId string, providerType string) ([]*extensions.Provider, error) {
	return installOrUpdateProvider(repositoryId, providerId, providerType)
}

func (a *App) ProvidersDestroy(repositoryId string, providerId string, providerType string) ([]*extensions.Provider, error) {
	pType, err := parseProviderType(providerType)
	if err != nil {
		return []*extensions.Provider{}, err
	}

	repository, err := getRepository(repositoryId)
	if err != nil {
		return []*extensions.Provider{}, err
	}

	provider := models.InstalledProvider{}
	err = models.DATABASE.
		Where(&models.InstalledProvider{ProviderId: providerId, RepositoryId: repositoryId, ProviderType: string(pType)}).
		First(&provider).
		Error
	if err != nil {
		return []*extensions.Provider{}, err
	}

	err = extensions.UninstallProvider(provider)
	if err != nil {
		return []*extensions.Provider{}, err
	}

	providers := repository.GetProviders(pType)
	return providers, nil
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
		return "", errors.New("Invalid provider type")
	}
	return providerType, nil
}

func installOrUpdateProvider(repositoryId string, providerId string, providerType string) ([]*extensions.Provider, error) {
	// Get the ProviderType
	pType, err := parseProviderType(providerType)
	if err != nil {
		return []*extensions.Provider{}, err
	}

	repository, err := getRepository(repositoryId)
	if err != nil {
		return []*extensions.Provider{}, err
	}

	_, err = extensions.InstallProvider(repository, providerId, pType)
	if err != nil {
		return []*extensions.Provider{}, err
	}

	providers := repository.GetProviders(pType)
	return providers, nil
}

func getRepository(repositoryId string) (extensions.Repository, error) {
	repository := extensions.Repository{}
	_, err := extensions.GetRepository(repositoryId, &repository)
	if err != nil || repository.ID == "" {
		return repository, errors.New("repository not found")
	}

	return repository, nil
}
