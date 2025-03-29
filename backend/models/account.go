package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Base
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

	// TODO: Implement token signing with a secret key
	signedToken, err := token.SignedString([]byte("iamarandomsecretkey"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Authenticate(username string, password string) (*Account, error) {
	account := &Account{}
	err := DATABASE.Where("username = ?", username).First(account).Error
	if err != nil {
		return nil, err
	}

	err = ComparePassword(account.Password, password)
	if err != nil {
		return nil, err
	}

	return account, nil
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
