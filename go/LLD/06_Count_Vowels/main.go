package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Ohh Captain My Captain")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write your world")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	wordFrequency := make(map[byte]int)

	for i := 0; i < len(input); i++ {
		ch := input[i]

		// check if character is vowel
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {

			wordFrequency[ch]++
		}
	}

	fmt.Println("Vowels Frequency")

	for k, v := range wordFrequency {
		fmt.Printf("%c : %d\n", k, v)
	}

}
