package controllers

import (
	"app_web/db/models"
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

		if intErr != nil || floatErr != nil {
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
			http.Redirect(writer, request, "/", http.StatusMovedPermanently)
		}
	}
}

func EditView(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(request.URL.Query().Get("id"))
	product := models.Recovery(id)
	temp.ExecuteTemplate(writer, "Edit", product)
}

func EditProduct(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		id, idParse := strconv.Atoi(request.FormValue("id"))
		nome := request.FormValue("nome")
		descricao := request.FormValue("descricao")
		quantidade, intErr := strconv.Atoi(request.FormValue("quantidade"))
		preco, floatErr := strconv.ParseFloat(request.FormValue("preco"), 64)

		if intErr != nil || floatErr != nil || idParse != nil {
			log.Println("Error to parse product price", floatErr.Error())
			log.Println("Error to parse product amount", intErr.Error())
		} else {
			produto := models.Produto{
				Id:         id,
				Nome:       nome,
				Descricao:  descricao,
				Quantidade: quantidade,
				Preco:      preco,
			}
			models.Edit(produto)
			http.Redirect(writer, request, "/", http.StatusMovedPermanently)
		}
	}
}

func DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	id, error := strconv.Atoi(request.URL.Query().Get("id"))
	if error != nil {
		log.Println("Error on delete", error.Error())
	} else {
		models.Delete(id)
		http.Redirect(writer, request, "/", http.StatusMovedPermanently)
	}
}
