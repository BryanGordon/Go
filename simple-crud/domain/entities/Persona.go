package entities

import "fmt"

type Persona struct {
	Nombre    string
	Alias     string
	Edad      int
	Direccion string
}

func (p *Persona) ChangeAlias(newAlias string) {
	p.Alias = newAlias
	fmt.Println("Alias cambiado correctamente..")
}

func (p *Persona) ChangeDireccion(newAddress string) {
	p.Direccion = newAddress
	fmt.Println("Direccion cambiada correctamente")
}
