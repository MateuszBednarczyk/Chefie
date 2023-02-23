package main

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"sync"

	"back/src/pkg/db"
	"back/src/pkg/handlers"
	"back/src/pkg/middlewares"
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

	var wg sync.WaitGroup
	wg.Add(1)
	go launchServer(&wg)
	wg.Wait()

}

func initializeHandlers(serverInstance *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheck()
	userHandler := handlers.NewUserHandler()
	refreshTokenHandler := handlers.NewRefreshTokenHandler()

	g := serverInstance.Group("api/" + apiVersion + "/check")
	g.Use(echojwt.JWT([]byte("secret")))

	serverInstance.GET("api/"+apiVersion+"/check", healthCheckHandler.HealthCheck, middlewares.JwtMiddleware)
	serverInstance.POST("api/"+apiVersion+"/register", userHandler.Register)
	serverInstance.POST("api/"+apiVersion+"/login", userHandler.Login)
	serverInstance.POST("api/"+apiVersion+"/token/refresh", refreshTokenHandler.Refresh)
}
