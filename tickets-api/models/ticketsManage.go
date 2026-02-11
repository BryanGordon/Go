package models

type TicketAvailable struct {
	Concert string `json:"concert"`
	Movie   string `json:"type"`
	Train   string `json:"train"`
}
