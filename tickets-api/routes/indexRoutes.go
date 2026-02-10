package routes

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/concert-tickets", GenerateTicketConcert).Methods("POST")
	routes.HandleFunc("/movies-tickets", GenerateTicketMovie).Methods("POST")

	return routes
}
