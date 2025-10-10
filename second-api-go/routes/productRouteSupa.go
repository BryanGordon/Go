package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"second-api-go/db"
	"second-api-go/models"

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
	var newProduct models.Product
	json.NewDecoder(req.Body).Decode(&newProduct)

	data, _, err := db.SupaCli.From("products").Insert(newProduct, false, "", "", "").Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateProductSupa(res http.ResponseWriter, req *http.Request) {
	var newName models.Product
	// var infoProduct models.Product
	var param = mux.Vars(req)
	var err error

	infoproduct, _, err := db.SupaCli.From("products").Select("*", "", false).Eq("id", param["id"]).Execute()
	convertedInfo := string(infoproduct)
	fmt.Print(convertedInfo)

	json.NewDecoder(req.Body).Decode(&newName)
	fmt.Print(newName)

	data, _, err := db.SupaCli.From("products").Update(newName.Name, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Could not update product name"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func DeleteProductSupa(res http.ResponseWriter, req *http.Request) {
	var param = mux.Vars(req)

	data, _, err := db.SupaCli.From("products").Delete("", "").Eq("name", param["name"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Could not delete product"))
		return
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write(data)
}
