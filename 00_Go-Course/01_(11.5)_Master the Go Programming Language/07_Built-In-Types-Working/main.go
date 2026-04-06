package main

import (
	"fmt"
	"sort"
)

// basic type (number, strings, boolens)
var myInt int
var myUnit uint
var myFloat float32
var myFloat64 float64

// aggregate types (array, struct)

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// Refrence Type : Pointer, Slices, Maps, Functions, channels

// Pointer is nothing more it just point to specific memory or location

func main() {
	var myStrings [3]string
	myStrings[0] = "Cat"
	myStrings[1] = "Cat"
	myStrings[2] = "Cat"
	fmt.Println(myStrings)

	// var myCar Car
	// myCar.NumberOfTires = 2
	// myCar.Luxury = false
	// myCar.Make = "ford"

	// myCar := Car{
	// 	NumberOfTires: 4,
	// 	Luxury:        false,
	// 	BucketSeats:   true,
	// 	Make:          "Volvo",
	// 	Model:         "XC()",
	// 	Year:          2019,
	// }

	// fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

	// var myInt int
	// myInt = 10
	// newPointer := &myInt
	// *newPointer = 20
	// fmt.Println(*newPointer)

	// It help us to change the value of variable withour even passing it

	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for _, x := range animals {
		fmt.Println("Animals Name", x)
	}

	fmt.Println("First tow element of the array is", animals[0:2])
	fmt.Println("Size of slice", len(animals))

	sort.Strings(animals)

	fmt.Println(animals)

}

// func DeleteFromSlice()
