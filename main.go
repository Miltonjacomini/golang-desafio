package main

import (
	"miltonjacomini/golang-desafio/routes"

	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Bem vinda(o) ao sistema de busca de filmes!")
	fmt.Println("Iniciando servidor...")
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
