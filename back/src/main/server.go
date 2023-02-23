package main

import (
	"back/src/pkg/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"sync"
)

var lock = &sync.Mutex{}

var serverInstance *echo.Echo

func launchServer(wg *sync.WaitGroup) *echo.Echo {
	lock.Lock()
	defer lock.Unlock()
	defer wg.Done()

	if serverInstance != nil {
		panic("Server already launched")
	} else {
		fmt.Println("Server launching")
		serverInstance := echo.New()
		serverInstance.Use(middleware.Logger())
		serverInstance.Use(middleware.Recover())
		services.InitializeServices()
		initializeHandlers(serverInstance)
		serverInstance.Logger.Fatal(serverInstance.Start(server + ":" + port))
	}
	return serverInstance
}
