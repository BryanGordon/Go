package routes

import (
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()

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
	return routes
}
