package controllers

import (
	"encoding/json"
	"net/http"
)

type PersonData struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int8   `json:"edad"`
}

type InitMessage struct {
	Message string `json:"mensaje"`
}

func GetHome(res http.ResponseWriter, req *http.Request) {
	mensaje := InitMessage{
		Message: "Funciona con struct",
	}

	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(mensaje)
}
