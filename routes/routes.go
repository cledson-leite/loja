package routes

import (
	"loja/controller"
	"net/http"
)

func Router()  {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/novo", controller.Novo)
	http.HandleFunc("/insert", controller.Insert)
	
}