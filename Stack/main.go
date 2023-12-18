package main

import "fmt"

type stack struct {
	items []int
}

func (s *stack) push(item int) {
	fmt.Println("Push item", item)
	s.items = append(s.items, item)
}

func (s *stack) pop() {

	l := len(s.items)

	if l > 0 {
		removedItem := s.items[l-1]
		s.items = s.items[:(l - 1)]
		fmt.Println("Pop item", removedItem)
	} else {
		fmt.Println("Stack is empty !!!!")
	}

}

func main() {

	var myStack stack

	myStack.push(1)
	fmt.Println(myStack.items)
	myStack.push(3)
	fmt.Println(myStack.items)
	myStack.push(5)
	fmt.Println(myStack.items)

	myStack.pop()
	fmt.Println(myStack.items)
	myStack.pop()
	fmt.Println(myStack.items)
	myStack.pop()
	fmt.Println(myStack.items)
	myStack.pop()
	fmt.Println(myStack.items)

}
