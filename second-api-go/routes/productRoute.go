package routes

import (
	"encoding/json"
	"net/http"
	"second-api-go/db"
	"second-api-go/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	var updateName models.Product
	var itemId models.Product
	param := mux.Vars(req)

	db.DB.Where("id = ?", param["id"]).First(&itemId)

	if itemId.Id == uuid.Nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Product not found"))
		return
	}

	json.NewDecoder(req.Body).Decode(&updateName)

	value := db.DB.Model(&updateName).Where("id = ?", itemId.Id).Update("name", updateName.Name)
	err := value.Error

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Can not update field"))
		return
	}

	json.NewEncoder(res).Encode(&updateName.Name)
}

func SearchProduct(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("search product"))
}

func DeleteProduct(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("delete product"))
}
