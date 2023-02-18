package repository

import (
	"back/src/pkg/db"
	"back/src/pkg/models"
)

func SaveUser(user *models.User) bool {
	result := db.Db().Create(&user)
	if result.Error != nil {
		return false
	}
	return true
}

func SelectUserByUsername(username string) *models.User {
	var result models.User
	db.Db().Where("username = ?", username).Find(&result)

	return &result
}
