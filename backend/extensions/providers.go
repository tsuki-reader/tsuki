package extensions

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"tsuki/backend/config"
	"tsuki/backend/helpers"
	"tsuki/backend/models"

	"github.com/tsuki-reader/nisshoku/providers"
	"gorm.io/gorm/clause"
)

// Types

type Provider struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	File      string `json:"file"`
	Icon      string `json:"icon"`
	Installed bool   `json:"installed"`
}

// Public

func InstallProvider(repository Repository, providerId string, providerType providers.ProviderType) (*models.InstalledProvider, error) {
	// Get the provider from the repository. If it doesn't exist, return an error
	foundProvider, err := repository.GetProvider(providerId, providerType)
	if err != nil {
		return nil, err
	}
	providerUrl := foundProvider.File

	// Download the provider
	response, err := helpers.SendRequest(providerUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	goScript := string(responseBytes)

	// TODO: Evaluate the provider script using yaegi_interp.EvaluateProvider(goScript)
	// _, err = yaegi_interp.EvaluateProvider(goScript)
	// if err != nil {
	// 	return nil, err
	// }

	// Write the go script to the location filesystem.go
	internalId := repository.BuildInternalProviderId(*foundProvider, providerType)
	providerFile := filepath.Join(config.CONFIG.Directories.Providers, internalId+".go")

	err = helpers.CreateAndWriteToFile(providerFile, goScript)
	if err != nil {
		return nil, err
	}

	// Insert provider into database
	clause := clause.OnConflict{
		Columns:   []clause.Column{{Name: "repository_id"}, {Name: "provider_id"}, {Name: "provider_type"}},
		UpdateAll: true,
	}
	installedProvider := models.InstalledProvider{
		Name:         foundProvider.Name,
		ProviderId:   providerId,
		FileLocation: providerFile,
		Icon:         foundProvider.Icon,
		RepositoryId: repository.ID,
		ProviderType: string(providerType),
	}
	if err := models.DATABASE.Clauses(clause).Create(&installedProvider).Error; err != nil {
		return nil, err
	}

	return &installedProvider, nil
}

func GetInstalledProvider(repository Repository, providerId string, providerType providers.ProviderType) (*models.InstalledProvider, error) {
	provider := models.InstalledProvider{}
	err := models.DATABASE.Where(&models.InstalledProvider{
		RepositoryId: repository.ID,
		ProviderId:   providerId,
		ProviderType: string(providerType),
	}).First(&provider).Error
	return &provider, err
}

func GetInstalledProviders(providerType providers.ProviderType) ([]models.InstalledProvider, error) {
	providers := []models.InstalledProvider{}
	err := models.DATABASE.Where(&models.InstalledProvider{ProviderType: string(providerType)}).Find(&providers).Error
	return providers, err
}

func UninstallProvider(provider models.InstalledProvider) error {
	// Purposefully ignore the returned error when removing file. If the file can't be removed/if it doesn't exist we don't
	// want that from stopping the uninstall from happening.
	err := os.Remove(provider.FileLocation)
	if err != nil {
		log.Println("Provider Uninstall: ", err)
	}

	err = models.DATABASE.Unscoped().Delete(&provider).Error
	return err
}
