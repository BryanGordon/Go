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
	var allTicket models.TicketAvailable
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
		allTicket.Concert = append(allTicket.Concert, newTicket)
		// ticketList = append(ticketList, newTicket)
		dataJson, err := json.MarshalIndent(ticketList, "", "")

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Data not saved"))
			return
		}

		err = os.WriteFile(filename, dataJson, 0644)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Could not update data file."))
			return
		}

		ticketData, err := json.Marshal(newTicket)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Error encoding new ticket value."))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(ticketData))
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

		err = os.WriteFile(filename, dataJson, 0644)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Could not update data file"))
			return
		}

		ticketData, _ := json.Marshal(newTicket)
		/*
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				res.Write([]byte("Could not encode new ticket data."))
			}
		*/
		res.WriteHeader(http.StatusOK)
		res.Write(ticketData)
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("File not found"))
}

func GetAllTickets(res http.ResponseWriter, req *http.Request) {
	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Error reading data"))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Write(data)
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("Could not found the file."))
}

func GenerateTicketTrain(res http.ResponseWriter, req *http.Request) {

}

func DeleteTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro
}

func InvalidateTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro

}
