package main

import (
	"log"
	"net/http"
	"second-api-go/db"
	"second-api-go/routes"
)

func main() {
	db.DbConnection()
	data := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":3000", data))
}
