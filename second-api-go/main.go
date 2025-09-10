package main

import (
	"log"
	"net/http"
	"second-api-go/routes"
)

func main() {
	data := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":3000", data))
}
