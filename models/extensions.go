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
