package models

type Movie struct {
	Id       uint    `json:"id"`
	Title    string  `json:"title"`
	Years    int     `json:"years"`
	Director string  `json:"director"`
	Duration int     `json:"duration"`
	Poster   string  `json:"poster"`
	Rate     float64 `json:"rate"`
}
