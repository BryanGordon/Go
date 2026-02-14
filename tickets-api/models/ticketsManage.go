package models

type TicketAvailable struct {
	Concert []Ticket `json:"concert"`
	Movie   []Ticket `json:"movie"`
	Train   []Ticket `json:"train"`
}
