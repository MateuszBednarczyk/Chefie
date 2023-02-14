package main

import (
	"back/db"
	"back/handlers"
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
	dbName     = "ex"
)

func main() {
	e := echo.New()
	initializeHandlers(e)
	dbConfig := createDbConfig()
	db.Init(&dbConfig)
	e.Logger.Fatal(e.Start(server + ":" + port))
}

func initializeHandlers(e *echo.Echo) {
	e.GET("api/"+apiVersion+"/check", handlers.HealthCheck)
}

func createDbConfig() db.Config {
	return db.Config{DbUsername: dbUsername, DbPassword: dbPassword, DbPort: dbPort, DbHost: dbHost, DbName: dbName}
}
