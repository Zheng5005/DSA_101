package main

import "fmt"

// Is a 2D array to store 1's and 0's to represent egdes
// # of rows = # of unique nodes
// # of columns = # of unique nodes

type Graph struct {
	matrix [][]int
	numVertices int
}

func NewGraph(numVertices int) *Graph {
	g := &Graph{
		numVertices: numVertices,
		matrix: make([][]int, numVertices),
	}

	for i := range g.matrix {
		g.matrix[i] = make([]int, numVertices)
	}

	return g
}

func (g *Graph) AddEdge(i, j int) {
	if i >=0 && i < g.numVertices && j >= 0 && j < g.numVertices {
		g.matrix[i][j] = 1
		g.matrix[j][i] = 1 // For undirected Graph
	}
}

func (g *Graph) HasEdge(i, j int) bool {
	if i >=0 && i < g.numVertices && j >= 0 && j < g.numVertices {
		return g.matrix[i][j] == 1
	}

	return false
}

func main() {
	g := NewGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	fmt.Println("Adjacency Matrix:")
	for i := 0; i < g.numVertices; i++ {
		fmt.Println(g.matrix[i])
	}

	fmt.Printf("Edge (0, 3) exist: %v\n", g.HasEdge(0, 3))
	fmt.Printf("Edge (1, 2) exist: %v\n", g.HasEdge(1, 2))
}
