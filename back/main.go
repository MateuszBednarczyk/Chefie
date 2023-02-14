package main

import (
	"back/handlers"
	"github.com/labstack/echo/v4"
)

var (
	server = "localhost"
	port   = "8000"
)

var apiVersion = "v1"

func main() {
	e := echo.New()
	initializeHandlers(e)
	e.Logger.Fatal(e.Start(server + ":" + port))
}

func initializeHandlers(e *echo.Echo) {
	e.GET("api/"+apiVersion+"/check", handlers.HealthCheck)
}
