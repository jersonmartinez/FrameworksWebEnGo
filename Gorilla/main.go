package main

import (
	"Gorilla/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.SetupRoutes(r)

	r.Use(LogginMiddleware)

	log.Println("Servidor Gorilla en ejecución en http://127.0.0.1:8080")
	http.ListenAndServe(":8080", r)
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		latency := time.Since(start)

		log.Printf("%s %s - %v - %s", r.Method, r.URL.Path, latency, http.StatusText(http.StatusOK))
	})
}
