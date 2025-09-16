package routes

import (
	"encoding/json"
	"net/http"
	"second-api-go/db"
	"second-api-go/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	var updateRol models.User
	var userID models.User
	param := mux.Vars(req)

	db.DB.Where("id = ?", param["id"]).First(&userID)

	if userID.Id == uuid.Nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("user not found"))
		return
	}

	json.NewDecoder(req.Body).Decode(&updateRol)

	updatedResult := db.DB.Model(&updateRol).Where("id = ?", userID.Id).Update("rol", updateRol.Rol)
	err := updatedResult.Error

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Could not update field"))
		return
	}

	json.NewEncoder(res).Encode(&updateRol.Rol)
}

func UpdateNameUser(res http.ResponseWriter, req *http.Request) {
	var updatedName models.User
	var userId models.User
	param := mux.Vars(req)

	db.DB.Where("id = ?", param["id"]).First(&userId)

	if userId.Id == uuid.Nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("user not found"))
		return
	}

	json.NewDecoder(req.Body).Decode(&updatedName)

	updatedResult := db.DB.Model(&updatedName).Where("id = ?", userId.Id).Update("name", updatedName.Name)
	err := updatedResult.Error

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Could not update field"))
		return
	}

	json.NewEncoder(res).Encode(&updatedName.Name)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("deleted user"))
}
