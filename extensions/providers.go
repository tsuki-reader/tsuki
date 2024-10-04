package extensions

import (
	"io"
	"path/filepath"
	"strings"
	"tsuki/core"
	"tsuki/database"
	"tsuki/helpers"
	"tsuki/models"

	"github.com/tsuki-reader/nisshoku/providers"
	"gorm.io/gorm/clause"
)

// Types

type Provider struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	File         string `json:"file"`
	Icon         string `json:"icon"`
	RepositoryId string `json:"repository_id"`
}

func (p *Provider) SetRepositoryID(r Repository) {
	p.RepositoryId = r.ID
}

func (p *Provider) BuildInternalProviderId(providerType providers.ProviderType) string {
	_providerType := strings.ToLower(string(providerType))
	return _providerType + "." + p.RepositoryId + "." + p.ID
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
	// TODO: Use yaegi to sanity check the golang string
	goScript := string(responseBytes)

	// Write the go script to the location filesystem.go
	internalId := foundProvider.BuildInternalProviderId(providerType)
	providerFile := filepath.Join(core.CONFIG.Directories.Providers, internalId+".go")

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
	if err := database.DATABASE.Clauses(clause).Create(&installedProvider).Error; err != nil {
		return nil, err
	}

	return &installedProvider, nil
}
