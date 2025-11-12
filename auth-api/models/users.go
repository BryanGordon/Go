package models

import "github.com/google/uuid"

type Users struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Rol         string    `json:"rol"`
	Email       string    `json:"email"`
	CreatedUser string    `json:"createdUser"`
	Nickname    string    `json:"nickname"`
	Password    string    `json:"password"`
}
