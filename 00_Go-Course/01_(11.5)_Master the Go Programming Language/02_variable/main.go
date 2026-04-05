package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const promt = "and press ENTER when ready."

func main() {
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2

	reader := bufio.NewReader(os.Stdin)
	guessTheNumber(firstNumber, secondNumber, substraction, reader)

}

func guessTheNumber(firstNumber, secondNumber, substraction int, reader *bufio.Reader) int {
	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("___________________________")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", promt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, promt)
	reader.ReadString('\n')

	fmt.Println("Now Multiply your number by", secondNumber, promt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", promt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", substraction, promt)
	reader.ReadString('\n')
	// take them trough the game

	answer := firstNumber*secondNumber - substraction
	fmt.Println("The answer is", answer)
	return answer

}
