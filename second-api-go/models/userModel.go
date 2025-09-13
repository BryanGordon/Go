package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Apellido string    `json:"apellido"`
	Correo   string    `json:"correo"`
	Rol      string    `json:"rol"`
}
