package user

import (
	"back/src/pkg/dto"
	"back/src/pkg/models"
	"back/src/pkg/repository"
	"back/src/pkg/services/http"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

type IRegisterService interface {
	Register(dto *dto.Register) *http.ServiceResponse
}

type registerService struct {
}

func NewRegisterService() *registerService {
	return &registerService{}
}

func (s *registerService) Register(dto *dto.Register) *http.ServiceResponse {
	var err error

	plainPassword := dto.Password
	if !isUserValid(dto) {
		return http.NewServiceResponse("Username and password cannot be null or blank", 400, []interface{}{})
	}

	passwordHash, err := hashPassword(plainPassword)
	user := models.User{
		Username:     dto.Username,
		PasswordHash: string(passwordHash),
	}

	log.Println(err)

	result := repository.SaveUser(&user)
	if result != nil {
		return http.NewServiceResponse("Couldn't save user", 500, []interface{}{})
	}
	return http.NewServiceResponse("Account has been created", 200, []interface{}{&user})
}

func isUserValid(dto *dto.Register) bool {
	return len(strings.TrimSpace(dto.Username)) > 0 && len(strings.TrimSpace(dto.Password)) >= 8
}

func hashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}
