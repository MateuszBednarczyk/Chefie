package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"back/src/pkg/db"
	"back/src/pkg/handlers"
	"back/src/pkg/middlewares"
	"back/src/pkg/services"
)

var (
	server     = "localhost"
	port       = "8000"
	apiVersion = "v1"
	dbUsername = "root"
	dbPassword = ""
	dbPort     = "5432"
	dbHost     = "localhost"
	dbName     = "Chefie"
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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	services.InitializeServices()
	initializeHandlers(e)
	e.Logger.Fatal(e.Start(server + ":" + port))
}

func initializeHandlers(e *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheck()
	userHandler := handlers.NewUserHandler()
	refreshTokenHandler := handlers.NewRefreshTokenHandler()

	g := e.Group("api/" + apiVersion + "/check")
	g.Use(echojwt.JWT([]byte("secret")))

	e.GET("api/"+apiVersion+"/check", healthCheckHandler.HealthCheck, middlewares.JwtMiddleware)
	e.POST("api/"+apiVersion+"/register", userHandler.Register)
	e.POST("api/"+apiVersion+"/login", userHandler.Login)
	e.POST("api/"+apiVersion+"/token/refresh", refreshTokenHandler.Refresh)

}
