package main

import (
	"fmt"
	"strings"
)

// Channels : Pushing data to go routine or taking data from the go routine

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// Create two channels
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	fmt.Println("Type something and press ENTER (enter Q to quiet)")

	for {
		fmt.Println("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if "q" == strings.ToLower(userInput) {
			break
		}

		ping <- userInput

		// wait for the response

		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels")

	close(ping)
	close(pong)
}
