package routes

import (
	"api-go/controllers"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/", controllers.GetHome).Methods("GET")
	routes.HandleFunc("/movies", GetMoviesHandler).Methods("GET")
	routes.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	routes.HandleFunc("/movies", PostMovie).Methods("POST")
	routes.HandleFunc("/movies", DeleteMovie).Methods("DELETE")

	return routes
}
