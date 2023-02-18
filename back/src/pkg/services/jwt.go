package services

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

type IJWTService interface {
	GenerateTokens(username string) *ServiceResponse
	RefreshToken(rawToken string) *ServiceResponse
}

type jwtService struct {
}

type JwtClaims struct {
	Username string `json:"name"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateTokens(username string) *ServiceResponse {
	baseTokenClaims := JwtClaims{
		username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}

	refreshTokenClaims := JwtClaims{
		username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, baseTokenClaims)
	token, err := rawToken.SignedString([]byte("secret"))
	if err != nil {
		return NewServiceResponse("Couldn't parse user's claims", 500, nil)
	}

	rawRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshToken, err := rawRefreshToken.SignedString([]byte("secret"))
	if err != nil {
		return NewServiceResponse("Invalid secret", 403, nil)

	}

	content := []interface{}{map[string]string{"token": "Bearer " + token, "refresh": "Bearer " + refreshToken}}

	return NewServiceResponse("Tokens generated successfully", 200, content)
}

func (s *jwtService) RefreshToken(rawToken string) *ServiceResponse {
	if rawToken == "" {
		return NewServiceResponse("Given token is empty", 400, []interface{}{})
	}

	tokenString := strings.TrimPrefix(rawToken, "Bearer ")
	if tokenString == rawToken {
		return NewServiceResponse("Given token is incorrect", 400, []interface{}{})
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return NewServiceResponse("Invalid token", 403, []interface{}{})
	}
	if !token.Valid {
		return NewServiceResponse("Invalid token", 403, []interface{}{})
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return NewServiceResponse("Invalid token", 403, []interface{}{})
	}
	tokens := JwtService().GenerateTokens(claims.Username)
	if tokens.Content == nil {
		return NewServiceResponse(tokens.Message, tokens.Code, nil)
	}

	return tokens
}
