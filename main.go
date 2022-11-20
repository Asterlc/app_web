package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func cnPostgres() *sql.DB {
	// user, dbname, password, host, ssl(true/false)
	con := "user=administrator dbname=produtos password=pass123456 host=localhost sslmode=disable"
	db, error := sql.Open("postgres", con)

	if error != nil {
		panic(error.Error())
	} else {
		return db
	}
}

type Produto struct {
	id              int
	Nome, Descricao string
	Quantidade      int
	Preco           float64
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("Iniciando na porta 9080")
	cnPostgres()
	http.HandleFunc("/", index)
	error := http.ListenAndServe(":9080", nil)
	if error != nil {
		panic(error.Error())
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	db := cnPostgres()
	row, error := db.Query("select * from produtos")
	if error != nil {
		panic(error.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for row.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := row.Scan(&id, &nome, &descricao, &quantidade, &preco)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
		fmt.Println(produtos)

	}

	temp.ExecuteTemplate(writer, "Index", produtos)
	defer db.Close()

}
