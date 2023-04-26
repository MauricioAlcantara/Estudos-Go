package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var view = template.Must(template.ParseGlob("src/views/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Notebook", Descricao: "Esta meio velho mais ainda e muito bom", Preco: 1499, Quantidade: 1},
		{Nome: "Mouse", Descricao: "Um mouse muito bom e novo", Preco: 250, Quantidade: 5},
		{Nome: "Jaqueta", Descricao: "Jaqueta de couro legítimo", Preco: 500, Quantidade: 2},
	}

	view.ExecuteTemplate(w, "Index", produtos)
}
