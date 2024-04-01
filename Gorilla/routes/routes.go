package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "¡Bienvenido al servidor Gorilla!")
	})

	r.HandleFunc("/saludo/{nombre}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nombre := vars["nombre"]
		w.Write([]byte("¡Hola, " + nombre + "!"))
	})
}
