package main

import "fmt"

//Push, Pop, Peek, IsEmpty

type Stack struct {
	items []int
}

func main() {
	fmt.Println("Doing the implemantation for stack")

	s := Stack{}

	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)

	fmt.Println(s)
	fmt.Println(s.Pop())  // 30
	fmt.Println(s.Peek()) // 20

}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
