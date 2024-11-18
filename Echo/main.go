package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"Echo/routes"
)

func main() {
	e := echo.New()

	//e.Use(middleware.Logger())

	routes.SetupRoutes(e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))

	e.Start(":8686")
}
