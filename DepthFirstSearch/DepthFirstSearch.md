## What is Depth First Search?
A search algorithm for traversing a tree or graph data structure, the algorithm starts at an arbitrary root node and explores as far as possible along each branch before backtracking. This approach uses a **Stack** to manage which node to visit first.

## What's the complexity?
O(V + E), where V is the number of vertices and E the number of edges

## How it works?
The core idea is to go deep into the graph, following a single path until a dead end (a node with no unvisited neighbors) is reached. At that point, it backtracks to the most recent unvisited junction and explores an alternative path.

The general step are:
- **Start** at the source vertex and mark it as visited
- **Push** the starting vertex onto a stack
- **Pop** a vertex from the stack to be the current node
- **Explore** all it's unvisited adjacent neighbors. For each unvisited neighbor, mark it as visited and push it onto the stack
- **Repeat** steps 3 and 4 until the stack is empty

## Advantages
- Memory efficient, it uses less memory than Breadth Fisrt Search (BFS)
- Simple implementation
- Good for deep/large graphs

## Drawbacks
- Doesn't guarantee finding the shortest path to a goal, unlike BFS
- Can be slow in worst-case scenarios
- Can get trapped in very deep or infinite branches/cycles in graphs
