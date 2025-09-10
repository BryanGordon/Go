package routes

import (
	"api-go/data"
	"api-go/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMoviesHandler(res http.ResponseWriter, req *http.Request) {
	var movies []models.Movie

	data.DB.Table("movie").Find(&movies)

	json.NewEncoder(res).Encode(&movies)

	res.Write([]byte("get movies"))
}

func GetMovie(res http.ResponseWriter, req *http.Request) {
	var movie models.Movie
	params := mux.Vars(req)
	data.DB.First(&movie, params["Id"])

	if movie.Id == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Movie not found"))
		return
	}

	json.NewEncoder(res).Encode(&movie)
}

func PostMovie(res http.ResponseWriter, req *http.Request) {
	var newMovie models.Movie

	json.NewDecoder(req.Body).Decode(&newMovie)

	newDataMovie := data.DB.Table("movie").Create(&newMovie)
	err := newDataMovie.Error

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
	}

	json.NewEncoder(res).Encode(&newMovie)

	res.Write([]byte("create a new movie"))
}

func DeleteMovie(res http.ResponseWriter, req *http.Request) {
	var movie models.Movie
	params := mux.Vars(req)
	data.DB.First(&movie, params["Id"])

	if movie.Id == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("User not found"))
		return
	}

	data.DB.Unscoped().Delete(&movie)
	res.WriteHeader(http.StatusOK)

	res.Write([]byte("delete a movie"))
}
