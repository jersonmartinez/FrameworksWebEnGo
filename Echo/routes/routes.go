package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		//return c.String(http.StatusOK, "¡Hola, mundo!")

		return c.JSON(http.StatusOK, map[string]string{
			"message": "¡Hola desde Echo!",
		})
	})

	e.GET("/saludo", func(c echo.Context) error {
		return c.String(http.StatusOK, "¡Saludos desde Echo!")
	})

	e.GET("/saludo/:nombre", func(c echo.Context) error {
		nombre := c.Param("nombre")
		return c.String(http.StatusOK, "¡Hola, "+nombre+"!")
	})
}
