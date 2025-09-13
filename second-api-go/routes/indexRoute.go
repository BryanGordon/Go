package routes

import "github.com/gorilla/mux"

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	routes.HandleFunc("/", GetProducts).Methods("GET")
	routes.HandleFunc("/add", AddProduct).Methods("POST")
	routes.HandleFunc("/{id}", UpdateProduct).Methods("PATCH")
	routes.HandleFunc("/search/{name}", SearchProduct).Methods("GET")
	routes.HandleFunc("/{name}", DeleteProduct).Methods("DELETE")

	routes.HandleFunc("/users", GetUsers).Methods("GET")
	routes.HandleFunc("/users/search/{name}", SearchUser).Methods("GET")
	routes.HandleFunc("/users/{id}", UpdateRolUser).Methods("PATCH")
	routes.HandleFunc("/users/{id}", UpdateNameUser).Methods("PATCH")
	routes.HandleFunc("/users/{name}", DeleteUser).Methods("DELETE")

	return routes
}
