package main

import "fmt"

// Interface type

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (d *Dog) Says() string {
	return d.Sound
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

func (d *Cat) Says() string {
	return d.Sound
}

func main() {

	// Ask a riddle
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	var cat Cat
	cat.Name = "car"
	cat.NumberOfLegs = 4
	cat.Sound = "Meow"
	cat.HasTail = true

	Riddle(&cat)

}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs, What animal is it?`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
