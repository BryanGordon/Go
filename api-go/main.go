package main

import (
	"api-go/data"
	"api-go/routes"
	"log"
	"net/http"
)

func main() {

	data.ConnectionDB()
	result := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":3000", result))

}
