package main

import (
	"back/src/pkg/db"
	"back/src/pkg/handlers"
	"back/src/pkg/services"
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
	services.InitializeServices()
	initializeHandlers(e)
	e.Logger.Fatal(e.Start(server + ":" + port))
}

func initializeHandlers(e *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheck()
	userHandler := handlers.NewUserHandler()

	e.GET("api/"+apiVersion+"/check", healthCheckHandler.HealthCheck)
	e.POST("api/"+apiVersion+"/register", userHandler.Register)
	e.POST("api/"+apiVersion+"/login", handlers.Login)
}
