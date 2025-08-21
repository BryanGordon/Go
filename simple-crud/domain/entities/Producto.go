package entities

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Local    string
	Cantidad int
}

func (pro *Producto) ChangePrice(newPrice float64) {
	pro.Precio = newPrice
	fmt.Println("Precio cambiado correctamente.")
}

func (pro *Producto) UpdateStock(newStock int) {
	pro.Cantidad = newStock
	fmt.Println("Se ha actualizado el stock")
}
