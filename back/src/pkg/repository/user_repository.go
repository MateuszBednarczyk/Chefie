package repository

import (
	"back/src/pkg/db"
	"back/src/pkg/models"
)

func SaveUser(user *models.User) error {
	result := db.Db().Create(&user)
	return result.Error
}

func SelectUserByUsername(username string) *models.User {
	var user models.User
	result := db.Db().Where("username = ?", username).Find(&user)
	if result.Error != nil {
		return nil
	}

	return &user
}

func IsUsernameAlreadyTaken(username string) bool {
	var user models.User
	_ = db.Db().Where("username = ?", username).Find(&user)
	if user.Username == "" {
		return false
	}
	return true
}
