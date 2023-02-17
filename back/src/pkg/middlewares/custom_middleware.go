package middlewares

import (
	"back/src/pkg/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		log.Println(authHeader)
		if authHeader == "" {
			return echo.ErrUnauthorized
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.ErrUnauthorized
		}
		token, err := jwt.ParseWithClaims(tokenString, &services.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			return echo.ErrUnauthorized
		}
		if !token.Valid {
			return echo.ErrUnauthorized
		}
		claims, ok := token.Claims.(*services.JwtClaims)
		if !ok {
			return echo.ErrUnauthorized
		}
		if claims.IsAdmin {
			return next(c)
		}
		return echo.ErrForbidden
	}
}
