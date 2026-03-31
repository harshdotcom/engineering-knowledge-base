package main

import "fmt"

func main() {

	myList := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6, 7}
	myNewMap := make(map[int]int)
	uniqueList := []int{}

	for i := range myList {
		if myNewMap[myList[i]] > 0 {
			continue
		} else {
			uniqueList = append(uniqueList, myList[i])
		}
		myNewMap[myList[i]]++
	}

	fmt.Println("Unique List", uniqueList)
}
