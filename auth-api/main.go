package main

import (
	"auth-api/connections"
	"auth-api/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	connections.SupaConnec()

	routes := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":3000",
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:4321"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		)(routes)))

}
