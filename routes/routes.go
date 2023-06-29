package routes

import (
	"miltonjacomini/golang-desafio/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
