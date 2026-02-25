package models

type Ticket struct {
	Id     string `json:"id"` // cambiar a uuid
	Type   string `json:"type"`
	Number string `json:"number"`
}
