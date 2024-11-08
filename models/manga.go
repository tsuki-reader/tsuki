package models

type MangaMapping struct {
	ID                  uint              `json:"id" gorm:"primarykey"`
	AnilistID           int               `json:"anilistId" gorm:"uniqueIndex:idx_anilistid_account"`
	ExternalID          string            `json:"externalId"`
	Progress            int               `json:"progress"`
	Chapters            int               `json:"chapters"`
	InstalledProviderID uint              `json:"-"` // TODO: Set null on delete
	InstalledProvider   InstalledProvider `json:"installedProvider"`
	AccountID           uint              `json:"-" gorm:"uniqueIndex:idx_anilistid_account"`
	Account             Account           `json:"-"`
}
