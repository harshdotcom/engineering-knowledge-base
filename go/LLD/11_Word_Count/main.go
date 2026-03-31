package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please write the string")
	input, _ := reader.ReadString('\n')
	wordMap := make(map[byte]int)
	input = strings.TrimSpace(input)
	for i := range input {
		wordMap[(input[i])]++
	}

	for k, v := range wordMap {
		fmt.Printf("%c : %d\n", k, v)
	}
}
