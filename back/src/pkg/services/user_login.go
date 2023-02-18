package services

import (
	"back/src/pkg/dto"
	"back/src/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type ILoginService interface {
	LoginUser(dto *dto.Credentials) *ServiceResponse
}

type loginService struct {
}

func NewLoginService() *loginService {
	return &loginService{}
}

func (s *loginService) LoginUser(dto *dto.Credentials) *ServiceResponse {
	service := JwtService()
	result := repository.SelectUserByUsername(dto.Username)

	if result == nil {
		return NewServiceResponse("Wrong password or user does not exist", 404, []interface{}{})
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(dto.Password))
	if err != nil {
		return NewServiceResponse("Wrong password or user does not exist", 403, []interface{}{})
	}

	tokens := service.GenerateTokens(result.Username)
	if err != nil {
		return NewServiceResponse("Couldn't create jwt tokens", 500, []interface{}{})
	}

	return NewServiceResponse("JWT Created", 200, []interface{}{&tokens})
}
