package models

import "gorm.io/gorm"

type InstalledProvider struct {
	gorm.Model
	Name         string `json:"name"`
	ProviderId   string `json:"id" gorm:"uniqueIndex:idx_repo_provider"`
	FileLocation string `json:"file"`
	Icon         string `json:"icon"`
	RepositoryId string `json:"repository_id" gorm:"uniqueIndex:idx_repo_provider"`
	ProviderType string `json:"provider_type" gorm:"uniqueIndex:idx_repo_provider"`
}
