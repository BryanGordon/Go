package routes

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.HandleFunc("/concert-tickets", GenerateTicketConcert).Methods("POST")
	routes.HandleFunc("/movies-tickets", GenerateTicketMovie).Methods("POST")
	routes.HandleFunc("/train-tickets", GenerateTicketTrain).Methods("POST")
	routes.HandleFunc("/get-all-tickets", GetAllTickets).Methods("GET")

	return routes
}
