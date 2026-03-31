package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	println("Write a number")
	fmt.Scan(&input)
	input = int(input)

	for i := 2; i <= int(math.Sqrt(float64(input))); i++ {
		if input%i == 0 {
			fmt.Println("Not a prime Number")
			return
		}
	}

	fmt.Println("You found a prime number")

}
