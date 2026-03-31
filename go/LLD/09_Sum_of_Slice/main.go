package main

import "fmt"

func main() {
	myNumList := []int{1, 3, 4, 5, 6, 7, 8, 9}
	var sumofNumber int

	for i := range myNumList {
		sumofNumber += myNumList[i]
	}

	fmt.Println("This is the sum of all number", sumofNumber)
}
