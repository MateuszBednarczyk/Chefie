package handlers

import (
	"back/src/pkg/services"
	"github.com/labstack/echo/v4"
)

type IRefreshTokenHandler interface {
	Refresh()
}

type refreshTokenHandler struct {
}

func NewRefreshTokenHandler() *refreshTokenHandler {
	return &refreshTokenHandler{}
}

func (h *refreshTokenHandler) Refresh(c echo.Context) error {
	rawToken := c.Request().Header.Get("Authorization")
	serviceResponse := services.JwtService().RefreshToken(rawToken)

	return c.JSON(serviceResponse.Code, NewHandlerResponse(serviceResponse))
}
