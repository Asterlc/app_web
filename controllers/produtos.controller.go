package controllers

import (
	"app_web/db/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	produtos := models.GetProdutos()

	temp.ExecuteTemplate(writer, "Index", produtos)
}
