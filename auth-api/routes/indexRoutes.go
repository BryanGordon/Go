package routes

import (
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/", UserList).Methods("GET")
	routes.HandleFunc("/", AddUser).Methods("POST")
	routes.HandleFunc("/search", SearchUser).Methods("GET")
	routes.HandleFunc("/update-name", UpdateUserName).Methods("PATCH")
	routes.HandleFunc("/update-email", UpdateUserMail).Methods("PATCH")
	routes.HandleFunc("/update-nick", UpdateUserNick).Methods("PATCH")
	routes.HandleFunc("/", RemoveUser).Methods("DELETE")

	routes.HandleFunc("/books", ListBooks).Methods("GET")
	routes.HandleFunc("/books", AddBooks).Methods("POST")
	routes.HandleFunc("/search/books", SearchBook).Methods("GET")
	routes.HandleFunc("/books/update-name", UpdateNameBooks).Methods("PATCH")
	routes.HandleFunc("/books/update-author", UpdateBookAuthor).Methods("PATCH")
	routes.HandleFunc("/books/update-", UpdateBookPages).Methods("PATCH")
	routes.HandleFunc("/books", RemoveBook).Methods("DELETE")
	return routes
}
