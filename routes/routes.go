package routes

import (
	"loja/controller"
	"net/http"
)

func Router()  {
		http.HandleFunc("/", controller.Index)

}