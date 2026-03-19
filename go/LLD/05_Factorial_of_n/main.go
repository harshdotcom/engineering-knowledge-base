package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write a number")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, _ := strconv.Atoi(input)
	fmt.Println(findFactorial(n))
}

func findFactorial(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	return number * findFactorial(number-1)
}
