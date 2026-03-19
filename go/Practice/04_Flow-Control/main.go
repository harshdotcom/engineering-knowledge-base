package main

import "fmt"

func main() {
	sum := 1
	defer fmt.Println("world")
	defer fmt.Println("Hello")
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)

	// for {
	// 	fmt.Println("Runnning this loop for all the time")
	// 	fmt.Println("Not running this loop for all the time")
	// }
}

// defer : will stack the funtion and will excecute like a last in first out!!
