package models

type Book struct {
	Id     string
	Name   string `json:"name"`
	Author string `json:"author"`
	Pages  string `json:"pages"`
	Year   int    `json:"year"`
}
