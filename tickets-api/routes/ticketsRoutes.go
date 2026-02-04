package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"ticket-api/models"

	"github.com/gorilla/mux"
)

var filename string = "tickets.json"

func GenerateTicketConcert(res http.ResponseWriter, req *http.Request) {
	var ticketList []models.Ticket
	params := mux.Vars(req)

	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("File could not read"))
			return
		}

		_ = json.Unmarshal(data, &ticketList)
		newTicket := models.Ticket{Id: params["id"], Type: "Concert", Number: 1231}
		ticketList = append(ticketList, newTicket)
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("File not found"))
}

func GenerateTicketMovie(res http.ResponseWriter, req *http.Request) {

}

func GenerateTicketTrain(res http.ResponseWriter, req *http.Request) {

}

func DeleteTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro
}

func InvalidateTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro

}
