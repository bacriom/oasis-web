package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func render(w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Archivos estáticos (css, js, imágenes)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w)
	})

	println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
