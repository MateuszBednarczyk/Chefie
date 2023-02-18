package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type IHealthCheckHandler interface {
	HealthCheck(c echo.Context) error
}

type healthCheckHandler struct {
}

func NewHealthCheck() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) HealthCheck(c echo.Context) error {

	return c.JSON(http.StatusOK, "OK!")
}
