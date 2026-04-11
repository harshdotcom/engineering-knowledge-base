package main

import (
	"fmt"
	"math"
)

func main() {
	answer := 7 + 3*4
	fmt.Println(answer, "Printing the anser")

	var radius = 12.0
	area := math.Pi * radius * radius

	fmt.Println("Area", area)
}
