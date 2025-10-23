package main

import (
	"auth-api/connections"
	"auth-api/routes"
	"log"
	"net/http"
)

func main() {
	connections.SupaConnec()

	routes := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":3000", routes))

}
