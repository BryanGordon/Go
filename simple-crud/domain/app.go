package domain

import (
	"fmt"
	"simple-crud/domain/entities"
)

var personDB = make([]entities.Persona, 0, 10)
var productoDB = make([]entities.Producto, 0, 5)

func savePerson(newPerson entities.Persona) {
	personDB = append(personDB, newPerson)
	fmt.Println("Usuario guardado con exito")
}

func saveProduct(newProduct entities.Producto) {
	productoDB = append(productoDB, newProduct)
	fmt.Println("El producto se ha guardado correctamente.")
}

func deletePerson(name string) {
	for index, userName := range personDB {
		if name == userName.Nombre {
			personDB = append(personDB[:index], personDB[index+1:]...) // Los ... sirven para sacar los valores de esa parte del array.
			fmt.Println("Los datos se han borrado exitosamente...\n", personDB)
			return
		}
	}

	fmt.Println("No se ha encontrado ningun dato...")
}

func deleteProduct(name string) {
	for index, nameProdu := range productoDB {
		if name == nameProdu.Nombre {
			productoDB = append(productoDB[:index], productoDB[index+1:]...)
			fmt.Println("Se ha borrado el producto correctamente\n", productoDB)
			return
		}
	}
	fmt.Println("No se ha encontrado ningun dato")
}

func updatePersonAddress(newAddress string, personName string) {
	for index, userName := range personDB {
		if personName == userName.Nombre {
			//userName.ChangeDireccion(newAddress) Funciona correctamente pero solo modifica la copia del dato, no el dato original
			personDB[index].ChangeDireccion(newAddress)
		} else {
			fmt.Println("No se ha encontrado ningun dato...")
		}
	}

}

func updateProductStock(newStock int, produtName string) {
	for index, product := range productoDB {
		if produtName == product.Nombre {
			productoDB[index].UpdateStock(newStock)
		} else {
			fmt.Println("No se ha encontrado el producto")
		}

	}

}

func App() {
	person := entities.Persona{Nombre: "Bryan", Alias: "BGE", Edad: 27, Direccion: "Cotacachi"}
	person2 := entities.Persona{Nombre: "Miguel", Alias: "Wateke", Edad: 20, Direccion: "Ibarra"}

	product1 := entities.Producto{Nombre: "PS5", Precio: 800, Local: "Quito", Cantidad: 10}
	product2 := entities.Producto{Nombre: "Iphone", Precio: 1200, Local: "Madrid", Cantidad: 300}

	savePerson(person)
	savePerson(person2)

	saveProduct(product1)
	saveProduct(product2)

	// %+v permite ver los datos de los structs sin tipos || %#v permite ver los datos de los structs con tipos
	fmt.Printf("Productos: %+v\n", productoDB)
	fmt.Printf("Personas: %+v\n", personDB)

	deletePerson("Bryan")
	deleteProduct("Iphone")

	updatePersonAddress("Atuntaqui", "Miguel")
	updateProductStock(200, "PS5")

	fmt.Printf("Productos: %+v\n", productoDB)
	fmt.Printf("Personas: %+v\n", personDB)
}
