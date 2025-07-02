package main

import "fmt"

type Node struct {
	Data int
	Next *Node // Pointer to the next Node
}

type LinkedList struct {
	Head *Node // Pointer to the fist node of the LinkedList
	Tail *Node // Pointer to the last node of the LinkedList
	length int // Number of nodes
}

func (l *LinkedList) Add(data int)  {
	newNode := &Node{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.length++
}

func (l *LinkedList) Print()  {
	current := l.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func main()  {
	myList := &LinkedList{}
	myList.Add(10)
	myList.Add(20)
	myList.Add(30)
	myList.Add(40)
	myList.Add(50)
	myList.Print()

	myList.Exercise1_rotate(4)
	fmt.Println("Rotated left by 4: ")
	myList.Print()
}

func (l *LinkedList) Exercise1_rotate(k int)  {
	// Given a singly linked list and an integer k, the task is to rotate the linked list to the left by k places.
	if l.Head == nil || l.Head.Next == nil || k == 0 {
		return
	}

	// Counting total nodes
	current := l.Head
	count := 1
	for current.Next != nil {
		current = current.Next
		count++
	}

	// If k is greater than or equal to list length, use modulo
	k = k % count
	if k == 0 {
		return
	}

	// Now current points to the last node
	current.Next = l.Head // this makes the LinkedList circular

	// Traverse to (k)th node
	temp := l.Head
	for i := 1; i < k; i++ {
		temp = temp.Next
	}

	// Setting new head and break the circle
	l.Head = temp.Next
	temp.Next = nil
}
