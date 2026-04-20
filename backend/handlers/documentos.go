package handlers

import (
	"html/template"
	"net/http"
	"oasis-web/backend/data"
	"sort"
	"strings"
	"time"
)

func TransparenciaHandler(w http.ResponseWriter, r *http.Request) {

	// 1. DOCUMENTOS
	documentos := []data.Documento{
		{
			ID:        "1j2WGNqaSd5UVaBplesHfyD1iWKi9smgQ",
			Titulo:    "Asamblea Marzo 2026",
			URL:       "https://docs.google.com/document/d/1j2WGNqaSd5UVaBplesHfyD1iWKi9smgQ/edit",
			Categoria: "Minutas",
			Fecha:     time.Date(2026, 3, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        "1l9zK8M3CfyPS65tJguTil8V9-UCAD0xt",
			Titulo:    "Firmas Asamblea Marzo 2026",
			URL:       "https://drive.google.com/file/d/1l9zK8M3CfyPS65tJguTil8V9-UCAD0xt/view?usp=sharing",
			Categoria: "Minutas",
			Fecha:     time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC),
		},
	}

	// 2. AGRUPAR POR CATEGORÍA
	categoriasMap := make(map[string][]data.Documento)

	for _, doc := range documentos {
		categoriasMap[doc.Categoria] = append(categoriasMap[doc.Categoria], doc)
	}

	// ordenar docs dentro de cada categoría
	for categoria := range categoriasMap {
		sort.Slice(categoriasMap[categoria], func(i, j int) bool {
			return categoriasMap[categoria][i].Fecha.After(
				categoriasMap[categoria][j].Fecha,
			)
		})
	}

	// convertir a slice (orden controlable)
	var categorias []data.CategoriaDocs

	orden := []string{"Minutas", "minutas"}
	for _, nombre := range orden {
		if docs, ok := categoriasMap[nombre]; ok {
			categorias = append(categorias, data.CategoriaDocs{
				Nombre: nombre,
				Docs:   docs,
			})
		}
	}

	// 3. TEMPLATE FUNCS (para usar lower en HTML)
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
	}

	// 4. CARGAR TEMPLATE
	tmpl := template.Must(
		template.New("").Funcs(funcMap).ParseFiles(
			"templates/layout.html",
			"templates/transparencia.html",
		),
	)

	// 5. DATA
	data := map[string]interface{}{
		"Categorias": categorias,
	}

	// 6. RENDER
	tmpl.ExecuteTemplate(w, "layout", data)
}
