package routes

import (
	"auth-api/connections"
	"auth-api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func UserList(res http.ResponseWriter, req *http.Request) {
	data, _, err := connections.Client.From("readers").Select("*", "", false).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("unable get data"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func SearchUser(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)

	data, _, err := connections.Client.From("readers").Select("*", "", false).Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		res.Write([]byte("User not found"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func AddUser(res http.ResponseWriter, req *http.Request) {
	var newUser models.Users

	json.NewDecoder(req.Body).Decode(&newUser)

	data, _, err := connections.Client.From("readers").Insert(newUser, false, "", "", "").Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't create new data"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateUserName(res http.ResponseWriter, req *http.Request) {
	var newName models.Users
	param := mux.Vars(req)

	json.NewDecoder(req.Body).Decode(&newName)
	updatedName := map[string]any{
		"name": newName.Name,
	}

	data, _, err := connections.Client.From("readers").Update(updatedName, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't update name"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateUserMail(res http.ResponseWriter, req *http.Request) {
	var newMail models.Users
	param := mux.Vars(req)

	json.NewDecoder(req.Body).Decode(&newMail)
	updatedMail := map[string]any{
		"email": newMail.Email,
	}

	data, _, err := connections.Client.From("readers").Update(updatedMail, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't update email"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateUserNick(res http.ResponseWriter, req *http.Request) {
	var newNick models.Users
	param := mux.Vars(req)

	json.NewDecoder(req.Body).Decode(&newNick)
	updatedNick := map[string]any{
		"nickname": newNick.Nickname,
	}

	data, _, err := connections.Client.From("readers").Update(updatedNick, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't update nickname"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func RemoveUser(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)

	data, _, err := connections.Client.From("readers").Delete("", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't delete user"))
		return
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write(data)
}
