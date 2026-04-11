package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Total number of pizzas we want to attempt making
const NumberOfPizzas = 10

// Global counters to track production stats
var pizzasMade, pizzasFailed, total int

// Producer represents a pizza production system
// It uses channels to communicate between goroutines safely
type Producer struct {
	data chan PizzaOrder // channel where pizza orders are sent to consumer
	quit chan chan error // special channel used to gracefully stop production
}

// PizzaOrder represents a single pizza job
type PizzaOrder struct {
	pizzaNumber int    // unique order number
	message     string // status message (success/failure reason)
	success     bool   // whether pizza was successfully made
}

// Close gracefully stops the producer using a handshake channel pattern
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch // send a channel to request shutdown
	return <-ch  // wait for confirmation
}

// makePizza simulates the process of making a pizza
// It introduces randomness to simulate real-world failures
func makePizza(pizzaNumber int) *PizzaOrder {

	// increment order number (simulate next order)
	pizzaNumber++

	// only process valid orders
	if pizzaNumber <= NumberOfPizzas {

		// simulate preparation time (1–5 seconds)
		delay := rand.Intn(5) + 1

		// NOTE: This should ideally use fmt.Printf instead of fmt.Println
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		// random outcome generator (simulate success/failure conditions)
		rnd := rand.Intn(12) + 1

		msg := ""
		success := false

		// track statistics globally
		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		// simulate cooking time
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		// determine failure or success reason
		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		// return final pizza order result
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	// return empty order if limit exceeded
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

// pizzeria runs as a goroutine and continuously produces pizzas
// It communicates via channels with the main function
func pizzeria(pizzaMaker *Producer) {

	// keeps track of current pizza number
	var i = 0

	// infinite loop until shutdown signal is received
	for {

		// simulate pizza creation
		currentPizza := makePizza(i)

		if currentPizza != nil {

			// update counter to latest pizza number
			i = currentPizza.pizzaNumber

			select {

			// send pizza order to consumer channel
			case pizzaMaker.data <- *currentPizza:

			// listen for shutdown signal (graceful exit)
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data) // close data channel to stop consumer loop
				close(quitChan)        // acknowledge shutdown
				return
			}
		}
	}
}

func main() {

	// Seeded randomness would be better (rand.Seed(time.Now().UnixNano()))
	// This ensures different results each run

	// welcome message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("__________________________________")

	// create producer with channels
	pizzaJob := &Producer{
		data: make(chan PizzaOrder), // channel for pizza results
		quit: make(chan chan error), // channel for shutdown coordination
	}

	// start producer in background (goroutine)
	go pizzeria(pizzaJob)

	// CONSUMER: receive pizza orders from channel
	// This loop automatically stops when channel is closed
	for i := range pizzaJob.data {

		// only process valid pizza numbers
		if i.pizzaNumber <= NumberOfPizzas {

			if i.success {
				color.Green(fmt.Sprintf("Order #%d is out for delivery", i.pizzaNumber))
				color.Green(i.message)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		}
	}

	// once channel closes, execution continues here
	color.Cyan("Done making pizzas...")

	// gracefully shutdown producer
	_ = pizzaJob.Close()

	// final report of production stats
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total.",
		pizzasMade, pizzasFailed, total)

	// performance summary based on failure rate
	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day....")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was a great day!")
	}
}
