package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input = ""
	var tries = 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	fmt.Println("Choose a difficult level (ease 10 tries, medium 5 tries, hard 3 tries)")
	fmt.Println("Or write exit for close the game")
	input, _ = reader.ReadString('\n')
	input = strings.TrimRight(input, "\r\n")

	switch input {
	case "exit":
		return
	case "ease":
		tries = 10
	case "medium":
		tries = 5
	case "hard":
		tries = 3
	}

	fmt.Printf("You choose a %s mode.", input)
	fmt.Println("Try guess the number!")

	numberToGuess := rand.Intn(500)
	fmt.Println(numberToGuess)

	for tries > 0 {
		guessNumber, _ := reader.ReadString('\n')
		guessNumber = strings.TrimRight(guessNumber, "\r\n")
		convertedGuessNumber, _ := strconv.Atoi(guessNumber)

		if convertedGuessNumber == numberToGuess {
			fmt.Println("You win")
			return
		} else if numberToGuess > convertedGuessNumber {
			fmt.Printf("The number is higher than %d \n", convertedGuessNumber)
			tries--
		} else if numberToGuess < convertedGuessNumber {
			fmt.Printf("The number is less than %d \n", convertedGuessNumber)
			tries--
		}
	}

	fmt.Printf("You lose. The number was %d", numberToGuess)
}
