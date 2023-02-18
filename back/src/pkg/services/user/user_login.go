package user

import (
	"back/src/pkg/dto"
	"back/src/pkg/repository"
	"back/src/pkg/services"
	"back/src/pkg/services/http"
	"golang.org/x/crypto/bcrypt"
)

type ILoginService interface {
	LoginUser(dto *dto.Credentials) *http.ServiceResponse
}

type loginService struct {
}

func NewLoginService() *loginService {
	return &loginService{}
}

func (s *loginService) LoginUser(dto *dto.Credentials) *http.ServiceResponse {
	service := services.JwtService()
	result := repository.SelectUserByUsername(dto.Username)

	if result == nil {
		return http.NewServiceResponse("Wrong password or user does not exist", 404, []interface{}{})
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), []byte(dto.Password))
	if err != nil {
		return http.NewServiceResponse("Wrong password or user does not exist", 403, []interface{}{})
	}

	token, err := service.CreateJWT(result.Username)
	if err != nil {
		return http.NewServiceResponse("Couldn't create jwt token", 500, []interface{}{})
	}

	return http.NewServiceResponse("JWT Created", 200, []interface{}{&token})
}
