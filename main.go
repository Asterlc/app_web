package main

import (
	"app_web/db"
	"app_web/routes"
	"fmt"
	"log"
	"net/http"
)

// CRIAR TABELA: create table produtos(id serial primary keynome varchar, descricao varchar, quantidade integer, preco decimal)
func main() {
	db := db.Connect()
	sqlRes, err := db.Exec(fmt.Sprintf("create table if not exists produtos(id serial primary key, nome varchar, descricao varchar, quantidade integer, preco decimal)"))
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println(sqlRes)
	}
	defer db.Close()
	fmt.Println("http://localhost:9080")
	routes.Routes()

	error := http.ListenAndServe(":9080", nil)
	if error != nil {
		panic(error.Error())
	}
}
