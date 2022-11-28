package models

import (
	"app_web/db"
	"log"
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

func Create(produto Produto) {
	db := db.Connect()

	prepare, err := db.Prepare("insert into produtos(nome,descricao,quantidade,preco) values($1,$2,$3,$4) ")
	if err != nil {
		panic(err.Error())
	}
	_, sqlError := prepare.Exec(produto.Nome, produto.Descricao, produto.Quantidade, produto.Preco)

	if sqlError != nil {
		log.Println(sqlError.Error())
	} else {
		log.Println("Product created:", produto.Nome)
	}
	defer db.Close()
}

func Recovery(id int) Produto {
	db := db.Connect()
	row, error := db.Query("select * from produtos where id=$1", id)
	if error != nil {
		log.Println("Error to edit a product", error.Error())
	}

	product := Produto{}

	for row.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := row.Scan(&id, &nome, &descricao, &quantidade, &preco)
		if err != nil {
			log.Println("Error row update", err.Error())
		}
		product.Id = id
		product.Nome = nome
		product.Descricao = descricao
		product.Quantidade = quantidade
		product.Preco = preco
	}
	defer db.Close()
	return product

}

func Edit(produto Produto) {
	db := db.Connect()
	row, error := db.Prepare("update produtos set nome=$1, descricao=$2, quantidade=$3, preco=$4 where id=$5")
	if error != nil {
		log.Println("Error in update", error.Error())
	}
	_, sqlError := row.Exec(produto.Nome, produto.Descricao, produto.Quantidade, produto.Preco, produto.Id)
	if sqlError != nil {
		log.Println("Error update product", sqlError.Error())
		return
	}
	log.Println("Product ID:", produto.Id, "have been updated")
	defer db.Close()
}

func Delete(id int) {
	db := db.Connect()
	prepare, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	prepare.Exec()
	result, sqlError := prepare.Exec(id)
	if sqlError != nil {
		log.Println(sqlError.Error())
	}
	log.Println("Deleted", result)
	defer db.Close()
}
