package controllers

import (
	"app_web/db/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	produtos := models.GetProdutos()

	temp.ExecuteTemplate(writer, "Index", produtos)
}

func NewProduct(writer http.ResponseWriter, request *http.Request) {
	temp.ExecuteTemplate(writer, "New", nil)
}

func AddProduct(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		nome := request.FormValue("nome")
		descricao := request.FormValue("descricao")
		quantidade, intErr := strconv.Atoi(request.FormValue("quantidade"))
		preco, floatErr := strconv.ParseFloat(request.FormValue("preco"), 64)
		fmt.Println(intErr)
		fmt.Println(floatErr)

		if intErr != nil && floatErr != nil {
			log.Println("Error preco", floatErr.Error())
			log.Println("Erro quantidade", intErr.Error())
		} else {
			produto := models.Produto{
				Nome:       nome,
				Descricao:  descricao,
				Quantidade: quantidade,
				Preco:      preco,
			}
			models.CreateProduct(produto)
			http.Redirect(writer, request, "/", 301)
		}
	}
}
