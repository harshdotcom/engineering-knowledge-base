package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance

	var bankBalance int
	var balance sync.Mutex

	// print out starting value
	fmt.Printf("Initial account balance: $d.00", bankBalance)
	fmt.Println()

	// define weekely revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time Job", Amount: 50},
		{Source: "Investment", Amount: 100},
	}

	wg.Add(len(incomes))
	// loop through 52 weeks and print out how much is made : keep running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from %s\n0", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
	// print out final balance
}
