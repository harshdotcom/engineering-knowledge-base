package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var name1, name2, name3 string

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	randomNumber := rand.Intn(30)
	fmt.Println("Here I'm printing randome number from 0 to 30: ", randomNumber)
	underRoot := findSquareRoot(7)
	fmt.Println(underRoot)
	fmt.Println(math.Pi)
	fmt.Println(addTowNumber(4, 5))
	fmt.Println(reverseFunction("Ram", "Shyam"))
	fmt.Println(reverseTwoInteger(4, 5))
	fmt.Println(name1, name2, name3)

	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value: %v\n", z, z)

}

func findSquareRoot(num float64) float64 {
	return math.Sqrt(num)
}

func addTowNumber(num1, num2 int) int {
	return num1 + num2
}

func reverseFunction(string1, string2 string) (string, string) {
	return string2, string1
}

func reverseTwoInteger(num1, num2 int) (x, y int) {
	y = num1
	x = num2
	return
}
