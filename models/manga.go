package models

import (
	"strings"

	"gorm.io/gorm"
)

type Manga struct {
	gorm.Model
	AnilistID   int `gorm:"unique"`
	Provider    string
	ProviderID  string
	Name        string
	Romaji      string
	Progress    int
	Chapters    int
	Status      string
	List        string
	Description string
	Score       int
	StartYear   int
	StartMonth  int
	Genres      string
	MediumCover string
	LargeCover  string
	BannerImage string
}

func (m *Manga) ParsedGenres() []string {
	g := strings.Split(m.Genres, ",")
	var genres []string = []string{}
	for _, genre := range g {
		if genre != "" {
			genres = append(genres, genre)
		}
	}

	return genres
}
