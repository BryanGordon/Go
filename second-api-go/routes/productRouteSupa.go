package routes

import (
	"net/http"
	"second-api-go/db"
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

func CreateProductSupa(res http.ResponseWriter, req *http.Request) {

}

func UpdateProductSupa(res http.ResponseWriter, req *http.Request) {

}

func DeleteProductSupa(res http.ResponseWriter, req *http.Request) {

}
