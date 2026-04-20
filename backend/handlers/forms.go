package handlers

import (
	"html/template"
	"net/http"
	"oasis-web/backend/models"
)

func SolicitudesHandler(w http.ResponseWriter, r *http.Request) {

	forms := []models.Form{
		{
			ID:   "form1",
			Name: "Acceso a cámaras",
			URL:  "https://docs.google.com/forms/d/e/1FAIpQLSe2l6-wo8JZ0LrjutTkLEgHDfJCDoJkHe1_8VMFIyaWTXPw6g/viewform?embedded=true",
		},
		{
			ID:   "form2",
			Name: "Accesos / Tags",
			URL:  "https://docs.google.com/forms/d/e/1FAIpQLScoKJpAfkVTDbxVqZTorLgdyr4pJEJAyfycu4q1MwEuT0X5YA/viewform?embedded=true",
		},
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/solicitudes.html",
	))

	tmpl.ExecuteTemplate(w, "layout", forms)
}
