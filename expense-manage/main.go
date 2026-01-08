package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
			if convertId == expensivesData.Id {
				expensives[index].Description = newDescription
				break
			} else if index == len(expensives)-1 {
				fmt.Println("No se ha podido encontrar el elemento...")
				return
			}
		}

		newJson, err := json.MarshalIndent(expensives, "", " ")

		if err != nil {
			fmt.Println("Error al crear el json.")
			return
		}

		err = os.WriteFile(fileName, newJson, 0644)

		if err != nil {
			fmt.Println("Error al modificar el json.")
			return
		}

		fmt.Println("Descripcion actualizada correctamente.")
		return
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

func delete(id string) {
	var expensives []Expensives
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error al leer el archivo.")
			return
		}

		_ = json.Unmarshal(data, &expensives)
		convertId, err := strconv.Atoi(id)

		for index, expensivesData := range expensives {
			if expensivesData.Id == convertId {
				expensives = append(expensives[:index], expensives[index+1:]...)
				break
			} else if index == len(expensives)-1 {
				fmt.Println("No se ha encontrado el elemento.")
				return
			}
		}

		newJson, err := json.MarshalIndent(expensives, "", " ")

		if err != nil {
			fmt.Println("Error al crear el json.")
			return
		}

		err = os.WriteFile(fileName, newJson, 0644)

		if err != nil {
			fmt.Println("Error al modificar el archivo.")
			return
		}

		fmt.Println("Gasto borrado de manera exitosa.")
		return
	}

	fmt.Println("No se ha encontrado el archivo.")
}

func summaryExpensives() {
	var expensives []Expensives
	summary := 0
	_, err := os.Stat(fileName)

	if err == nil {
		data, err := os.ReadFile(fileName)

		if err != nil {
			fmt.Println("Error al lerr el archivo.")
			return
		}

		_ = json.Unmarshal(data, &expensives)

		for _, expensivesData := range expensives {
			summary += expensivesData.Amount
		}

		fmt.Printf("El valor total de los gastos es: %d \n", summary)
		return
	}

	fmt.Println("No se ha encontrado el archivo")
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
			convAmountEx, _ := strconv.Atoi(amountEx)
			newExpense := Expensives{Id: globalId, Date: "2026", Description: descriptionEx, Amount: convAmountEx}
			createFile(newExpense)
		}

		if input == "list" {
			listExpenses()
		}

		if input == "update-description" {
			fmt.Println("Ingrese la nueva descripcion.")
			newDescriptionEx, _ := reader.ReadString('\n')
			newDescriptionEx = strings.TrimRight(newDescriptionEx, "\r\n")

			fmt.Println("Ingrese el ID del gasto.")
			idEx, _ := reader.ReadString('\n')
			idEx = strings.TrimRight(idEx, "\r\n")

			updateDescription(idEx, newDescriptionEx)
		}

		if input == "update-amount" {
			fmt.Println("Ingrese el nuevo monto.")
			newAmountEx, _ := reader.ReadString('\n')
			newAmountEx = strings.TrimRight(newAmountEx, "\r\n")

			fmt.Println("Ingrese el ID del gasto.")
			idEx, _ := reader.ReadString('\n')
			idEx = strings.TrimRight(idEx, "\r\n")

			updateAmount(idEx, newAmountEx)
		}

	}

}
