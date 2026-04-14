package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"oasis-web/backend/data"
)

func ComiteDetalleHandler(w http.ResponseWriter, r *http.Request) {

	nombre := strings.TrimPrefix(r.URL.Path, "/comite/")

	// 🔥 CASO: /comite/
	if nombre == "" {
		tmpl, err := template.ParseFiles(
			"templates/layout.html",
			"templates/comite.html",
		)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		tmpl.ExecuteTemplate(w, "layout", nil)
		return
	}

	// 🔥 CASO: /comite/presidente
	miembro, ok := data.Comite[nombre]
	if !ok {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/comite/detalles.html",
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", miembro)
}
