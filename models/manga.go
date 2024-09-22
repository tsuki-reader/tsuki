package models

import "gorm.io/gorm"

type Manga struct {
	gorm.Model
	AnilistID   int
	Provider    string
	Name        string
	Romaji      string
	Progress    int
	Chapters    int
	Status      string
	Description string
	Score       int
	StartYear   int
	StartMonth  int
	Genres      string
}
