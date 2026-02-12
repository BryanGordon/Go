package models

type TicketAvailable struct {
	Concert []Ticket `json:"concert"`
	Movie   string   `json:"type"`
	Train   string   `json:"train"`
}
