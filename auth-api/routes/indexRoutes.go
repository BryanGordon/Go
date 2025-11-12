package routes

import (
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()
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
	routes.HandleFunc("/admin/users", UserList).Methods("GET")
	routes.HandleFunc("/admin/users", AddUser).Methods("POST")
	routes.HandleFunc("/admin/users/search/{id}", SearchUser).Methods("GET")
	routes.HandleFunc("/admin/users/update-name/{id}", UpdateUserName).Methods("PATCH")
	routes.HandleFunc("/admin/users/{id}", RemoveUser).Methods("DELETE")

	/* Todo */
	// - Envolver la ruta del admin con auth
	// - No envolver la parte del user (Envolver despues)
	// - Registar los usuarios en supabase para el login
	// - Agregar la columna de rol a la tabla de readers
	// - Crear el login en el frontend
	// - Solucionar errores
	return routes
}
