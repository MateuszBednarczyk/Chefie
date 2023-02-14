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
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Couldn't read user from json")
	}
	user, err := services.RegisterUser(&result)
	if err != nil {
		return c.JSON(http.StatusConflict, "Couldn't save user")
	}
	return c.JSON(http.StatusOK, user)
}
