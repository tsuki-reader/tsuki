package models

import (
	"gorm.io/gorm"
)

type MangaMapping struct {
	gorm.Model
	AnilistID  int `gorm:"unique"`
	Provider   string
	ProviderID string
	Progress   int
	Chapters   int
}
