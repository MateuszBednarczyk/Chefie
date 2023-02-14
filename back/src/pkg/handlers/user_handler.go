package handlers

import (
	"back/src/pkg/models"
	"back/src/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) error {
	var result models.User
	err := c.Bind(&result)
	if services.Valid(&result) != true {
		return echo.NewHTTPError(http.StatusBadRequest, "Username and password can't be null")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Couldn't read user from json")
	}
	user, err := services.RegisterUser(&result)
	if err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Couldn't save user")
	}
	return c.JSON(http.StatusOK, user)
}
