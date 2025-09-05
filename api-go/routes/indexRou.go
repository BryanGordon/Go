package routes

import (
	"api-go/controllers"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	api := routes.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", controllers.GetHome).Methods("GET")

	return routes
}
