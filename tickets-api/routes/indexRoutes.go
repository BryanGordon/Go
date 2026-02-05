package routes

import "github.com/gorilla/mux"

func Routes() *mux.Router {
	routes := mux.NewRouter()

	routes.NewRoute().Handler()

	return routes
}
