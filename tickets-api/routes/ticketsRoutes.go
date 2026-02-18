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

		_ = json.Unmarshal(data, &allTicket)
		newTicket := models.Ticket{Id: "1", Type: "Concert", Number: generateRandomNumber()}
		allTicket.Concert = append(allTicket.Concert, newTicket)
		dataJson, err := json.MarshalIndent(allTicket, "", " ")

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

		ticketData, _ := json.Marshal(newTicket)

		res.WriteHeader(http.StatusOK)
		res.Write([]byte(ticketData))
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("File not found"))
}

func GenerateTicketMovie(res http.ResponseWriter, req *http.Request) {
	var ticketsAvailable models.TicketAvailable

	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("File could not read"))
			return
		}

		_ = json.Unmarshal(data, &ticketsAvailable)
		newTicket := models.Ticket{Id: "2", Type: "Movie", Number: generateRandomNumber()}
		ticketsAvailable.Movie = append(ticketsAvailable.Movie, newTicket)
		dataJson, err := json.MarshalIndent(ticketsAvailable, " ", "")

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
	var ticketList models.TicketAvailable

	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Could not read the file."))
			return
		}

		_ = json.Unmarshal(data, &ticketList)

		newTicket := models.Ticket{Id: "3", Type: "Train", Number: generateRandomNumber()}
		ticketList.Train = append(ticketList.Train, newTicket)
		newDataJson, err := json.MarshalIndent(ticketList, " ", " ")

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("File error."))
			return
		}

		err = os.WriteFile(filename, newDataJson, 0644)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Could not update file"))
			return
		}

		newTicketData, _ := json.Marshal(newTicket)

		res.WriteHeader(http.StatusOK)
		res.Write(newTicketData)
		return
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("File not found"))
}

func DeleteTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro
}

func InvalidateTicket(res http.ResponseWriter, req *http.Request) {
	// pedir id string como parametro

}
