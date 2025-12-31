package main

import "fmt"

// Node represent a single element in the tree
type Node struct {
	Value int
	Left *Node
	Right *Node
}

// BTS represent the entire binary search tree
type BST struct {
	Root *Node
}

// Insert adds a new value to the BST
func (bst *BST) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &Node{Value: value}
		return
	}

	bst.Root.Insert(value)
}

// Insert recursively adds a value to a node's subtree
func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search finds a value in the BST
func (bst *BST) Search(value int) bool {
	return bst.Root.Search(value)
}

// Search recursively looks for a value from a given node
func (n *Node) Search(value int) bool {
	if n == nil {
		return false //Not found
	}

	if value == n.Value {
		return true
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// InOrderTraversal performs in-order traversal of the BST
func (bst *BST) InOrderTraversal() {
	bst.Root.inOrder()
	fmt.Println()
}

// inOrder recursively traverses nodes in order
func (n *Node) inOrder() {
	if n == nil {
		return
	}

	n.Left.inOrder()
	fmt.Printf("%d ", n.Value)
	n.Right.inOrder()
}

func main() {
	tree := BST{}
	tree.Insert(10)
	tree.Insert(5)
  tree.Insert(20)
  tree.Insert(3)
  tree.Insert(7)
  tree.Insert(18)
	tree.Insert(25)

	fmt.Print("In-order traversal: ")
	tree.InOrderTraversal()

	fmt.Printf("Search for 18: %t\n", tree.Root.Search(18))
	fmt.Printf("Search for 99: %t\n", tree.Root.Search(99))
}
