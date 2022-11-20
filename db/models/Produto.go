package models

import (
	"app_web/db"
	"database/sql"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Quantidade      int
	Preco           float64
}

func Get() []Produto {
	db := db.Connect()
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func Create(produto Produto) (bool, sql.Result) {
	db := db.Connect()

	prepare, err := db.Prepare("insert into produtos(nome,descricao,quantidade,preco) values($1,$2,$3,$4) ")
	if err != nil {
		panic(err.Error())
	}
	result, sqlError := prepare.Exec(produto.Nome, produto.Descricao, produto.Quantidade, produto.Preco)

	if sqlError != nil {
		panic(sqlError.Error())
	}
	defer db.Close()
	return true, result
}

// func Update(id int, produto Produto) {
// 	db := db.Connect()
// }

func Delete(id int) (bool, sql.Result) {
	db := db.Connect()
	prepare, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec()
	result, sqlError := prepare.Exec(id)
	if sqlError != nil {
		panic(sqlError.Error())
	}
	defer db.Close()
	return true, result
}
