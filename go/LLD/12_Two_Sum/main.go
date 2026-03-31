package main

import "fmt"

func main() {
	fmt.Println("Two Sum Problem!!")
	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myNumMap := make(map[int]int)
	target := 100
	for i := range myList {
		goalNum := target - myList[i]
		if myNumMap[goalNum] > 0 {
			fmt.Println("Yeah we found a pair", myNumMap[goalNum]-1, i)
			return
		}
		myNumMap[myList[i]]++
	}

	fmt.Println("Sorry to say but we have't any pair which follow your condition!!")
}
