package main

import (
	"log"
	"net/http"
	"ticket-api/routes"
)

func main() {
	data := routes.Routes()

	log.Fatal(http.ListenAndServe(":3000", data))
}
