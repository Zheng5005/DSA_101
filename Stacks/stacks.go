package main

import "fmt"


type Stack []interface{} //this is a generic stack

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
  }
  index := len(*s) - 1
  element := (*s)[index]
  *s = (*s)[:index]
  return element, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
  }
	return (*s)[len(*s)-1], true
}

func main() {
    var myStack Stack // Create a stack instance

    myStack.Push(10)
    myStack.Push("hello")
    myStack.Push(true)

    fmt.Println("Stack after pushes:", myStack)

    item, ok := myStack.Pop()
    if ok {
        fmt.Println("Popped:", item) // Output: Popped: true
    }

    fmt.Println("Stack after pop:", myStack)

    topItem, ok := myStack.Peek()
    if ok {
        fmt.Println("Top item:", topItem) // Output: Top item: hello
    }

    for !myStack.IsEmpty() {
        item, _ := myStack.Pop()
        fmt.Println("Popped during loop:", item)
    }
    fmt.Println("Is stack empty?", myStack.IsEmpty()) // Output: Is stack empty? true
}
