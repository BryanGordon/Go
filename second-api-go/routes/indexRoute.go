package routes

import "github.com/gorilla/mux"

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	/* Routes with mysql */
	routes.HandleFunc("/", GetProducts).Methods("GET")
	routes.HandleFunc("/add", AddProduct).Methods("POST")
	routes.HandleFunc("/{id}", UpdateProduct).Methods("PATCH")
	routes.HandleFunc("/search/{name}", SearchProduct).Methods("GET")
	routes.HandleFunc("/{name}", DeleteProduct).Methods("DELETE")

	routes.HandleFunc("/users", GetUsers).Methods("GET")
	routes.HandleFunc("/users/add", AddUsers).Methods("POST")
	routes.HandleFunc("/users/search/{name}", SearchUser).Methods("GET")
	routes.HandleFunc("/users/{id}", UpdateRolUser).Methods("PATCH")
	routes.HandleFunc("/users/update/{id}", UpdateNameUser).Methods("PATCH")
	routes.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	/* Routes with supabase */
	routes.HandleFunc("/supa/users", GetUsersSupa).Methods("GET")
	routes.HandleFunc("/supa/products", GetDataSupa).Methods("GET")

	routes.HandleFunc("/supa/users/search/{name}", SearchUserSupa).Methods("GET")

	return routes
}
