package main

import (
	"fmt"
)

var one = "One" // Package level variable

func main() {
	var one = "this is the block level variable"
	println(one)
	myFunc()
	// newString := packageone.publicVar
	// fmt.Println(newString)
}

func myFunc() {
	fmt.Println(one)
}
