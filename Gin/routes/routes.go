package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Mi aplicación",
			"Heading": "¡Hola, mundo!",
			"Message": "Bienvenido a mi aplicación web con Gin y plantillas HTML.",
		})
	})

	r.Static("/static", "./static")
}
