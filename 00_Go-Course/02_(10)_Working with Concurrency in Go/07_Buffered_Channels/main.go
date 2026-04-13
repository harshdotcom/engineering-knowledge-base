package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 100)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("sending ", i, "to the channel")
		ch <- i
		fmt.Println("sent", i, "to the channel")
	}

	fmt.Println("Done")
	close(ch)
}
