package models

import "github.com/google/uuid"

type Product struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Total uint      `json:"total"`
	Price float64   `json:"price"`
}
