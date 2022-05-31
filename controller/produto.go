package controller

import (
	"loja/models"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index (writer http.ResponseWriter, request *http.Request) {
	produtos := models.GetAll()
	templates.ExecuteTemplate(writer, "Index", produtos)
}