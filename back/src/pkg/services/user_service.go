package services

import (
	"back/src/pkg/db"
	"back/src/pkg/dto"
	"back/src/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type RegisterService interface {
	Register(dto *dto.Register) (*models.User, error)
}

type RegisterServiceStruct struct {
}

func (s *RegisterServiceStruct) Register(dto *dto.Register) (*models.User, error) {
	var err error

	plainPassword := dto.Password
	if !isUserValid(dto) {
		return nil, err
	}

	passwordHash, err := hashPassword(plainPassword)
	user := models.User{
		Username:     dto.Username,
		PasswordHash: string(passwordHash),
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		return nil, nil
	}

	return &user, err
}

func isUserValid(dto *dto.Register) bool {
	return len(strings.TrimSpace(dto.Username)) > 0 && len(strings.TrimSpace(dto.Password)) >= 8
}

func hashPassword(p string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
}
