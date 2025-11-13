package routes

import (
	"auth-api/connections"
	"auth-api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
	var supaCli = connections.Client
	/* Routes for all users without authentication */
	/*
		routes.HandleFunc("/", UserList).Methods("GET")
		routes.HandleFunc("/", AddUser).Methods("POST")
		routes.HandleFunc("/search/{id}", SearchUser).Methods("GET")
		routes.HandleFunc("/update-name/{id}", UpdateUserName).Methods("PATCH")
		routes.HandleFunc("/update-email/{id}", UpdateUserMail).Methods("PATCH")
		routes.HandleFunc("/update-nick/{id}", UpdateUserNick).Methods("PATCH")
		routes.HandleFunc("/{id}", RemoveUser).Methods("DELETE")

		routes.HandleFunc("/books", ListBooks).Methods("GET")
		routes.HandleFunc("/books", AddBooks).Methods("POST")
		routes.HandleFunc("/search/books/{id}", SearchBook).Methods("GET")
		routes.HandleFunc("/books/update-name/{id}", UpdateNameBooks).Methods("PATCH")
		routes.HandleFunc("/books/update-author/{id}", UpdateBookAuthor).Methods("PATCH")
		routes.HandleFunc("/books/update-pages/{id}", UpdateBookPages).Methods("PATCH")
		routes.HandleFunc("/books/{id}", RemoveBook).Methods("DELETE")
	*/

	/*Routes with authentication*/
	routes.HandleFunc("/books", ListBooks).Methods("GET") // Accessible only for users

	// All these routes is available aonly for admin users
	routes.Handle("/admin/users", middlewares.AuthMiddleware(supaCli, middlewares.AdminAuthMiddleware(supaCli, http.HandlerFunc(UserList)))).Methods("GET")
	routes.Handle("/admin/users", middlewares.AuthMiddleware(supaCli, middlewares.AdminAuthMiddleware(supaCli, http.HandlerFunc(AddUser)))).Methods("POST")
	routes.Handle("/admin/users/search/{id}", middlewares.AuthMiddleware(supaCli, middlewares.AdminAuthMiddleware(supaCli, http.HandlerFunc(SearchUser)))).Methods("GET")
	routes.Handle("/admin/users/update-name/{id}", middlewares.AuthMiddleware(supaCli, middlewares.AdminAuthMiddleware(supaCli, http.HandlerFunc(UpdateUserName)))).Methods("PATCH")
	routes.Handle("/admin/users/{id}", middlewares.AuthMiddleware(supaCli, middlewares.AdminAuthMiddleware(supaCli, http.HandlerFunc(RemoveUser)))).Methods("DELETE")

	/* Todo */
	// - No envolver la parte del user (Envolver despues)
	// - Crear el login en el frontend
	// - Solucionar errores
	return routes
}
