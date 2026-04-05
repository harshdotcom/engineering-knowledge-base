package main

import "fmt"

type Queue struct {
	items []int
}

//Enqueue, Dequeue, Front

func main() {
	q := Queue{}

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println(q)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Front())

}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	if (len(q.items)) == 0 {
		return -1
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front
}

func (q *Queue) Front() int {
	if (len(q.items)) == 0 {
		return -1
	}
	return q.items[0]
}
