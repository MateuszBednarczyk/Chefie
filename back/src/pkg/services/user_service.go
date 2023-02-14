package services

import (
	"back/src/pkg/db"
	"back/src/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func RegisterUser(u *models.User) (*models.User, error) {
	var err error
	plainPassword := u.PasswordHash
	passwordBytes, err := HashPassword(plainPassword)
	u.PasswordHash = string(passwordBytes)
	db.DB.Create(&u)
	if err != nil {
		return nil, err
	}
	return u, err
}
