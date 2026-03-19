package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write a string to check it is palindrome or not??")
	input, _ := reader.ReadString('\n')
	checkPalindrome(input)

}

func checkPalindrome(ourString string) {
	start := 0
	end := len(ourString) - 1

	for start < end {
		if ourString[start] != ourString[end] {
			fmt.Println("It is not a palindrome")
			return
		}
		start++
		end--
	}
	fmt.Println("It is a palindrome")
}
