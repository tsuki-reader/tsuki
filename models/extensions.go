package models

import (
	"tsuki/extensions/yaegi_interp"
	"tsuki/helpers"

	"github.com/tsuki-reader/nisshoku/providers"
	"gorm.io/gorm"
)

type InstalledProvider struct {
	gorm.Model
	Name         string `json:"name"`
	ProviderId   string `json:"id" gorm:"uniqueIndex:idx_repo_provider"`
	FileLocation string `json:"file"`
	Icon         string `json:"icon"`
	RepositoryId string `json:"repository_id" gorm:"uniqueIndex:idx_repo_provider"`
	ProviderType string `json:"provider_type" gorm:"uniqueIndex:idx_repo_provider"`
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
