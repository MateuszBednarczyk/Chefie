package main

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"back/src/pkg/services"
)

var lock = &sync.Mutex{}

func launchServer(wg *sync.WaitGroup, serverInstance *echo.Echo, ch chan string) {
	lock.Lock()
	defer lock.Unlock()
	defer wg.Done()

	if serverInstance != nil {
		ch <- "Server is already running"
	}
	fmt.Println("Server launching")
	serverInstance = echo.New()
	serverInstance.Use(middleware.Logger())
	serverInstance.Use(middleware.Recover())
	services.InitializeServices()
	initializeHandlers(serverInstance)
	serverInstance.Logger.Fatal(serverInstance.Start(server + ":" + port))
}
