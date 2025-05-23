package extensions

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"tsuki/backend/config"
	"tsuki/backend/helpers"
	"tsuki/backend/models"

	"github.com/tsuki-reader/nisshoku/providers"
)

// Types

type Repository struct {
	Name           string      `json:"name"`
	ID             string      `json:"id"`
	Logo           string      `json:"logo"`
	URL            string      `json:"url"`
	MangaProviders []*Provider `json:"manga_providers"`
	ComicProviders []*Provider `json:"comic_providers"`
}

func (r *Repository) Update() error {
	oldRepoId := r.ID
	err := InstallRepository(r.URL, true, r)

	if err == nil && oldRepoId != "" && oldRepoId != r.ID {
		UninstallRepository(oldRepoId)
	}

	return err
}

func (r *Repository) Uninstall() error {
	err := UninstallRepository(r.ID)
	return err
}

func (r *Repository) GetProvider(providerId string, providerType providers.ProviderType) (*Provider, error) {
	var _providers []*Provider

	switch providerType {
	case providers.Comic:
		_providers = r.ComicProviders
	case providers.Manga:
		_providers = r.MangaProviders
	}

	foundProviderIdx := slices.IndexFunc(_providers, func(p *Provider) bool { return p.ID == providerId })
	if foundProviderIdx == -1 {
		return nil, errors.New("Could not find provider with ID " + providerId)
	}

	result := _providers[foundProviderIdx]
	result.Installed = models.RecordExists(
		&models.InstalledProvider{ProviderId: providerId, RepositoryId: r.ID, ProviderType: string(providerType)},
		&models.InstalledProvider{},
	)

	return result, nil
}

func (r *Repository) GetProviders(providerType providers.ProviderType) []*Provider {
	_providers := []*Provider{}

	switch providerType {
	case providers.Comic:
		_providers = r.ComicProviders
	case providers.Manga:
		_providers = r.MangaProviders
	}

	for _, p := range _providers {
		p.Installed = models.RecordExists(
			&models.InstalledProvider{ProviderId: p.ID, RepositoryId: r.ID, ProviderType: string(providerType)},
			&models.InstalledProvider{},
		)
	}

	return _providers
}

func (r *Repository) BuildInternalProviderId(provider Provider, providerType providers.ProviderType) string {
	_providerType := strings.ToLower(string(providerType))
	return _providerType + "." + r.ID + "." + provider.ID
}

// Public

func InstallRepository(jsonUrl string, update bool, repository *Repository) error {
	response, err := helpers.SendRequest(jsonUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(repository)
	if err != nil {
		return err
	}
	repository.URL = jsonUrl

	if repository.Name == "" || repository.ID == "" {
		return errors.New("Repository did not provide necessary information")
	}

	if !validateRepositoryId(repository.ID) {
		return errors.New("Repository ID failed validation check")
	}

	bytes, err := json.Marshal(repository)
	if err != nil {
		return err
	}
	repositoryJson := string(bytes)

	// Now we want to store the json file in the repositories directory
	repositoryFilename := repository.ID + ".json"
	repositoryLocation := filepath.Join(config.CONFIG.Directories.Repositories, repositoryFilename)

	// Check if the repository already exists
	repositoryExists := helpers.FileExists(repositoryLocation)
	if repositoryExists && !update {
		return errors.New("Repository already exists")
	}

	err = helpers.CreateAndWriteToFile(repositoryLocation, repositoryJson)
	if err != nil {
		return err
	}

	return nil
}

// Returns the location, and an optional error, and mutates given repository.
// If the repository could not be found, it returns an empty string and a nil error.
func GetRepository(repositoryId string, repository *Repository) (string, error) {
	// Build the repository location
	repositoryFile := repositoryId + ".json"
	repositoryLocation := filepath.Join(config.CONFIG.Directories.Repositories, repositoryFile)

	// Check if the file exists
	repositoryExists := helpers.FileExists(repositoryLocation)
	if !repositoryExists {
		return "", nil
	}

	// Read the file
	repositoryJson, err := helpers.ReadFileContents(repositoryLocation)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(repositoryJson), repository)
	if err != nil {
		return "", err
	}

	// TODO: Decide. Do we want to do this or make the frontend request the providers.
	for _, p := range repository.MangaProviders {
		p.Installed = models.RecordExists(
			&models.InstalledProvider{RepositoryId: repositoryId, ProviderType: string(providers.Manga), ProviderId: p.ID},
			&models.InstalledProvider{},
		)
	}

	return repositoryLocation, nil
}

func GetRepositories() ([]Repository, error) {
	var repositories []Repository

	err := filepath.WalkDir(config.CONFIG.Directories.Repositories, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d == nil {
			return errors.New("there was an error walking")
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".json" {
			return nil
		}

		repositoryId := strings.SplitN(d.Name(), ".", 2)[0]
		repository := Repository{}
		_, err = GetRepository(repositoryId, &repository)
		if err != nil || repository.ID == "" {
			return errors.New("repository error")
		}
		repositories = append(repositories, repository)

		return nil
	})

	if repositories == nil {
		return []Repository{}, nil
	}

	return repositories, err
}

func UninstallRepository(repositoryId string) error {
	repositoryLocation := filepath.Join(config.CONFIG.Directories.Repositories, repositoryId+".json")

	err := os.Remove(repositoryLocation)
	return err
}

// Private

func validateRepositoryId(repositoryId string) bool {
	matched, _ := regexp.MatchString("[^a-zA-Z\\-\\d]", repositoryId)
	return !matched
}
