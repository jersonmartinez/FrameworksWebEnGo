package main

import (
	"Gorilla/routes"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", map[string]interface{}{
			"Title":   "Mi aplicación",
			"Heading": "¡Hola, mundo!",
			"Message": "Bienvenido a mi aplicación con Gorilla y plantillas HTML.",
		})
	})

	routes.SetupRoutes(r)

	r.Use(LogginMiddleware)

	log.Println("Servidor Gorilla en ejecución en http://127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = "templates/" + tmpl
	t, err := template.ParseFiles(tmpl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		latency := time.Since(start)

		log.Printf("%s %s - %v - %s", r.Method, r.URL.Path, latency, http.StatusText(http.StatusOK))
	})
}
