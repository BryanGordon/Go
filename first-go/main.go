package main //Siempre hay que traer el paquete main, tambien define el punto de entrada
import (
	"fmt"
	"strings"
)

// Este es el punto de entrada de la aplicación

// tipos de datos
// byte === unit8 | para trabajar con datos binarios
// rune === int32 | para representar un solo caracter que ocupa mas de un byte
// complex64, complex128 | valores que tienen una parte real y otra imaginaria

func main() {
	fmt.Println("Hola go")

	//Numeros
	entero := 10
	decimal := 3.14
	suma := entero + int(decimal)

	// Texto
	mensaje := "Hola, "
	concatenado := mensaje + "Bryan"
	toMayus := strings.ToUpper(concatenado)

	//Array y Slices
	arrayFijo := [3]int{1, 2, 3}             // Hay que colocar simpre un tamaño por defecto
	sliceVariable := []int{4, 5, 6}          // Forma de instanciar una array dinamico, Slice es otro forma de array que solo presenta un trozo
	sliceVariable = append(sliceVariable, 7) // append sirve para agregar elementos
	fmt.Println(len(arrayFijo))              // len muestra cual es el total de elementos del array
	fmt.Println(cap(arrayFijo))              // cap() Muestra cual es la capacidad total del array

	// Maps
	diccionario := map[string]int{ // Se typa primero la clave "string" y luego el tipo de dato de los valores "int"
		"precio":  1,
		"precio2": 2,
	}

	// Structs | Estructuras | Objetos
	type Persona struct { // colocar siempre struct para indentificar que es un objeto
		Nombre string // Colocar mayusculas en todos los elementos que se quiera exportar/importar, caso contrario en minusculas
		Edad   int
	}

}
