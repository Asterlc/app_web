package routes

import (
	"app_web/controllers"
	"net/http"
)

func Routes() {

	http.HandleFunc("/", controllers.Index)
}
