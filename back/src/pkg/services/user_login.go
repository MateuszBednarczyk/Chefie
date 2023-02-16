package services

import (
	"back/src/pkg/db"
	"back/src/pkg/dto"
	"back/src/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type ILoginService interface {
	LoginUser(dto *dto.Credentials) string
}

type loginService struct {
}

func NewLoginService() *loginService {
	return &loginService{}
}

func (s *loginService) LoginUser(dto *dto.Credentials) string {
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
