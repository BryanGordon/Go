package main

import (
	"context"
	"log"
	"net/http"
	"ticket-api/manage"
	"ticket-api/routes"

	"github.com/gorilla/handlers"
)

var allowedAccess = []string{"http://localhost:5173"}
var allowedMethods = []string{"GET", "POST", "OPTIONS"}
var allowedHeaders = []string{"Content-Type"}

func main() {
	data := routes.Routes()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	manage.CleanDataFile(ctx)

	log.Fatal(http.ListenAndServe(":3000",
		handlers.CORS(handlers.AllowedOrigins(allowedAccess), handlers.AllowedMethods(allowedMethods), handlers.AllowedHeaders(allowedHeaders))(data)))
}
