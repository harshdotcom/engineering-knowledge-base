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

	fmt.Println("Press Any Key on the keyboard. Press ESC to quit")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Println("You Pressed", char, key)
		} else {
			fmt.Println("You Pressed", char, key)
		}

		if key == keyboard.KeyEsc {
			break
		}

		fmt.Println("Program existing")
	}
}
