package routes

import (
	"net/http"
	"second-api-go/db"

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
	data, _, err := db.SupaCli.From("users").Select("*", "", false).Eq("name", param["name"]).Execute()

	/* Si el usuario no existe devuelve un array vacio, arreglar eso*/
	/* Probar si la api key publica se conecta y hace las mismas funciones que la actual*/
	/* Probar la primera api key generica, tal vez funciona*/

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

}

func UpdateUserSupa(res http.ResponseWriter, req *http.Request) {

}

func DeleteUserSupa(res http.ResponseWriter, req *http.Request) {

}
