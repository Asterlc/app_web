package main

import (
	"app_web/routes"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando na porta 9080")
	routes.Routes()

	error := http.ListenAndServe(":9080", nil)
	if error != nil {
		panic(error.Error())
	}
}
