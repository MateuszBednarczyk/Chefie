package handlers

import (
	"back/src/pkg/dto"
	"back/src/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) error {
	var result dto.Register
	err := c.Bind(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't read user from json")
	}
	user, err := services.RegisterUser(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Couldn't save user")
	}
	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	var credentials dto.Credentials
	err := c.Bind(&credentials)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Couldn't read from DTO")
	}
	jwt := services.LoginUser(&credentials)
	return c.JSON(http.StatusOK, jwt)
}
