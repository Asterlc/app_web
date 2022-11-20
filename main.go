package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fmt.Println("Iniciando na porta 8080")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	temp.ExecuteTemplate(writer, "Index", nil)
}
