package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type User struct {
	name    string
	age     int
	address string
	pinCode string
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {

	// ++++++++++++++++++  Pointers ++++++++++++++++++

	// How Pointers work in Go!!
	var p *int
	i := 20
	p = &i
	var z = 10
	y := &z
	fmt.Println(*p, *y)

	// ++++++++++++++++++  Struct   ++++++++++++++++++
	// How structs works in Go! : A Struct is a collection of fields
	v := Vertex{}
	v.X = 2
	v.Y = 8
	newPointer := &v
	v.X = 7
	user1 := User{"Harsh", 24, "Agra", "284162"}
	user2 := User{"Harsh", 24, "Agra", "284162"}
	fmt.Println(Vertex{2, 3}, v, user1, user2, *&newPointer.X)

	// ++++++++++++++++++  Struct Literals   ++++++++++++++++++
	fmt.Println(v1, p, v2, v3, *p)

	// ++++++++++++++++++  Array   ++++++++++++++++++

	var myArray [5]string
	myArray[0] = "Hello"
	myArray[1] = "World!!"

	fmt.Println(myArray)

	// ++++++++++++++++++  Slices   ++++++++++++++++++
	//Slices are like references to arrays

	primes := [5]int{2, 3, 4, 5, 6}
	var s []int = primes[:]
	fmt.Println(s)

	// ++++++++++++++++++  Slices litrels   ++++++++++++++++++
	q := []int{1, 2, 3, 4, 5}
	fmt.Println(q)

	m := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(m)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// How to declare slice of slice

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println(board)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
