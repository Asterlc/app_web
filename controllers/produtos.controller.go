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
	produtos := models.Get()

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
			log.Println("Error to parse product price", floatErr.Error())
			log.Println("Error to parse product amount", intErr.Error())
		} else {
			produto := models.Produto{
				Nome:       nome,
				Descricao:  descricao,
				Quantidade: quantidade,
				Preco:      preco,
			}
			models.Create(produto)
			http.Redirect(writer, request, "/", 301)
		}
	}
}

// func EditProduct(writer http.ResponseWriter, request *http.Request) {
// 	id, error := strconv.Atoi(request.URL.Query().Get("id"))
// }

func DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	id, error := strconv.Atoi(request.URL.Query().Get("id"))
	if error != nil {
		log.Println("Error on delete", error.Error())
	} else {
		models.Delete(id)
		http.Redirect(writer, request, "/", 301)

	}
}
