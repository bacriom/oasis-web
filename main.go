package main

import (
	"html/template"
	"net/http"
	"os"

	"oasis-web/backend/handlers"
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

	// 📁 static
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 🔥 rutas específicas PRIMERO
	http.HandleFunc("/comite/", handlers.ComiteDetalleHandler)
	http.HandleFunc("/transparencia/", handlers.TransparenciaHandler)
	http.HandleFunc("/solicitudes/", handlers.SolicitudesHandler)

	// 🔥 handler general al final
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" {
			renderTemplate(w, "home.html")
			return
		}

		page := r.URL.Path[1:] + ".html"

		if _, err := os.Stat("templates/" + page); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		renderTemplate(w, page)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Servidor corriendo en puerto " + port)

	http.ListenAndServe(":"+port, nil)
}
