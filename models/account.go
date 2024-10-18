package models

import (
	"tsuki/core"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	AnilistToken string `json:"anilist_token"`
	AnilistName  string `json:"anilist_name"`
	Username     string `json:"username" gorm:"unique"`
	Password     string `json:"-"`
}

func (account *Account) GenerateJWTToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       account.ID,
		"username": account.Username,
	})

	signedToken, err := token.SignedString([]byte(core.CONFIG.Server.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashed string, plain string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err
}
