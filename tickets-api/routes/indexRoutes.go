package routes

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/concert-tickets", GenerateTicketConcert)
	routes.HandleFunc("/movies-tickets", GenerateTicketMovie)

	return routes
}
