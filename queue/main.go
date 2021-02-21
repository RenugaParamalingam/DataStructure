package main

import "fmt"

func main() {
	var q Queue
	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")
	q.Print()

	q.Dequeue()
	q.Print()

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}

// Queue structure
type Queue []string

// Enqueue inserts element to the queue
func (q *Queue) Enqueue(element string) {
	*q = append(*q, element)
	fmt.Println("Enqueued: ", element)
}

// Dequeue removes element in front of queue
func (q *Queue) Dequeue() {
	if q.Length() < 1 {
		fmt.Println("Queue is empty. Can't deuqueue.")

		return
	}

	fmt.Println("Dequeued: ", (*q)[0])
	(*q) = (*q)[1:]
}

// Print prints all elements in queue
func (q *Queue) Print() {
	fmt.Println("Elements in queue")

	for _, v := range *q {
		fmt.Println(v)
	}
}

// Length returns length of queue
func (q *Queue) Length() int {
	return len(*q)
}
