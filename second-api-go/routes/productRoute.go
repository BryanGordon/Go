package routes

import (
	"encoding/json"
	"net/http"
	"second-api-go/db"
	"second-api-go/models"
)

func AddProduct(res http.ResponseWriter, req *http.Request) {
	var newProduct models.Product

	json.NewDecoder(req.Body).Decode(&newProduct)

	addedProduct := db.DB.Create(&newProduct)
	err := addedProduct.Error

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
	}

	json.NewEncoder(res).Encode(&newProduct)
}

func GetProducts(res http.ResponseWriter, req *http.Request) {
	var allProducts []models.Product

	db.DB.Find(&allProducts)

	json.NewEncoder(res).Encode(&allProducts)
}

func UpdateProduct(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("update product"))
}

func SearchProduct(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("search product"))
}

func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("delete product"))
}
