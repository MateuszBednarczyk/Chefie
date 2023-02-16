package handlers

import (
	"back/src/pkg/dto"
	"back/src/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HuserHandler interface {
	Register(c echo.Context) error
}

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) Register(c echo.Context) error {
	service := services.GetRegisterService()
	var result dto.Register
	err := c.Bind(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't read user from json")
	}
	user, err := service.Register(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Couldn't save user")
	}
	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	service := services.GetLoginService()
	var credentials dto.Credentials
	err := c.Bind(&credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Couldn't read from DTO")
	}
	jwt := service.LoginUser(&credentials)
	return c.JSON(http.StatusOK, jwt)
}
