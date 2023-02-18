package services

import (
	"back/src/pkg/dto"
	"back/src/pkg/repository"
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
	service := JwtService()
	result := repository.SelectUserByUsername(dto.Username)

	err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(dto.Password))
	if err != nil {
		return "Wrong password or user does not exist"
	}

	token, err := service.CreateJWT(result.Username)
	if err != nil {
		return "something went wrong while creating jwt"
	}

	return token
}
