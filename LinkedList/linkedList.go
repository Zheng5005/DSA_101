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
	myList.Print()
}
