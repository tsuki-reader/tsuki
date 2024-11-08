package models

import (
	"tsuki/extensions/yaegi_interp"
	"tsuki/helpers"

	"github.com/tsuki-reader/nisshoku/providers"
)

type InstalledProvider struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name"`
	ProviderId   string `json:"providerId" gorm:"uniqueIndex:idx_repo_provider"`
	FileLocation string `json:"-"`
	Icon         string `json:"-"`
	RepositoryId string `json:"repositoryId" gorm:"uniqueIndex:idx_repo_provider"`
	ProviderType string `json:"providerType" gorm:"uniqueIndex:idx_repo_provider"`
}

func (ip *InstalledProvider) ScriptContents() (string, error) {
	return helpers.ReadFileContents(ip.FileLocation)
}

func (ip *InstalledProvider) Load() (providers.Provider, error) {
	script, err := ip.ScriptContents()
	if err != nil {
		return nil, err
	}

	value, err := yaegi_interp.EvaluateProvider(script)
	if err != nil {
		return nil, err
	}

	NewProvider := value.Interface().(func(providers.ProviderContext) providers.Provider)

	// TODO: Pull comic and manga library out of settings
	context := providers.ProviderContext{
		MangaLibraryPath: "/home/crxssed/Downloads/",
		ComicLibraryPath: "/home/crxssed/Downloads/",
		ProviderType:     providers.ProviderType(ip.ProviderType),
	}

	provider := NewProvider(context)

	return provider, nil
}

// Get the first search result from this provider. Only returns nil in cases where the provider fails to load.
func (ip *InstalledProvider) GetFirstSearchResult(query string) (*providers.ProviderResult, error) {
	provider, err := ip.Load()
	if err != nil {
		return nil, err
	}

	searchResults, err := provider.Search(query)
	if len(searchResults) > 0 {
		return &searchResults[0], nil
	}

	return nil, nil
}

func (ip *InstalledProvider) GetChapterList(externalId string) ([]providers.Chapter, error) {
	if externalId == "" {
		return []providers.Chapter{}, nil
	}

	provider, err := ip.Load()
	if err != nil {
		return []providers.Chapter{}, err
	}

	chapters, err := provider.GetChapters(externalId)
	if err != nil {
		return []providers.Chapter{}, err
	}

	return chapters, nil
}
