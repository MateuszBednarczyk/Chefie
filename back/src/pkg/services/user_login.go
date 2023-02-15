package services

import (
	"back/src/pkg/db"
	"back/src/pkg/dto"
	"back/src/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(dto *dto.Credentials) string {
	var result models.User
	db.DB.Where("username = ?", dto.Username).Find(&result)
	err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(dto.Password))
	if err != nil {
		return ""
	}

	token, err := CreateJWT(result.Username)
	if err != nil {
		return "something went wrong while creating jwt"
	}

	return token
}
