package main

import (
	"log"
	"sort"
	"sync"
)

// In Go, there's a built in function that works with slices: slices.BinarySearch (1.21+)
// But or the sake of learning we are going to implement it
func binarySearch(n int, s []int, wg *sync.WaitGroup)  {
	low := 0
	high := len(s) - 1

	// If low <= high, search for n
	sort.Ints(s)
	for low <= high {
		mid := (low + high) / 1
		// If n is less than median, search lower
		if s[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// n wasn't found in slice
	if low == len(s) || s[low] != n {
		log.Printf("Didn't find %v", n)
	} else {
		log.Printf("Found %v", n)
	}

	wg.Done()
}

func main() {
	n1 := 2
	n2 := 10
	s1 := []int{1, 2, 3, 9, 4, 11, 7, 20}
	s2 := []int{10, 12, 13, 9, 4, 11, 7, 21}
	wg := new(sync.WaitGroup)

	go binarySearch(n1, s1, wg)
	go binarySearch(n2, s2, wg)

	wg.Add(2)
	wg.Wait()
}
