package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

var fileName = "expensives-list"

type Expensives struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

func createFile(newData Expensives) {
	var expensives []Expensives
	isExist := false

	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)
		isExist = true

		if err == nil && len(data) > 0 {
			_ = json.Unmarshal(data, &expensives)
		}
	}

	expensives = append(expensives, newData)
	newJson, err := json.MarshalIndent(expensives, "", " ")

	if err != nil {
		fmt.Println("An error ocurred creating the json file")
		panic(err)
	}

	err = os.WriteFile(fileName, newJson, 0644)

	if err != nil {
		fmt.Println("Error ocurred writing the file")
		panic(err)
	}

	if isExist {
		fmt.Println("Se ha agregado un gasto..")
	} else {
		fmt.Println("Archivo creado exitosamente..")
	}
}

/*
func addExpense(newExpense Expensives) {
	var expenses []Expensives

}
*/

func listExpenses() {
	var expensives []Expensives
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error al leer el archivo")
			return
		}
		_ = json.Unmarshal(data, &expensives)

		for _, data := range expensives {
			fmt.Printf("-- ID: %d , Date: %s , Description: %s , Cantidad: $%d . \n", data.Id, data.Date, data.Description, data.Amount)
		}

		return
	}

	fmt.Println("No se ha encontrado el archivo")
}

func updateDescription(id string, newDescription string) {
	var expensives []Expensives
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error al leer el archivo.")
			return
		}

		_ = json.Unmarshal(data, &expensives)
		convertId, _ := strconv.Atoi(id)

		for index, expensivesData := range expensives {
			if expensives[index].Id == convertId {

			}
		}
	}

	fmt.Println("\033[31mNo se ha encontrado el archivo.\033[0m")
}

func updateAmount(id string, newAmount string) {
	var expenses []Expensives
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error al leer el archivo.")
			return
		}

		_ = json.Unmarshal(data, &expenses)
		convertedId, _ := strconv.Atoi(id)

		for index, expensivesData := range expenses {
			if convertedId == expensivesData.Id {
				convertedAmount, _ := strconv.Atoi(newAmount)
				expenses[index].Amount = convertedAmount
				break
			} else if index == len(expenses)-1 {
				fmt.Println("No se ha encontrado el elemento...")
				return
			}
		}

		newJson, err := json.MarshalIndent(expenses, "", " ")

		if err != nil {
			fmt.Println("Ocurrio un error.")
			return
		}

		err = os.WriteFile(fileName, newJson, 0644)

		if err != nil {
			fmt.Println("Error al modificar el json")
			return
		}

		fmt.Println("Monto actualizado correctamente.")
		return
	}

	fmt.Println("No se ha encontrado el archivo.")
}

func main() {
	var input = ""
	globalId := 1
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Iniciando app..")

	for {
		fmt.Println("Ingrese un comando")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Saliendo de la aplicacion...")
			time.Sleep(1 * time.Second)
			break
		}

		if input == "add" {
			fmt.Println("Ingrese la descripcion del gasto")
			descriptionEx, _ := reader.ReadString('\n')
			descriptionEx = strings.TrimRight(descriptionEx, "\r\n")

			fmt.Println("Ingrese el monto")
			amountEx, _ := reader.ReadString('\n')
			amountEx = strings.TrimRight(amountEx, "\r\n")
			// Cambiar el tipo de data amountEx a integer
			newExpense := Expensives{Id: globalId, Date: "2026", Description: descriptionEx, Amount: amountEx}
			createFile(newExpense)
		}

		if input == "list" {
			listExpenses()
		}

		if input == "update" {
			fmt.Println("Agregando gastos")
		}
	}

}
