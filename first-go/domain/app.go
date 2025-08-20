package domain

import (
	"first-go/domain/entities" //Colocar primero el nombre de la carpeta del proyecto y luego la ruta a la que pertenezca el codigo
	"fmt"
)

func App() {
	persona := entities.Persona{Name: "Bryan", Apellido: "Gordon", Edad: 27}
	tipo := entities.TipoElemento{Nombre: "Alguien"}

	persona.Saludar()
	persona.CumplirAnios()
	fmt.Print(persona)
	fmt.Print("\n", tipo)

}
