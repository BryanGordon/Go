package routes

import (
	"net/http"
	"second-api-go/db"

	"github.com/gorilla/mux"
)

func GetDataSupa(res http.ResponseWriter, req *http.Request) {
	data, _, err := db.SupaCli.From("products").Select("*", "", false).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Data not found"))
		return
	}

	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func SearchProductSupa(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)

	data, _, err := db.SupaCli.From("products").Select("*", "", false).Eq("name", param["name"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Product not found"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func CreateProductSupa(res http.ResponseWriter, req *http.Request) {

}

func UpdateProductSupa(res http.ResponseWriter, req *http.Request) {

}

func DeleteProductSupa(res http.ResponseWriter, req *http.Request) {

}
