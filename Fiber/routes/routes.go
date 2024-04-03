package routes

import "github.com/gofiber/fiber/v2"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SetupRoutes(app *fiber.App) {

	app.Get("/", handlerInicio)
	app.Get("/about", handlerAbout)
	app.Get("/saludo/:nombre", handlerSaludo)

	app.Get("/contact", func(c *fiber.Ctx) error {
		return c.SendString("Página de contacto para nuestra aplicación")
	})

	app.Post("/api/usuarios", func(c *fiber.Ctx) error {

		var usuario User
		if err := c.BodyParser(&usuario); err != nil {
			return err
		}

		return c.JSON(usuario)

	})
}

func handlerInicio(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{
		"Title":   "Mi aplicación",
		"Heading": "¡Hola, mundo!",
		"Message": "Bienvenido a mi aplicación web con Fiber y plantillas HTML.",
	})

	//return c.SendString("¡Bienvenido a mi aplicación con Fiber!")
}

func handlerAbout(c *fiber.Ctx) error {
	return c.SendString("Página de información sobre nuestra aplicación")
}

func handlerSaludo(c *fiber.Ctx) error {
	nombre := c.Params("nombre")

	return c.SendString("¡Hola, " + nombre + "!")
}
