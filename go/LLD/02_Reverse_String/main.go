package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write a string:")
	input, _ := reader.ReadString('\n')

	runes := []rune(input)

	start := 0
	end := len(runes) - 1
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	ans := string(runes)
	fmt.Println(ans)

}
