package models

type MangaMapping struct {
	ID                  uint              `json:"id" gorm:"primarykey"`
	AnilistID           int               `json:"anilistId" gorm:"unique"`
	ExternalID          string            `json:"externalId"`
	Progress            int               `json:"progress"`
	Chapters            int               `json:"chapters"`
	InstalledProviderID int               `json:"-"` // TODO: Set null on delete
	InstalledProvider   InstalledProvider `json:"installedProvider"`
}
