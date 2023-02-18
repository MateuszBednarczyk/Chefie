package handlers

import (
	"back/src/pkg/dto"
	"back/src/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IUserLoginRegisterHandler interface {
	Register(c echo.Context) error
}

type userLoginRegisterHandler struct {
}

func NewUserHandler() *userLoginRegisterHandler {
	return &userLoginRegisterHandler{}
}

func (h *userLoginRegisterHandler) Register(c echo.Context) error {
	service := services.RegisterService()
	var result dto.Register
	err := c.Bind(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't read serviceResponse from json")
	}

	serviceResponse := service.Register(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Couldn't save serviceResponse")
	}

	return c.JSON(serviceResponse.Code, serviceResponse)
}

func Login(c echo.Context) error {
	service := services.LoginService()
	var credentials dto.Credentials
	err := c.Bind(&credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Couldn't read from DTO")
	}

	serviceResponse := service.LoginUser(&credentials)
	if serviceResponse.Content[0] == "" || len(serviceResponse.Content) == 0 {
		return c.JSON(serviceResponse.Code, serviceResponse)
	}

	return c.JSON(serviceResponse.Code, serviceResponse)
}
