package routes

import (
	"auth-api/connections"
	"auth-api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func ListBooks(res http.ResponseWriter, req *http.Request) {
	data, _, err := connections.Client.From("books").Select("*", "", false).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Error al solicitar los datos"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func AddBooks(res http.ResponseWriter, req *http.Request) {
	var newBook models.Book

	json.NewDecoder(req.Body).Decode(&newBook)

	data, _, err := connections.Client.From("books").Insert(newBook, false, "", "", "").Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("No se ha podido añadir el libro..."))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func SearchBook(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)

	data, _, err := connections.Client.From("users").Select("*", "", false).Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Book not found"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateNameBooks(res http.ResponseWriter, req *http.Request) {
	var newName models.Book
	param := mux.Vars(req)

	json.NewDecoder(req.Body).Decode(&newName)
	updatedName := map[string]any{
		"name": newName.Name,
	}

	data, _, err := connections.Client.From("books").Update(updatedName, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't update name"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateBookPages(res http.ResponseWriter, req *http.Request) {
	var newPages models.Book
	param := mux.Vars(req)

	json.NewDecoder(req.Body).Decode(&newPages)
	updatedPages := map[string]any{
		"pages": newPages.Pages,
	}

	data, _, err := connections.Client.From("books").Update(updatedPages, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't update pages"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func UpdateBookAuthor(res http.ResponseWriter, req *http.Request) {
	var newAuthor models.Book
	param := mux.Vars(req)

	json.NewDecoder(req.Body).Decode(&newAuthor)
	updatedAuthor := map[string]any{
		"author": newAuthor.Author,
	}

	data, _, err := connections.Client.From("books").Update(updatedAuthor, "", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't update author"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func RemoveBook(res http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)

	data, _, err := connections.Client.From("books").Delete("", "").Eq("id", param["id"]).Execute()

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Couldn't delete the book"))
		return
	}

	res.WriteHeader(http.StatusAccepted)
	res.Write(data)
}
