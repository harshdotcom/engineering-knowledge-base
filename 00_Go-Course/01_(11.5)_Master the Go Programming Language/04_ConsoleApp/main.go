package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Write the name, reader")
	// userInput, _ := reader.ReadString('\n')
	// userInput = strings.Replace(userInput, "\n", "", -1)
	// fmt.Println(userInput)

	coffees := make(map[int]string)
	coffees[1] = "Cappuchino"
	coffees[2] = "Latte"

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("MENU")
	fmt.Println("_______")
	fmt.Println("1-Cappuchino")
	fmt.Println("2-Latte")
	fmt.Println("Q-quit the program")

	fmt.Println("Press Any Key on the keyboard. Press ESC to quit")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You choose a %s", coffees[i])

		fmt.Println(t)

		if char == 'q' || char == 'Q' {
			break
		}

		fmt.Println("Program existing")
	}
}
