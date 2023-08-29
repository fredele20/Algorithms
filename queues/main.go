package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue - add an item to the end of the list (Queue)
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue - removes the first item from the list (Queue)
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(10)
	myQueue.Enqueue(20)
	myQueue.Enqueue(30)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}