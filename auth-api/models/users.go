package models

import "github.com/google/uuid"

type Users struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	CreatedUser string    `json:"created-user"`
	Nickname    string    `json:"nickname"`
	Password    string    `json:"password"`
}
