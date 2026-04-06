package main

import (
	"bufio"
	"fmt"
	"myApp/mylogger"
	"os"
	"time"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("is is", i)
	// }

	// rand.Seed(time.Now().UnixNano())

	// i := 1000

	// for i >= 100 {
	// 	i = rand.Intn(1000) + 1
	// 	fmt.Println("i is", i)

	// 	if i > 100 {
	// 		fmt.Println("is is", i, "so keeps going")
	// 	}
	// }

	// fmt.Println("Got", i, "and broke out of loop")

	// read input from the user 5 times and write it to a log

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something")
	for i := 0; i <= 5; i++ {
		fmt.Println("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}

}
