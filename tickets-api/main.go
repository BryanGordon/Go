package main

import (
	"log"
	"net/http"
	"ticket-api/routes"

	"github.com/gorilla/handlers"
)

var allowedAccess = []string{"http://localhost:5173"}
var allowedMethods = []string{"GET", "POST", "OPTIONS"}
var allowedHeaders = []string{"Content-Type"}

func main() {
	data := routes.Routes()

	log.Fatal(http.ListenAndServe(":3000",
		handlers.CORS(handlers.AllowedOrigins(allowedAccess), handlers.AllowedMethods(allowedMethods), handlers.AllowedHeaders(allowedHeaders))(data)))
}
