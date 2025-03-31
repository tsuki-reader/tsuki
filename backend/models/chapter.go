package models

type Chapter struct {
	Base
	Title               string            `json:"title"`
	ExternalID          string            `json:"external_id"`
	Provider            string            `json:"provider"`
	AbsoluteNumber      int               `json:"absolute_number"`
	InstalledProviderID uint              `json:"-"` // TODO: Set null on delete
	InstalledProvider   InstalledProvider `json:"installed_provider"`
}
