package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/phramos07/finalproject/handlers"
)

func main() {
	server := echo.New()

	healthHandler := handlers.NewHealthHandler()
	server.GET("/live", healthHandler.IsAlive)

	if err := server.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
