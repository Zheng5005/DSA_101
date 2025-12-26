package main

import "fmt"

func bubbleSort(arr []int)  {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr) - i; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func bubbleSortOptimal(arr []int)  {
	n := len(arr)
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
	}
}

func main()  {
	nums := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", nums)
  bubbleSort(nums)
  fmt.Println("Sorted array:", nums)
}
