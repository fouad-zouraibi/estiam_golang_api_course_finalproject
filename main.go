package main

/*
This project was created by :
    Fouad ZOURAIBI,
    Taha GARGOURI,
    Haider JUOINI.
*/
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
	// userHandler := handlers.NewUserHandler(userService)
	// server.GET("/users/:id", userHandler.Get)

	// Register a new endpoint for POST user
	userHandler := handlers.NewUserHandler(userService)
	server.POST("/users", userHandler.Create) // Register the handler function for POST /users

	if err := server.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}