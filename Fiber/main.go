package main

import (
	"log"

	"Fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	app := fiber.New()
	engine := html.New("./views", ".html")

	app = fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/static", "./static")

	app.Use(logginMiddleware)

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func logginMiddleware(c *fiber.Ctx) error {
	log.Printf("Solicitud recibida: %s %s", c.Method(), c.Path())

	return c.Next()
}
