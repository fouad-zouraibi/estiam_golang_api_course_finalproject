package main

// this is a test to test commits with ssh

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/phramos07/finalproject/config"
	"github.com/phramos07/finalproject/handlers"
	"github.com/phramos07/finalproject/repos"
	"github.com/phramos07/finalproject/services"
)

func main() {
	server := echo.New()

	// load config
	config := config.Load()
	userRepo := repos.NewUserRepository(config.DbConn)
	userService := services.NewUserService(userRepo)

	healthHandler := handlers.NewHealthHandler()
	server.GET("/live", healthHandler.IsAlive)

	// REMOVE THAT ENDPOINT
	userHandler := handlers.NewUserHandler(userService)
	server.GET("/users/:id", userHandler.Get)

	// TODO: Register a new endpoint for POST user

	if err := server.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
