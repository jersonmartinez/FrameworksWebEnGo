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
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "¡Hola, mundo!")
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "¡Hola, %s!", nombre)
	})

	r.POST("/usuarios", func(c *gin.Context) {
		var nuevoUsuario Usuario

		if err := c.BindJSON(&nuevoUsuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar el JSON"})
			return
		}

		if nuevoUsuario.Nombre == "" || nuevoUsuario.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre y correo electrónico son campos requeridos"})
		}

		usuarios = append(usuarios, nuevoUsuario)

		c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario registrado", "datos": usuarios})
	})
}
