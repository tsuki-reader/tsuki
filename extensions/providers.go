package extensions

import (
	"io"
	"path/filepath"
	"strings"
	"tsuki/core"
	"tsuki/helpers"

	"github.com/tsuki-reader/nisshoku/providers"
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

func InstallProvider(repository Repository, providerId string, providerType providers.ProviderType, update bool, provider *Provider) error {
	// Get the provider from the repository. If it doesn't exist, return an error
	foundProvider, err := repository.GetProvider(providerId, providerType)
	if err != nil {
		return err
	}
	providerUrl := foundProvider.File

	// Download the provider
	response, err := helpers.SendRequest(providerUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// TODO: Use yaegi to sanity check the golang string
	goScript := string(responseBytes)

	// Write the go script to the location comics.tsuki-repo.filesystem.go
	internalId := foundProvider.BuildInternalProviderId(providerType)
	providerDirectory := filepath.Join(core.CONFIG.Directories.Providers, internalId)
	providerFile := filepath.Join(providerDirectory, "provider.go")

	// TODO: Check if provider exits
	err = helpers.CreateDirectory(providerDirectory)
	if err != nil {
		return err
	}

	err = helpers.CreateAndWriteToFile(providerFile, goScript)
	if err != nil {
		return err
	}

	return nil
}
