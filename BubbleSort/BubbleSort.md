## What is a Bubble Sort?
An algorithm that repeatedly steps through a list, compares adjacent elements, and swaps them if they're in the wrong order, causing larger elements to "bubble up" to their correct position.

## What's the complexity?
O(n^2) -> Worst case

## How it works?
- **Compare & Swaps**: the algorithm iterates through the list, comparing th first elements with the second, then the second with the third and so on.
- **Move Largest**: if the first element is larger (for ascending sort), they are swapped. This process ensures the largest element "bubbles" to the end of the list in the first pass.
- **Repeat**: the passes continue, with each pass placing the next largest unsorted element in it's correct spot, reducing the comparison range
- **Termination**: the list sorted when a full pass occurs with no swaps

## Advantages
- Simple to understand & implement
- Stable
- Minimal space complexity (O(1) for swaps)

## Drawbacks
- Poor performance for most real-world scenarios
- Limited real-world use
