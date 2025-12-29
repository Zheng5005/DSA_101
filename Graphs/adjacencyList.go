package main

import "fmt"

// Is an array/arraylist of linkedList
// Each linkedList has a unique node at the head.
// All adjacent neighbors to thato node are added to that node's linkedList

type GraphL struct {
	adjList map[int][]int
}

func NewGraphL() *GraphL {
	return &GraphL{
		adjList: make(map[int][]int),
	}
}

func (g *GraphL) AddVertex(v int) {
	if _, exists := g.adjList[v]; !exists {
		g.adjList[v] = []int{}
	}
}

func (g *GraphL) AddEdge(i, j int) {
	// Ensure both vertices exists in the map
	g.AddVertex(i)
	g.AddVertex(j)

	// Add the edge (i -> j) and (j -> i) for an undirected graph
	g.adjList[i] = append(g.adjList[i], j)
	g.adjList[j] = append(g.adjList[j], i)
}

func (g *GraphL) PrintGraph() {
	for vertex, neighbors := range g.adjList {
		fmt.Printf("Vertex %d is connected to: %v\n", vertex, neighbors)
	}
}

func main() {
	graph := NewGraphL()

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 5)

	graph.AddVertex(6) // Add an isolated vertex

	graph.PrintGraph()
}
