## What is a Selection Sort?
Is an in-place comparison sorting algorithm that works by repeatedly finding the minimum element from the unsorted portion of a list and swapping it with the first unsorted element

## What's the complexity?
O(n^2)

## How it works?
- **Initialization**: the sorted sublist is initially empty, and the entire array is the unsorted sublist
- **Find Minimum**: Iterate through the unsorted part of the array to find the smallest element
- **Swap**: Swap the found minimum element with the first element of the unsorted section
- **Advance Boundary**: Move the boundary between the sorted and unsorted parts one element to the right. The first element is now considered sorted.
- **Repeat**: Repeat steps 2-4 for the remaining unsorted portion until the entire array is sorted.

## When to use Selection Search
- Small datasets
- Memory-constrained environments

## Advantages
- Simplicity and Ease of Implementation
- Predictable Performance
- Efficient for small datasets

## Drawbacks
- Inefficient for large datasets
- Not adaptive for sorted datasets
- High number of comparisons
