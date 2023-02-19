package middlewares

import (
	"back/src/pkg/handlers"
	"back/src/pkg/services"
	"github.com/labstack/echo/v4"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		serviceResponse := services.JwtService().IsTokenValid(authHeader)

		if len(serviceResponse.Content) == 0 {
			return c.JSON(serviceResponse.Code, handlers.NewHandlerResponse(serviceResponse))
		}

		if serviceResponse.Content[0] == false {
			return c.JSON(serviceResponse.Code, handlers.NewHandlerResponse(serviceResponse))
		}

		if serviceResponse.Content[1].(*services.JwtClaims).IsAdmin == true {
			return next(c)
		}

		return c.JSON(serviceResponse.Code, handlers.NewHandlerResponse(serviceResponse))
	}
}
