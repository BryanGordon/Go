package main //Siempre hay que traer el paquete main, tambien define el punto de entrada
import "first-go/domain"

/* Package a importar*/
// errors
// "strings"

// Este es el punto de entrada de la aplicación

// tipos de datos
// byte === unit8 | para trabajar con datos binarios
// rune === int32 | para representar un solo caracter que ocupa mas de un byte
// complex64, complex128 | valores que tienen una parte real y otra imaginaria
// nil === null
// * === Puntero
// & === Direccion de memoria

/* Funciones

func suma(a int, b int) int { // Siempre hay que colocar el tipo de dato que retorna
	result := a + b
	return result
}

func suma2(a, b int) int { // Otra forma de simplificar los tipos de datos en los parametros
	return a + b
}

func division(a, b float64) (float64, float64) { // Funcion que devuelve mas de un valor
	cociente := a / b
	resto := float64(int(a) % int(b))

	return cociente, resto
}

func divisionError(a, b float64) (float64, error) { // Funcion que es capaz de manejar errores, el error siempre va al final
	if b == 0 {
		fmt.Errorf("No se puede dividir para 0")                      // Warning por usar errors.New(), se tendria que utilizar uno de los dos y retornarlo
		return 0, errors.New("ha ocurrido un error al dividir por 0") // erros.New() crea errores; hay que pasar como primer valor un 0, porque es el valor por defecto.
	}

	cociente := a / b

	return cociente, nil // retornamos nil ya que no hemos tenido ningun error
}

func showNombres(nombres ...string) { // Funcion de parametros variables, los 3 puntos indican que no se sabe el numero total de parametros
	for _, nombres := range nombres {
		fmt.Println(nombres)
	}
}

func contador() func() int { // Funcion de clausura (funcion que retorna otra funcion)
	cont := 0

	return func() int {
		cont++
		return cont
	}
}

type rectangulo struct { // Forma de definir tipos; Las clases no existen en go, se tiene que crear un tipo y luego una funcion asociada
	ancho, alto float64
}

func (rec rectangulo) area() float64 { // El primer parentesisi representa que esta asociado a un type o interface
	return rec.alto * rec.ancho
}

*/

/* Punteros
// Para modificar el elemento original (No recomendable, mala practica)
// Nunca se debe cambiar el argumento directamente (Hacer copia)

// Las direcciones de memoria no funcionan igual en go, sino que se hace una copia del argumento
func incrementarValor(value *int) { // Para poder acceder al argumento en memoria hay que utilizar *
	*value++ // Tambien hay que colocar el * para que funcione
}

// new() permite crear un puntero con un valor predeterminado del tipo correspondiente
*/

/* Interfaces
// No se implementan (al menos manualmente)
// Se cumplen

type forma interface {
	area() float64
}

type circle struct {
	radio float64
}

func (cir circle) area() float64 {
	return math.Pi * cir.radio * cir.radio
}

func imprimirArea(f forma) {
	fmt.Printf("El area es: %.2f\n", f.area()) // %.2f significa que solo va a utilizar dos decimales
}

var cualquiera interface{} // Forma de declarar una interfaz que sea cualquier cosa/generica; como un type any (TS)

*/

