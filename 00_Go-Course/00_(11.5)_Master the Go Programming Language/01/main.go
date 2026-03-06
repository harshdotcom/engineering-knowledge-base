package main

import "fmt"

func main() {
	fmt.Println("Hello World!!")
	fmt.Print("This is some text")
	fmt.Print("This is some more text")
	sayHelloworld("Harsh Jha Jha")
}

func sayHelloworld(whattosay string) {
	fmt.Println(whattosay)
}
