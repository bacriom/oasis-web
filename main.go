package main

import (
	"html/template"
	"net/http"
	"os"
)

func renderTemplate(w http.ResponseWriter, page string) {
	files := []string{
		"templates/layout.html",
		"templates/" + page,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// 📁 Archivos estáticos
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handler dinámico
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// HOME
		if r.URL.Path == "/" {
			renderTemplate(w, "home.html")
			return
		}
		// Obtener nombre de página
		page := r.URL.Path[1:] + ".html"

		// Verificar si existe el archivo
		if _, err := os.Stat("templates/" + page); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		renderTemplate(w, page)
	})

	// 🔥 Puerto dinámico (CLAVE para deploy)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Servidor corriendo en puerto " + port)

	http.ListenAndServe(":"+port, nil)
}
