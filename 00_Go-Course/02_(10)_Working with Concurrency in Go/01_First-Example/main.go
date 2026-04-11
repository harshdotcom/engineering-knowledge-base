package main

import (
	"fmt"
	"sync"
	"time"
)

// Go routine : Very light weight threads and they all manage by the go co-rotine scheduler

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	// fmt.Println("Hello world")

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"theta",
		"epsiton",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)

	}

	wg.Wait()
	wg.Add(1)
	go printSomething("This is the first things to be printed", &wg)

	time.Sleep(1 * time.Second)
}
