package main

import "fmt"

func addExpense() {

}

func listExpenses() {

}

func main() {
	const input = ""

	for {
		if input == "exit" {
			fmt.Println("Saliendo de la aplicacion...")
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
