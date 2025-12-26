## What is a Insertion Sort?
A simple and stable sorting algorithm that works bu building a sorted list one element at a time.

## What's the complexity?
O(n^2)

## How it works?
- The first element is considered the initial sorted protion of the array.
- It iterates through the remaining unsorted elements, taking one element (the "key") at a time.
- The key is compared with elements in the sorted portion, moving backward from right to left.
- If a sorted element is greater than the key, it is shifted one position to the right to make space.
- This shifting continues until a smaller element is found, or the beginning of the array is reached.
- The key is then inserted into the correct position
- This process repeats until all elements from the unsorted part have been placed into the growing sorted part

## When to use Insertion Search
- Sorting small lists where overhead of more complex algorithms is nor worth it.
- Sorting list that are already mostly sorted

## Advantages
- Simple to implement
- Stable
- Adaptative (efficient for nearly sorted data)

## Drawbacks
- Inefficient for large datasets
