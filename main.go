package main

import (
	"html/template"
	"net/http"
	"os"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

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

	// RUTAS
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ⚠️ Evita que "/" capture todo
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		renderTemplate(w, "home.html")
	})

	http.HandleFunc("/privacidad", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "privacidad.html")
	})

	http.HandleFunc("/reglamento", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "reglamento.html")
	})

	http.HandleFunc("/seguridad", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "seguridad.html")
	})

	http.HandleFunc("/transparencia", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "transparencia.html")
	})

	http.HandleFunc("/comite", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "comite.html")
	})

	http.HandleFunc("/anuncios", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "anuncios.html")
	})

	http.HandleFunc("/solicitudes", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "solicitudes.html")
	})

	// 🔥 Puerto dinámico (CLAVE para deploy)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Servidor corriendo en puerto " + port)

	http.ListenAndServe(":"+port, nil)
}
