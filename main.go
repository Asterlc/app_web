package main

import (
	"app_web/db/models"
	"fmt"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("Iniciando na porta 9080")
	http.HandleFunc("/", index)
	error := http.ListenAndServe(":9080", nil)
	if error != nil {
		panic(error.Error())
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	produtos := models.GetProdutos()

	temp.ExecuteTemplate(writer, "Index", produtos)
}
