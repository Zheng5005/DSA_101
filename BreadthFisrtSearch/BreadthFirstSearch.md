## What is Breadth First Search?
A graph traversal technique that systematically explores a graph or tree layer by layer starting from a given source node, it uses a **queue** to manage which nodes to visit next, ensuring all nodes at the current depth are explores before moving to the next level.

## What's the complexity?
O(V + E), where V is the number of vertices and E the number of edges

## How it works?
It operates by "expanding" the search outward from the starting point like a fire spreading in concentric rings
- **Initialize** a queue and a way to track visited nodes
- **Enqueue** the starting (source) node into the queue and mark it as visited
- **Loop** as long as the queue is not empy:
    - **Dequeue** a node from the front of the queue
    - **Process** the node
    - **Enqueue all unvisited neighbors** of the processed node, marking each new neighbor as visited as it's added to the queue.
- The process continues until all reachable nodes have been visited ot the target node is found

## Advantages
- Guaranteed finding the shortest path in a unweighted graph
- It will find a solution if one exist

## Drawbacks
- High memory usage
- Can be slow and exhaustive in large or deep graphs where the solution is far away
