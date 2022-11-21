package main

import (
	"app_web/routes"
	"fmt"
	"net/http"
)

// CRIAR TABELA: create table produtos(nome varchar, descricao varchar, quantidade integer, preco decimal)
func main() {
	fmt.Println("http://localhost:9080")
	routes.Routes()

	error := http.ListenAndServe(":9080", nil)
	if error != nil {
		panic(error.Error())
	}
}
