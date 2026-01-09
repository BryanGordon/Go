package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input = ""
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Welcome to the Number Guessing Game!")
		fmt.Println("I'm thinking of a number between 1 and 100.")
		fmt.Println("You have 5 chances to guess the correct number.")
		fmt.Println("Choose a difficult level (ease 10 tries, medium 5 tries, hard 3 tries)")
		fmt.Println("Or write exit for close the game")
		input, _ = reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")

		if input == "exit" {
			break
		}
	}
}
