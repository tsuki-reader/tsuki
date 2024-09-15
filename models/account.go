package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name  string `json:"name"`
	Token string `json:"token"`
}
