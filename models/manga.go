package models

type MangaMapping struct {
	ID         uint `json:"id" gorm:"primarykey"`
	AnilistID  int  `gorm:"unique"`
	ProviderID string
	ExternalID string
	Progress   int
	Chapters   int
}
