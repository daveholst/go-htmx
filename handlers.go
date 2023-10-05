package main

import (
	"go-htmx/models"
	"html/template"
	"net/http"
)

func getPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]models.Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	// Returns the full template
	tmpl.Execute(w, films)

}

func postFilm(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("index.html"))

	// Returns the named block 'film-list-element' only
	tmpl.ExecuteTemplate(w, "film-list-element", models.Film{Title: title, Director: director})
}
