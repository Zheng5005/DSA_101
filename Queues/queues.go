package main

import "fmt"

type Queue []interface{}

func (q *Queue) Enqueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil // Or handle error
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func main() {
	var myQueue Queue // Create a stack instance

	myQueue.Enqueue(10)
	myQueue.Enqueue("hello")
	myQueue.Enqueue(true)

	fmt.Println("Queue after Enqueues:", myQueue)

	item := myQueue.Dequeue()
	fmt.Println("Dequeued:", item) // Output: Popped: true

	fmt.Println("Queue after Dequeue:", myQueue)

	for !myQueue.IsEmpty() {
		item := myQueue.Dequeue()
		fmt.Println("Dequeued during loop:", item)
	}
	fmt.Println("Is Queue empty?", myQueue.IsEmpty()) // Output: Is stack empty? true
}
