package main

import "fmt"

func main() {
	var s Stack
	s.Push("a")
	s.Print()

	s.Pop()
	s.Print()

	s.Pop()
}

// Stack structure
type Stack []string

// Push new value to stack
func (s *Stack) Push(element string) {
	fmt.Println("Pushing: ", element)
	*s = append(*s, element)
}

// Pop removes top element
func (s *Stack) Pop() {
	if s.Length() < 1 {
		fmt.Println("Stack is empty. Can't pop")
		return
	}

	fmt.Println("Poping: ", (*s)[len(*s)-1])
	*s = (*s)[:len(*s)-1]
}

// Length returns length of stack
func (s *Stack) Length() int {
	return len(*s)
}

// Print prints elements in stack
func (s *Stack) Print() {
	fmt.Println("Printing stack")

	if s.Length() < 1 {
		fmt.Println("Stack is empty")
		return
	}

	for _, v := range *s {
		fmt.Println(v)
	}
}
