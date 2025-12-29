package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var fileName = "expensives-list"

type Expensives struct {
	Id          string `json:"id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

func createFile(newData Expensives) {
	var expensives []Expensives
}

func addExpense() {

}

func listExpenses() {

}

func main() {
	var input = ""
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Iniciando app..")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	for {
		if input == "exit" {
			fmt.Println("Saliendo de la aplicacion...")
			time.Sleep(1 * time.Second)
			break
		}

		if input == "list" {
			fmt.Println("Enseñar lista")
		}

		if input == "add" {
			fmt.Println("Agregando gastos")
		}
	}

}
