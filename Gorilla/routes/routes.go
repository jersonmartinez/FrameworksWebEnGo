package routes

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	renderer := Renderer()

	r.HandleFunc("/", renderTemplate(renderer, "index.html"))

	r.HandleFunc("/{page}", dynamicPageHandler(renderer)).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
}

func Renderer() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}

func renderTemplate(renderer *template.Template, templaName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := renderer.ExecuteTemplate(w, templaName, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func dynamicPageHandler(renderer *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		page := vars["page"]

		if !strings.HasSuffix(page, ".html") {
			page += ".html"
		}

		if _, err := os.Stat("templates/" + page); err == nil {
			err := renderer.ExecuteTemplate(w, page, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		}
	}
}
