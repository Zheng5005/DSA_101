package main

import "fmt"

type Graph struct {
	vertices map[int][]int
}

func (g *Graph) AddEdge(source, destination int) {
	if g.vertices == nil {
		g.vertices = make(map[int][]int)
	}

	g.vertices[source] = append(g.vertices[source], destination)
}

func (g *Graph) DFS(startVertex int) {
	visited := make(map[int]bool)

	g.DFSUtil(startVertex, visited)
	fmt.Println()
}

func (g *Graph) DFSUtil(v int, visited map[int]bool) {
	visited[v] = true
	fmt.Printf("%d ", v)

	for _, neighbor := range g.vertices[v] {
		if !visited[neighbor] {
			g.DFSUtil(neighbor, visited)
		}
	}
}

// For cases where you might be concerned about the call stack depth in very deep graphs
func (g *Graph) DFSIterative(startVertex int) {
	visited :=  make(map[int]bool)
	stack := []int{startVertex} // use a slice as a stack

	for len(stack) > 0 {
		// Pop the top element from the stack (LIFO)
		current := stack[len(stack) - 1]
		stack := stack[:len(stack) - 1]

		if !visited[current] {
			visited[current] = true
		  fmt.Printf("%d ", current)

			// Push unvisited neighbors onto the stack
			for _, neighbor := range g.vertices[current] {
				if !visited[neighbor] {
					stack = append(stack, neighbor)
				}
			}
		}
	}

	fmt.Println()
}

func main() {
	g := &Graph{}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(2, 3)
	g.AddEdge(3, 3)

	fmt.Println("Depth First Search starting from vertex 2:")
	g.DFS(2)
}
