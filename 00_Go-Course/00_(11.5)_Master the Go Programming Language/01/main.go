package main

import (
	"bufio"
	"fmt"
	doctor "myapp/02_doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var whattosay string
	whattosay = doctor.Intro()
	fmt.Println(whattosay)

	for {
		userInput, _ := reader.ReadString('\n')
		response := doctor.Response(userInput)

		userInput = strings.TrimSpace(userInput)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(response)
		}

	}

	// fmt.Println("Hello World!!")
	// fmt.Print("This is some text")
	// fmt.Print("This is some more text")
	// // just print the text, no line breaking
	// // var whattosay = "Hello world"
	// whattosay := "Hello world!"
	// sayHelloworld(whattosay)
}

func sayHelloworld(whattosay string) {
	fmt.Println(whattosay)
}
