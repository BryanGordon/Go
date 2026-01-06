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

func listExpenses() {

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

		if input == "list" {
			newData := []Expensives
			fmt.Println("Enseñar lista")
			createFile(newData)
		}

		if input == "add" {
			fmt.Println("Agregando gastos")
		}
	}

}
