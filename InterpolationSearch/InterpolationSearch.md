## What is a Interpolation Search?
Is an improvement over binary search best used for "uniformly" distributed data, this search "guesses" where a value might be based on calculated probe results, if probe is incorrect, search area is narrowed, and a new probe is calculated

## What's the complexity?
O(n) -> worst case (non-uniform data)

## How it works?
The algorithm refines the search area by estimating the target's location based on the values at the current bounds (low and high). The estimated position (**pos**) is calculated using the following formula:

$$
pos = low + \frac{(high - low) * (target - arr[low])}{arr[high] - arr[low]}
$$

where **low** and **high** are the bounds, and **arr** is the array

The process involves initializing **low** and **high**, then repeatedly calculating **pos** and compoaring **arr[pos]** to the **target**. The search range is adjusted based on the comparison, continuing until the target is found or the range is empty.

## When to use Interpolation Search
- **Sorted Data**: the input array must be in ascending order
- **Uniform Distribution**: the algorithm performs best when the elements are uniformly distributed (evenly spaced). If the data is clustered, its performance can degrade to O(n), similar to a linear search
- **Large Datasets**: it offers lookup times than binary search for large, uniformly distributed datasets

## Advantages
- Faster on uniform data
- Good for time-sensitive apps
- Can quickly determine if a target isn't present, leading to faster retrieving data quickly in time-sensitivescenarios

## Drawbacks
- Requires sorted data
- Performance heavily relies on data distribution, it's inefficient if data isn't uniform
- More complex than binary search
