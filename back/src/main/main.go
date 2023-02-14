package main

import (
	"back/src/pkg/db"
	"back/src/pkg/handlers"
	"github.com/labstack/echo/v4"
)

var (
	server     = "localhost"
	port       = "8000"
	apiVersion = "v1"
	dbUsername = "root"
	dbPassword = ""
	dbPort     = "5432"
	dbHost     = "localhost"
	dbName     = "foodie"
)

func main() {
	db.Init(&db.Config{
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbPort:     dbPort,
		DbHost:     dbHost,
		DbName:     dbName,
	})
	e := echo.New()
	initializeHandlers(e)
	e.Logger.Fatal(e.Start(server + ":" + port))
}

func initializeHandlers(e *echo.Echo) {
	e.GET("api/"+apiVersion+"/check", handlers.HealthCheck)
	e.POST("api/"+apiVersion+"/register", handlers.Register)
}
