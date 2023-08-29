package main

import "fmt"


type Stack struct {
	items []int
}


// Push - adds an item to the end of the list (stack)
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop - remove the last item from the list (stack)
func (s *Stack) Pop() int {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}

