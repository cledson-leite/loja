package controller

import (
	"log"
	"loja/models"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index (writer http.ResponseWriter, request *http.Request) {
	produtos := models.GetAll()
	templates.ExecuteTemplate(writer, "Index", produtos)
}
func Novo (writer http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(writer, "Novo", nil)
}
func Insert (writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		nome := request.FormValue("nome")
		descricao := request.FormValue("descricao")

		converteFloat, err := strconv.ParseFloat(request.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro na converção: ", err)
		}

		converteInt, err := strconv.Atoi(request.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro na converção: ", err)
		}

		models.PostProduto(nome, descricao, converteFloat, converteInt)
	}
	http.Redirect(writer, request, "/", 301)
}