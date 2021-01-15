package main

import "fmt"

func main() {
	var s Stack
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Push("d")
	s.Print()
	s.Pop()
	s.Print()
}

// Stack structure
type Stack []string

// Push new value to stack
func (s *Stack) Push(element string) {
	*s = append(*s, element)
	fmt.Printf("element %v pushed to stack  \n", element)
}

// Pop removes top element
func (s *Stack) Pop() {
	fmt.Printf("poping element %v from stack \n", (*s)[len(*s)-1])
	*s = (*s)[:len(*s)-1]
}

// Length returns length of stack
func (s *Stack) Length() int {
	return len(*s)
}

// Print prints elements in stack
func (s *Stack) Print() {
	fmt.Println("Printing stack")
	for _, v := range *s {
		fmt.Println(v)
	}
}
