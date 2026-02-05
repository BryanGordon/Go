package routes

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"os"
	"ticket-api/models"
)

var filename string = "tickets.json"

func generateRandomNumber() int {
	number := rand.Int()
	return number
}

func GenerateTicketConcert(res http.ResponseWriter, req *http.Request) {
	var ticketList []models.Ticket
	// params := mux.Vars(req)

	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("File could not read"))
			return
		}

		_ = json.Unmarshal(data, &ticketList)
		newTicket := models.Ticket{Id: "1", Type: "Concert", Number: generateRandomNumber()}
		ticketList = append(ticketList, newTicket)
		dataJson, err := json.MarshalIndent(ticketList, "", "")

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Data not saved"))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write(dataJson)
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("File not found"))
}

func GenerateTicketMovie(res http.ResponseWriter, req *http.Request) {
	var ticketList []models.Ticket

	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("File could not read"))
			return
		}

		_ = json.Unmarshal(data, &ticketList)
		newTicket := models.Ticket{Id: "2", Type: "Movie", Number: generateRandomNumber()}
		ticketList = append(ticketList, newTicket)
		dataJson, err := json.MarshalIndent(ticketList, "", "")

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Could not write the file"))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write(dataJson)
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("File not found"))
}

func GenerateTicketTrain(res http.ResponseWriter, req *http.Request) {

}

func DeleteTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro
}

func InvalidateTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro

}
