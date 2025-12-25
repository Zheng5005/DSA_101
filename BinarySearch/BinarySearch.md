## What is a Binary Search?
Search algorithm that finds the position of a target value within a **sorted** array, halff of the array is eliminated during each step

## What's the complexity?
O(log n)

## How it works?
1- Set the lower bound (**low**) to the first index and the upper bound (**high**) to the last index of the sorted array, the search space is the entire array
2- Calculate the middle index (**mid**) of the current search space, tipically done by the formula: **mid = (low + high) / 2**
3- Compare the value at the mid index with the target value:
    - **If match is found**: the search ends successfully, and the index is returned
    - **If the target is smaller**: the target must be in the left half of the array. The algorithm eliminates the right half by setting the **high** index to **mid - 1**
    - **If the target is larger**: the target must be in the right half of the array. The algorithm eliminates the left half by setting the **low** index to **mid + 1**
4- Steps 2 and 3 are repeted until the targert is found or the search space is exhausted (which happens when the **low** index becomes greater than the **high** index)

## When to use Binary Search
- Approximate matching
- Finding the next-smallest or next largest element

## Advantages
- Speed and Efficiency
- Halves the search space
- Effective for larger datasets

## Drawbacks
- Requires sorted datasets
- Ineficient for dynamic data
- Works efficiently with data structures that allow for immediate random access to any element, but is inefficient for structures like linked list
