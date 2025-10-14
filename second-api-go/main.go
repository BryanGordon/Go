package main

import (
	"log"
	"net/http"
	"second-api-go/db"
	"second-api-go/routes"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

var allowedOri = []string{"*"}
var allowedMeth = []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"}
var allowedHeaders = []string{"Content-Type", "Authorization"}

func main() {
	/* Conexion con local db mysql */
	// db.DbConnection()
	/* Conexion con supabase */
	godotenv.Load() // Siempre llamar en la funcion principal

	db.SupaConnect()

	data := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":3000",
		handlers.CORS(handlers.AllowedOrigins(allowedOri), handlers.AllowedMethods(allowedMeth), handlers.AllowedHeaders(allowedHeaders))(data)))

}
