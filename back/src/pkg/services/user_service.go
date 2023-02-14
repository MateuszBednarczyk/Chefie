package services

import (
	"back/src/pkg/db"
	"back/src/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(u *models.User) (*models.User, error) {
	var err error
	plainPassword := u.PasswordHash
	passwordBytes, err := hashPassword(plainPassword)
	u.PasswordHash = string(passwordBytes)
	result := db.DB.Create(&u)
	if result.Error != nil {
		return nil, err
	}
	return u, err
}

func hashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}

func Valid(u *models.User) bool {
	if len(u.Username) <= 0 || len(u.PasswordHash) <= 0 {
		return false
	}
	return true
}
