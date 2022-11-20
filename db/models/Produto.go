package models

import (
	"app_web/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Quantidade      int
	Preco           float64
}

func GetProdutos() []Produto {
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

		p.Nome = nome
		p.Descricao = descricao
		p.Quantidade = quantidade
		p.Preco = preco

		produtos = append(produtos, p)
	}
	db.Close()
	return produtos
}
