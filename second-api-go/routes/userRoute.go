package routes

import (
	"encoding/json"
	"net/http"
	"second-api-go/db"
	"second-api-go/models"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	var usersList []models.User

	result := db.DB.Find(&usersList)
	err := result.Error

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("No users found"))
	}

	json.NewEncoder(res).Encode(&usersList)
}

func SearchUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Get user info"))
}

func UpdateRolUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Updated user rol"))
}

func UpdateNameUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Updated user name"))
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("deleted user"))
}
