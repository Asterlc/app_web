package routes

import (
	"app_web/controllers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.NewProduct)
	http.HandleFunc("/insert", controllers.AddProduct)
}
