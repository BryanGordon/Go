package entities // Poner el nombre del package igual al nombre de la carpeta que pertenece
import "fmt"

type Persona struct {
	Name     string
	Apellido string
	Edad     int
}

func (p Persona) Saludar() {
	fmt.Println("Hola, ", p.Name)
}

func (p *Persona) CumplirAnios() {
	p.Edad++ // No hace falta * porque infiere que es un puntero
}
