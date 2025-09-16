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

func AddUsers(res http.ResponseWriter, req *http.Request) {
	var newUser models.User

	json.NewDecoder(req.Body).Decode(&newUser)

	newData := db.DB.Create(&newUser)
	err := newData.Error

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(res).Encode(&newUser)
}

func SearchUser(res http.ResponseWriter, req *http.Request) {
	var searchedUser models.User
	param := mux.Vars(req)

	db.DB.Where("name = ?", param["name"]).Find(&searchedUser)

	if searchedUser.Name == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(res).Encode(&searchedUser)
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
