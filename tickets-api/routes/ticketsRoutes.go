package routes

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"ticket-api/models"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
)

var filename string = "tickets.json"
var letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateRandomNumber() string {
	number := rand.IntN(99999999999)
	convertNumber := strconv.Itoa(number)
	return convertNumber
}

func generateRandomTicketNumber() string {
	number := rand.IntN(99999999999)
	letter := letters[rand.IntN(len(letters))]
	convertNumber := strconv.Itoa(number)

	result := convertNumber + " - " + string(letter)
	return result
}

func GenerateTicketConcert(res http.ResponseWriter, req *http.Request) {
	var allTicket models.TicketAvailable

	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("File could not read"))
			return
		}

		_ = json.Unmarshal(data, &allTicket)
		keyID, err := uuid.NewV6()

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Error creating new ticket."))
			return
		}

		newTicket := models.Ticket{Id: keyID, Type: "Concert", Number: generateRandomNumber()}
		allTicket.Concert = append(allTicket.Concert, newTicket)
		dataJson, err := json.MarshalIndent(allTicket, " ", " ")

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
		keyId, err := uuid.NewV6()

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Error creating new ticket."))
			return
		}

		newTicket := models.Ticket{Id: keyId, Type: "Movie", Number: generateRandomNumber()}
		ticketsAvailable.Movie = append(ticketsAvailable.Movie, newTicket)
		dataJson, err := json.MarshalIndent(ticketsAvailable, " ", " ")

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
		keyId, err := uuid.NewV6()

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Error creating new ticket."))
			return
		}

		newTicket := models.Ticket{Id: keyId, Type: "Train", Number: generateRandomTicketNumber()}
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

func ValidateTicket(res http.ResponseWriter, req *http.Request) bool {
	var params = mux.Vars(req)
	var tickets models.TicketAvailable

	_, err := os.Stat(filename)

	if err == nil {
		data, err := os.ReadFile(filename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Error reading data."))
			return false
		}

		_ = json.Unmarshal(data, &tickets)

		for _, value := range tickets.Concert {
			if value.Id.String() == params["id"] {
				// Cambiar estado de
				newJson, err := json.MarshalIndent(tickets, " ", " ")

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					res.Write([]byte("File error."))
					return false
				}

				err = os.WriteFile(filename, newJson, 0644)

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					res.Write([]byte("Error updating file."))
					return false
				}

				res.WriteHeader(http.StatusOK)
				res.Write([]byte("Ticket validated."))
				return true
			}
		}

		for _, value := range tickets.Movie {
			if value.Id.String() == params["id"] {
				// cambiar valor del validacion
				newJson, err := json.MarshalIndent(tickets, " ", " ")

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					res.Write([]byte("File error."))
					return false
				}

				err = os.WriteFile(filename, newJson, 0644)

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					res.Write([]byte("Error updating file."))
					return false
				}

				res.WriteHeader(http.StatusOK)
				res.Write([]byte("Ticket validated"))
				return true
			}
		}

		for _, value := range tickets.Train {
			if value.Id.String() == params["id"] {
				// Cambiar valor
				newJson, err := json.MarshalIndent(tickets, " ", " ")

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					res.Write([]byte("File error."))
					return false
				}

				err = os.WriteFile(filename, newJson, 0644)

				if err != nil {
					res.WriteHeader(http.StatusBadRequest)
					res.Write([]byte("Error updating file."))
					return false
				}

				res.WriteHeader(http.StatusOK)
				res.Write([]byte("Ticket Validated"))
				return true
			}
		}
	}

	res.WriteHeader(http.StatusBadRequest)
	res.Write([]byte("Couldn't read data file."))
	return false
}
