package middlewares

import (
	"back/src/pkg/services"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("Authorization").(*jwt.Token)
		if !ok {
			return echo.ErrUnauthorized
		}
		claims, ok := token.Claims.(services.JwtClaims)
		if !ok {
			return echo.ErrUnauthorized
		}
		if claims.IsAdmin == true {
			return nil
		}
		return nil
	}
}
