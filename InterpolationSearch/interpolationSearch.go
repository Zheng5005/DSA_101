package main

import "fmt"

// InterpolationSearch finds the index of a target value in a sorted, uniformly distributed integer array.
// It returns the index if found, otherwise -1.
func interpolationSearch(arr []int, query int) int {
	// Find indexes of two corners
	lo := 0
	hi := len(arr) - 1

	// Since array is sorted, an element present in array must be in range defined by corner
	for lo <= hi && query >= arr[lo] && query <= arr[hi] {
		// Probing the position with keeping uniform distribution in mind.
		midIndex := lo + (query-arr[lo])*(hi-lo)/(arr[hi]-arr[lo])
		midItem := arr[midIndex]

		// Target is found
		if midItem == query {
			return midIndex
		} else if midItem < query {
			// Target is in upper part
			lo = midIndex + 1
		} else if midItem > query {
			// Target is in lower part
			hi = midIndex - 1
		}
	}

	return -1
}
func main()  {
	// Example array (must be sorted and ideally uniformly distributed)
  data := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
  target := 60

	index := interpolationSearch(data, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
