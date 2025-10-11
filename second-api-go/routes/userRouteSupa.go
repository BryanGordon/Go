package routes

import (
	"encoding/json"
	"net/http"
	"second-api-go/db"
	"second-api-go/models"

	"github.com/gorilla/mux"
)

func GetUsersSupa(res http.ResponseWriter, req *http.Request) {
	data, _, err := db.SupaCli.From("users").Select("*", "", false).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Data not found"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func SearchUserSupa(res http.ResponseWriter, req *http.Request) {
	var param = mux.Vars(req)
	data, _, err := db.SupaCli.From("users").Select("*", "exact", false).Eq("name", param["name"]).Execute()

	/* Si el usuario no existe devuelve un array vacio, arreglar eso*/

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("User not found"))
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func CreateUserSupa(res http.ResponseWriter, req *http.Request) {
	var newUser models.User

	json.NewDecoder(req.Body).Decode(&newUser)

	data, _, err := db.SupaCli.From("users").Insert(newUser, false, "", "minimal", "").Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)

}

func UpdateUserRolSupa(res http.ResponseWriter, req *http.Request) {
	var param = mux.Vars(req)
	var newUserRol models.User

	json.NewDecoder(req.Body).Decode(&newUserRol)
	updatedRolUser := map[string]any{
		"rol": newUserRol.Rol,
	}

	data, _, err := db.SupaCli.From("users").Update(updatedRolUser, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Could not update the field"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateUserNameSupa(res http.ResponseWriter, req *http.Request) {
	var newUserName models.User
	var param = mux.Vars(req)

	json.NewDecoder(req.Body).Decode(&newUserName)
	updatedName := map[string]any{
		"name": newUserName.Name,
	}

	data, _, err := db.SupaCli.From("users").Update(updatedName, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Could not update the field"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func DeleteUserSupa(res http.ResponseWriter, req *http.Request) {
	var param = mux.Vars(req)

	data, _, err := db.SupaCli.From("users").Delete("", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("User could not delete"))
		return
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write(data)
}
