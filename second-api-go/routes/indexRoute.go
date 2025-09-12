package routes

import "github.com/gorilla/mux"

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", GetProducts).Methods("GET")
	routes.HandleFunc("/add", AddProduct).Methods("POST")
	routes.HandleFunc("/{id}", UpdateProduct).Methods("PATCH")
	routes.HandleFunc("/search/{name}", SearchProduct).Methods("GET")
	routes.HandleFunc("/{name}", DeleteProduct).Methods("DELETE")

	return routes
}
