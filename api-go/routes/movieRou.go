package routes

import (
	"api-go/data"
	"api-go/models"
	"encoding/json"
	"net/http"
)

func GetMoviesHandler(res http.ResponseWriter, req *http.Request) {
	var movies []models.Movie

	data.DB.Table("movie").Find(&movies)

	json.NewEncoder(res).Encode(&movies)

	res.Write([]byte("get movies"))
}

func GetMovie(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("get unique movie"))
}

func PostMovie(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("create a new movie"))
}

func DeleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("delete a movie"))
}
