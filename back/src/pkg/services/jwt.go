package services

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type IJWTService interface {
	CreateJWT(username string) (string, error)
}

type jwtService struct {
}

type JwtClaims struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) CreateJWT(username string) (string, error) {
	claims := JwtClaims{
		username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return "Bearer " + token, err
}
