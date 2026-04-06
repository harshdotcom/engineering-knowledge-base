package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i is", i)
		for j := 0; j < 20; j++ {
			fmt.Println(" j is", j)
		}
	}
}
