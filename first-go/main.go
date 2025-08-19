package main //Siempre hay que traer el paquete main, tambien define el punto de entrada
import (
	"fmt"
	// "strings"
)

// Este es el punto de entrada de la aplicación

// tipos de datos
// byte === unit8 | para trabajar con datos binarios
// rune === int32 | para representar un solo caracter que ocupa mas de un byte
// complex64, complex128 | valores que tienen una parte real y otra imaginaria

func main() {
	/* Basics
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

	persona1 := Persona{Nombre: "Bryan", Edad: 27}
	persona2 := Persona{Nombre: "Ivan"} // La edad toma el valor de 0

	// Imprimir resultados
	fmt.Println("Suma", suma)
	fmt.Println("Mensaje", toMayus)
	fmt.Println("Array Fijo", arrayFijo)
	fmt.Println("Slice Variable", sliceVariable)
	fmt.Println("Diccionario", diccionario)
	fmt.Println("Persona", persona1)
	fmt.Println("Persona 2", persona2)
	*/

	/* Condicionales, bucles

	//Condicionales
	edad := 20

	if edad >= 18 {
		fmt.Println("Es mayor de edad")
	} else {
		fmt.Println("Es menor de edad")
	}

	//Bucles
	for i := 0; i < 3; i++ {
		fmt.Printf("Iteracion: %d\n", i) //El %d sirve para asignar un valor dentro de un string, valor se coloca despues de la coma
	}

	// Range uso del rango del slice
	slice := []string{"Algo", "Hola", "Otro"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value) // %s para colocar valores de tipo string
	}

	for index, _ := range slice { // Si no se va a usar el valor, se puede colocar un guion bajo para ignorarlo
		fmt.Printf("Index con _: %d\n", index)
	}

	for index := range slice { // Simplificado (sin "_")
		fmt.Printf("Solo Index: %d\n", index)
	}

	for _, value := range slice { // No se pude simplificar
		fmt.Printf("Segundo Valor: %s\n", value)
	}

	// Defer Permite ejecutar codigo siempre al final de la compilacion del metodo principal/asociado
	defer fmt.Println("Uso del defer")

	fmt.Println("Prueba para el defer")
	fmt.Println("Prueba para el defer")
	fmt.Println("Prueba para el defer")

	*/

	/* Funciones */

}

func suma(a int, b int) {
	result := a + b
	fmt.Println("Alog")
	return result
}
