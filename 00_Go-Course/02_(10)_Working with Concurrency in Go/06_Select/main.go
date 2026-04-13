package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is From server 1"
	}
}

func server2(ch chan string) {

	for {
		time.Sleep(3 * time.Second)
		ch <- "This is From server 2"
	}
}

func main() {

	fmt.Println("Select with channels")
	fmt.Println("____________________")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Println("Case one:", s1)
		case s2 := <-channel2:
			fmt.Println("Case two", s2)

		}
	}

}
