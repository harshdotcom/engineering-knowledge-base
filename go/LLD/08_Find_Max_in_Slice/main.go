package main

import (
	"fmt"
	"math"
)

func main() {
	myNumList := []int{1, 3, 4, 5, 6, 7, 8, 9}
	var maxNumber int

	for i := range myNumList {
		maxNumber = int(math.Max(float64(myNumList[i]), float64(maxNumber)))
	}

	fmt.Println("This is the max number", maxNumber)
}
