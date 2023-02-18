package repository

import (
	"back/src/pkg/db"
	"back/src/pkg/models"
)

func SaveUser(user *models.User) error {
	result := db.Db().Create(&user).Error
	return result
}

func SelectUserByUsername(username string) *models.User {
	var user models.User
	result := db.Db().Where("username = ?", username).Find(&user).Error
	if result.Error != nil {
		return nil
	}
	return &user
}