/* Concurrencia y polimorfismo
// Todo el codigo se ejecuta de forma sincrona
// Se usa la palabra reservada go
// Al usar la palabra go al lado de un codigo se crea una GoRoutine
// GoRoutine es un hilo de ejecucion ligero virtual
// Channel permite comunicar cosas entre GoRoutines

func saludar(canal chan<- string) { // Tipo chan define el Channel y la flecha indica que recibe un string (importante definir bien el tipo de dato)
	time.Sleep(1 * time.Second)        // Se ejecuta cada segundo
	canal <- "Hola desde la GoRoutine" // Se envia el string hacia el canal
}

func imprimirSaludo(canal <-chan string) { // Asi se recibe un dato desde la GoRoutine
	fmt.Println("Esperando mensaje...")
	mensaje := <-canal
	fmt.Println(mensaje)
}

*/

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
	otherSlice := make([]int, 5, 10) 				// Permite crear slices con una longuitud y una capacidad respectivamente, con un solo parametro crea la misma longuitud y capacidad

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

	/* Declaracion y utilizacion de funciones

	cociente, resto := division(20, 2) // Cada variable obtiene el valor que retorna la funcion respectivamente

	cocienteError, error := divisionError(10, 3) // Se hace de las misma manera pero despues hay que controlar el error
	cocienteError2, _ := divisionError(10, 3)    // Se puede evitar usar el error

	if error != nil {
		// panic("La funcion divisionError produjo un error") // La funcion panic() cierra aborta completamente la ejecucion del programa
		fmt.Printf("Hubo un error %s", error)
		return
	}

	fmt.Println("Division", cociente, resto)
	fmt.Println("Division con error completa", cocienteError)
	fmt.Println("Division con error solo resultado", cocienteError2)

	conta := contador()               // La funcion se ejecuta sin mas
	fmt.Println("Contador:", conta()) // Hay que usar la variable creada anteriormente como si fuera una funcion
	fmt.Println("Contador:", conta())
	fmt.Println("Contador:", conta())
	fmt.Println("Contador:", conta())
	fmt.Println("Contador:", conta())

	rectan := rectangulo{alto: 10, ancho: 5} // Definir primero los valores en el type/interface
	fmt.Println(rectan.area())               // La variable tiene acceso a las funciones del type/interface

	*/

	/* Uso de punteros
	valor := 10
	fmt.Println("Puntero antes del cambio:", valor)

	incrementarValor(&valor) // Se tiene que poner la & para hacer referencia a la direccion de memoria
	fmt.Println("Puntero despues del valor:", valor)

	puntero := new(int)                                 // Crea un puntero hacia un espacio de memoria
	fmt.Println("Espacio de memoria con new:", puntero) // Imprime la direccion de memoria
	fmt.Println("Valor de memoria con new:", *puntero)  // Imprime el valor que hay en la direccion de memoria
	*puntero = 20
	fmt.Println("Valor de memoria con new, cambiado:", *puntero)
	*/

	/* Uso de interfaces
	circulo := circle{radio: 5}

	imprimirArea(circulo) // La funcion infiere que circulo es una forma porque poseen un metodo con el mismo nombre, al ser diferente marcaria un error; el nombre debe ser exacto
	*/

	/* Uso de Concurrencia, polimorfismo
	nuevoCanal := make(chan string) // Con make se crea un canal
	go saludar(nuevoCanal)          // Se ejecuta la GoRoutine
	imprimirSaludo(nuevoCanal)

	otroCanal := make(chan int)
	go func() { // Se crea una GoRoutine que ejecuta una funcion
		for i := 0; i < 3; i++ {
			otroCanal <- i
		}
		close(otroCanal) // Se tiene que cerrar el canal
	}() // Se coloca () para ejecutar la funcion

	for num := range otroCanal { // Se puede utilizar los valores ejecutados desde la GoRutine/Channel
		fmt.Println("Valor recibido desde la GoRutine:", num)
	}

	// Un canal es Writer y Reader ya que se puede leer y escribir datos (recibir y enviar)
	// Un ejemplo de solo Writer es range

	// Ejemplo con uso de mutex
	// Mutex permite reserva recursos cuando un mismo canal se esta ejecutando mas de una vez
	var contar int
	var mu sync.RWMutex // Es un lector y escritor de mutex (Reader/Writer)

	// Writer Elemento que permite escribir/enviar datos
	go func() {
		for i := 0; i < 3; i++ {
			mu.Lock()   // Bloquea el codigo que se ejecuta dentro de la funcion
			contar++    // No se va a poder utilizar "contar" hasta que la ejecucion no termine
			mu.Unlock() // Desbloquea el uso del codigo posterior
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader Elemento que permite escribir/recibir datos
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mu.RLock() // Permite bloquear la lectura
				fmt.Printf("Leyendo desde la GoRoutine %d: %d\n", id, contar)
				mu.RUnlock() // Permite desbloquear la lectura
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
	// En consola se ve como se repiten los datos, eso es porque estan bloqueados

	*/

	/* uso de packages*/
	domain.App()

}
