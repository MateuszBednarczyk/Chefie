package services

import (
	"back/src/pkg/dto"
	"back/src/pkg/models"
	"back/src/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type IRegisterService interface {
	Register(dto *dto.Register) *ServiceResponse
}

type registerService struct {
}

func NewRegisterService() *registerService {
	return &registerService{}
}

func (s *registerService) Register(dto *dto.Register) *ServiceResponse {
	plainPassword := dto.Password
	if !isUserValid(dto) {
		return NewServiceResponse("Username and password cannot be null or blank", 400, []interface{}{})
	}

	passwordHash, err := hashPassword(plainPassword)
	if err != nil {
		return NewServiceResponse("Couldn't hash password", 500, []interface{}{})
	}
	user := models.User{
		Username:     dto.Username,
		PasswordHash: string(passwordHash),
	}

	if repository.IsUsernameAlreadyTaken(user.Username) == true {
		return NewServiceResponse("Username "+user.Username+" is already taken", 409, []interface{}{})
	}

	result := repository.SaveUser(&user)
	if result != nil {
		return NewServiceResponse("Couldn't save user", 500, []interface{}{})
	}
	return NewServiceResponse("Account has been created", 200, []interface{}{&user})
}

func isUserValid(dto *dto.Register) bool {
	return len(strings.TrimSpace(dto.Username)) > 0 && len(strings.TrimSpace(dto.Password)) >= 8
}

func hashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}
