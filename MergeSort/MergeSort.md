## What is a Merge Sort?
An efficient, stable sorting algorithm that uses the **divide-and-conquer** approach to sort elements in an array or a linked list

## What's the complexity?
O(n log n)

## How it works?
The process involves 3 steps, and the process is most commonly implemented using recursion:
- **Divide**: the algorithm repeatedly divides the unsorted list into 2 halves until each sublist contains only 1 element, a list with 1 element is considered sorted by definition.
- **Conquer**: each of the single-element sublist is an inherently sorted array.
- **Combine (Merge)**: the algorithm then repeatedly merges adjacent pairs of sorted sublist to produce new, larger sorted sublists. This merging process compares elements from the  sublist and places the smaller item into a new, combined list in the correct order. This continues until all elements have been merged into a single sorted list.

## Advantages
- Consistent O(n log n) performance across al cases
- Works well with linked list
- Efficient for large datasets and well-suited for external sorting

## Drawbacks
- Requires additinal memory space, which can be a limitation for memory-constrined enviroments.
- Standard implementation is more complex than simpler sorts like bubble sort.
- Slower than some other algorithm (like insertion sort) for very small datasets due to arecursive overhead.
