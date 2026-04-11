package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
	color     *color.Color
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0, color: color.New(color.FgRed)},
	{name: "Socrates", leftFork: 0, rightFork: 1, color: color.New(color.FgGreen)},
	{name: "Aristotle", leftFork: 1, rightFork: 2, color: color.New(color.FgYellow)},
	{name: "Pascal", leftFork: 2, rightFork: 3, color: color.New(color.FgBlue)},
	{name: "Locke", leftFork: 3, rightFork: 4, color: color.New(color.FgMagenta)},
}

var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second

func main() {
	fmt.Println("Dining philosopher problem")
	fmt.Println("__________________________")

	dine()

	fmt.Println("The table is empty.")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated, i)
	}

	wg.Wait()
}

func diningProblem(
	p Philosopher,
	wg *sync.WaitGroup,
	forks map[int]*sync.Mutex,
	seated *sync.WaitGroup,
	index int,
) {
	defer wg.Done()

	p.color.Printf("%s is seated at the table.\n", p.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if index%2 == 0 {
			forks[p.leftFork].Lock()
			p.color.Printf("\t%s takes the left fork.\n", p.name)

			forks[p.rightFork].Lock()
			p.color.Printf("\t%s takes the right fork.\n", p.name)
		} else {
			forks[p.rightFork].Lock()
			p.color.Printf("\t%s takes the right fork.\n", p.name)

			forks[p.leftFork].Lock()
			p.color.Printf("\t%s takes the left fork.\n", p.name)
		}

		p.color.Printf("\t%s is eating.\n", p.name)
		time.Sleep(eatTime)

		forks[p.leftFork].Unlock()
		forks[p.rightFork].Unlock()

		p.color.Printf("\t%s is thinking.\n", p.name)
		time.Sleep(thinkTime)

		p.color.Printf("\t%s put down the forks.\n", p.name)
	}

	p.color.Println(p.name, "is satisfied.")
	p.color.Println(p.name, "left the table")
}
