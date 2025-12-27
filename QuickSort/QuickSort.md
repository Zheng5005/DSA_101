## What is a Quick Sort?
Is a highly efficient sorting algorithm that uses a **divide-and-conquer** strategy. It works bu selecting a "pivot" element fromt the array and partitioning the other elements into 2 sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively

## What's the complexity?
O(n^2) -> Worst case

## How it works?
The process can be summarized in 3 main steps:
- **Pick a Pivot**: Select an element from the array to serve as the pivot. The choice of pivot can vary, with common methods including picking the first, last, a random element, or the median-of-three to improve performance.
- **Partition the Array**: Rearrange the array so that all elements smaller than the pivot are moved to its left, and all elements greater than the pivot are moved to its right. After this step, the pivot is in its final sorted position.
- **Recursively Sort sub-arrays**: Apply the quicksort algorithm recursively to the 2 sub-arrays (left and right of the pivot) until each sub-array contains only 1 or 0 elements, at which point they are already sorted.

## Advantages
- Good cache locality, which makes it faster than algorithms like merge sort in many practical implementations
- In-place sorting, so it is memory efficient and a good choice when space is limited
- Fast in average-case scenarios (O(n log n))

## Drawbacks
- Worst case performance is O(n^2)
- Difficult to implement efficiently
- Unstable sort, meaning the relative order of equal elements may not be preserved
