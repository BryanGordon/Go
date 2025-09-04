package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Funciona"))
}

func main() {
	routes := mux.NewRouter()

	routes.HandleFunc("/", home)

	http.ListenAndServe(":8080", routes)
}
