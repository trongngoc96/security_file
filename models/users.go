package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              uint    `json:"id"`
	EMAIL           string  `json:"email"`
	PASSWORD        string  `json:"password"`
	USERNAME        string  `json:"username"`
	KEYSTORE        string  `json:"keystore"`
	KEYSTORE_ADMIN  string  `json:"keystore_admin"`
	ADDRESS         string  `json:"address"`
	PASSWORD_WALLET string  `json:"password_wallet"`
	BALANCE         float64 `json:"balance"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.PASSWORD = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
