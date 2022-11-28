package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const CREATE_TABLE_PRODUTOS = "create table if not exists produtos(id serial primary key, nome varchar, descricao varchar, quantidade integer, preco decimal)"

func Connect() *sql.DB {
	// user, dbname, password, host, ssl(true/false)
	con := "user=administrator dbname=produtos password=pass123456 host=localhost sslmode=disable"
	db, error := sql.Open("postgres", con)

	if error != nil {
		panic(error.Error())
	} else {
		return db
	}
}

func CreateAppTable() {
	db := Connect()
	defer db.Close()
	_, err := db.Exec(CREATE_TABLE_PRODUTOS)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("Tabela OK")
	}
}
