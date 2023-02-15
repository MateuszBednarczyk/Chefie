package services

import (
	"back/src/pkg/db"
	"back/src/pkg/dto"
	"back/src/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	LoginUser(dto *dto.Credentials) string
}

type LoginServiceStruct struct {
}

func (s *LoginServiceStruct) LoginUser(dto *dto.Credentials) string {
	service := GetJWTService()
	var result models.User
	db.DB.Where("username = ?", dto.Username).Find(&result)
	err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(dto.Password))
	if err != nil {
		return ""
	}

	token, err := service.CreateJWT(result.Username)
	if err != nil {
		return "something went wrong while creating jwt"
	}

	return token
}
