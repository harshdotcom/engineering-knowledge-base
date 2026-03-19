package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	formatted := now.Format("02-01-2006 15:04:05")
	fmt.Println(formatted)
}
