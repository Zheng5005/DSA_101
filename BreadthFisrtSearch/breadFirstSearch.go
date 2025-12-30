package main

import "fmt"

// Graph represents an undirected graph using an adjacency list
type Graph struct {
    adjList map[int][]int
}

// NewGraph creates a new graph instance
func NewGraph() *Graph {
    return &Graph{adjList: make(map[int][]int)}
}

// AddEdge adds an edge to the graph (undirected)
func (g *Graph) AddEdge(u, v int) {
    g.adjList[u] = append(g.adjList[u], v)
    g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) BFS(startNode int) {
	// A queue for managing nodes to visit (FIFO principle)
	queue := []int{startNode}

	// A map to keep track of visited nodes to prevent cycles
	visited := make(map[int]bool)
	visited[startNode] = true

	// Loop until the queue is empty
	for len(queue) > 0 {
		// Dequeue the front node (FIFO)
		currNode := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", currNode) // Process the Node

		// Enqueue all unvisited neighbors
		for _, neighbor := range g.adjList[currNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
  graph.AddEdge(1, 2)
  graph.AddEdge(3, 2)
  graph.AddEdge(3, 4)
  graph.AddEdge(4, 5)

	fmt.Println("Breadth First Search traversal starting from node 0:")
	graph.BFS(0)
	fmt.Println()
}
