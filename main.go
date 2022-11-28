package main

import (
	"app_web/db"
	"app_web/routes"
	"fmt"
	"net/http"
)

// CRIAR TABELA: create table produtos(id serial primary keynome varchar, descricao varchar, quantidade integer, preco decimal)
func main() {
	db.CreateAppTable()
	fmt.Println("http://localhost:9080")
	routes.Routes()

	error := http.ListenAndServe(":9080", nil)
	if error != nil {
		panic(error.Error())
	}
}
