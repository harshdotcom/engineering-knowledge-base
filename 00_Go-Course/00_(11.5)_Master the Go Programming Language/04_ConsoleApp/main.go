package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Write the name, reader")
	// userInput, _ := reader.ReadString('\n')
	// fmt.Println(userInput)

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press Any Key on the keyboard")
}
