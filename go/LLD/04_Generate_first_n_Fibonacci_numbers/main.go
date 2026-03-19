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
	fmt.Println("Write a number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	fib := make([]int, n)
	if n >= 1 {
		fib[0] = 0
	}
	if n >= 2 {
		fib[1] = 1
	}

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	fmt.Println("Fibonacci sequence:")
	fmt.Println(fib)

}
