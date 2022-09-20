package routes

import (
	"net/http"
	"web/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
