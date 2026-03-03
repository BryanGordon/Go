package models

import "github.com/gofrs/uuid"

type Ticket struct {
	Id     uuid.UUID `json:"id"`
	Type   string    `json:"type"`
	Number string    `json:"number"`
}
