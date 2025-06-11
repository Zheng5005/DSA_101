package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // The value of the item
  priority int    // The priority of the item
  index    int    // The index of the item in the heap, needed for updates
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// Less defines the priority order. For a min-priority queue, it's pq[i].priority < pq[j].priority.
// For a max-priority queue, it's pq[i].priority > pq[j].priority.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority // Example: Min-priority queue
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
  item := x.(*Item)
  item.index = n
  *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
  n := len(old)
  item := old[n-1]
  old[n-1] = nil // Avoid memory leak
	item.index = -1 // For safety
  *pq = old[0 : n-1]
  return item
}

func mainPriorityQueue()  {
	pq := make(PriorityQueue, 0)
  heap.Init(&pq)

  // Push items
  heap.Push(&pq, &Item{value: "task A", priority: 3})
  heap.Push(&pq, &Item{value: "task B", priority: 1})
  heap.Push(&pq, &Item{value: "task C", priority: 2})

  // Pop items (will be in priority order)
  for pq.Len() > 0 {
  	item := heap.Pop(&pq).(*Item)
		fmt.Print(item)
	}
}
