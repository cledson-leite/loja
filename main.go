package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome, Descricao string
	Preco float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000",nil)
}

func index (writer http.ResponseWriter, request *http.Request) {
	produtos := []Produto{
		{Nome: "Tênis", Descricao: "Confortável", Preco: 89, Quantidade: 5},
		{Nome: "Camiseta", Descricao: "Azul, bem Bonita", Preco: 39, Quantidade: 10},
		{Nome: "Fone", Descricao: "Muito bom", Preco: 59, Quantidade: 3},
		{Nome: "Produto novo", Descricao: "Muito Legal", Preco: 1.99, Quantidade: 1},
	}
	templates.ExecuteTemplate(writer, "Index", produtos)
}