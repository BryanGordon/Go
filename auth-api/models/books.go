package models

import "github.com/google/uuid"

type Book struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Author string    `json:"author"`
	Pages  int       `json:"pages"`
	Year   int       `json:"year"`
}
