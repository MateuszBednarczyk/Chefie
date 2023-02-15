package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HhealthCheck interface {
	HealthCheck(c echo.Context) error
}

type healthCheck struct {
}

func NewHealthCheck() HhealthCheck {
	return &healthCheck{}
}

func (h *healthCheck) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK!")
}
