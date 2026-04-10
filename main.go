package main

import (
	"html/template"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func render(w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// 📁 Archivos estáticos
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 🌐 Ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w)
	})

	// 🔥 Puerto dinámico (CLAVE para deploy)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Servidor corriendo en puerto " + port)

	http.ListenAndServe(":"+port, nil)
}
