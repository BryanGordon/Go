package routes

import "github.com/gorilla/mux"

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", GetProducts).Methods("GET")
	routes.HandleFunc("/add", AddProduct).Methods("POST")
	routes.HandleFunc("/{id}", UpdateProduct).Methods("GET")
	routes.HandleFunc("/search", SearchProduct).Methods("POST")
	routes.HandleFunc("/{id}", DeleteProduct).Methods("DELETE")

	return routes
}
