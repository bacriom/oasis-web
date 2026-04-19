package data

import (
	"time"
)

type Documento struct {
	ID        string
	Titulo    string
	URL       string
	Categoria string
	Fecha     time.Time
}

type CategoriaDocs struct {
	Nombre string
	Docs   []Documento
}
