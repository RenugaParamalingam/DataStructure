package main

import "fmt"

func main() {
	s := NewSinglyLinkedList()
	s.Push(1)
	s.Push(3)
	s.Push(4)

	s.Print()
	fmt.Println("node at index 1: ", s.GetByIndex(1))

	fmt.Println("push at index 1 with data 2")
	s.PushAtIndex(2, 1)
	s.Print()
	fmt.Println("length: ", s.Length())

	fmt.Println("\n after reorder")
	s.ReorderList()
	s.Print()

	s.DeleteAtIndex(1)
	fmt.Println("After deleting at index ", 1)
	s.Print()

	s.Reverse()
	fmt.Println("After reverse ")
	s.Print()

	fmt.Println("hasCycleUsingPointer: ", s.HasCycleUsingPointer())
	fmt.Println("hasCycleUsingMap: ", s.HasCycleUsingMap())

	node2 := s.GetByIndex(2)
	node2.next = s.Head.next
	fmt.Println("Cycle added")
	// s.Print()

	fmt.Println("hasCycleUsingPointer: ", s.HasCycleUsingPointer())

	fmt.Println("hasCycleUsingMap: ", s.HasCycleUsingMap())
}

// Node holds structure for a node in singly linked list
type Node struct {
	value int
	next  *Node
}

// SinglyLinkedList holds structure of singly linked list
type SinglyLinkedList struct {
	Head *Node
}

// NewSinglyLinkedList initialize a singly linked list
func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

// Push add element to the singly linked list
func (s *SinglyLinkedList) Push(data int) {
	if s.Head == nil {
		s.Head = &Node{value: data, next: nil}
		return
	}

	current := s.Head

	for current.next != nil {
		current = current.next
	}

	current.next = &Node{value: data, next: nil}
}

// PushAtIndex push data at the given index
func (s *SinglyLinkedList) PushAtIndex(data, index int) {
	head := s.Head

	if index == 0 {
		s.Head = &Node{value: data, next: nil}
		if head.next != nil {
			s.Head.next = head.next
		}
		return
	}

	count := 0
	prev := s.Head
	reqIndex := s.Head

	for node := s.Head; node != nil; node = node.next {
		if count == index {
			reqIndex = node
			break
		}

		prev = node

		count++
	}

	node := &Node{
		value: data,
	}

	prev.next = node
	node.next = reqIndex
}

// DeleteAtIndex deletes node of the given index
func (s *SinglyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if s.Head.next != nil {
			s.Head = s.Head.next

			return
		}
		s.Head = nil

		return
	}

	prev := s.Head
	count := 0
	for node := s.Head; node != nil; node = node.next {
		if count == index {
			prev.next = node.next
			break
		}

		prev = node
		count++
	}
}

// GetByIndex returns node for the index
func (s *SinglyLinkedList) GetByIndex(index int) *Node {
	count := 0

	for node := s.Head; node != nil; node = node.next {
		if count == index {
			return node
		}

		count++
	}

	return nil
}

// Print prints all the elements in the singly linked list
func (s *SinglyLinkedList) Print() {
	if s.Head == nil {
		fmt.Println("List is empty")
	}

	current := s.Head
	for current.next != nil {
		fmt.Print(current.value, " -> ")
		current = current.next

		// Print tail
		if current.next == nil {
			fmt.Println(current.value)
		}
	}
}

// Length returns length of the singly linked list
func (s *SinglyLinkedList) Length() int {
	count := 0

	for node := s.Head; node != nil; node = node.next {
		count++
	}

	return count
}

// Reverse reverse singly linked list and print it
func (s *SinglyLinkedList) Reverse() {
	var prev *Node
	node := s.Head

	for node != nil {
		temp := node.next
		node.next = prev

		prev = node
		node = temp
	}
	s.Head = prev

}

// ReorderList reorder list in the following pattern
// Given a singly linked list L: L0→L1→…→Ln-1→Ln,
// reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
func (s *SinglyLinkedList) ReorderList() {
	l := s.Length()

	for i := 1; i <= l/2-1; i++ {
		prev := s.GetByIndex(i)
		nex := prev.next
		reorderNodePrev := s.GetByIndex(l - 1 - i)
		reorderNode := reorderNodePrev.next
		fmt.Println(prev, nex, reorderNodePrev, reorderNode)
		reorderNodePrev.next = nil

		prev.next = reorderNode
		reorderNode.next = nex

	}
}

// HasCycleUsingPointer detects and return true if cycle exists using pointers
func (s *SinglyLinkedList) HasCycleUsingPointer() bool {
	slow := s.Head
	fast := s.Head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

// HasCycleUsingMap detects and return true if cycle exists using map
func (s *SinglyLinkedList) HasCycleUsingMap() bool {
	visited := make(map[*Node]bool, 0)

	node := s.Head
	for node.next != nil {
		if _, ok := visited[node]; ok {
			return true
		}

		visited[node] = true
		node = node.next
	}

	return false
}
