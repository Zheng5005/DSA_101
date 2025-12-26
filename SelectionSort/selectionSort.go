package main

import "fmt"

func SelectionSort(arr []int) []int  {
	if len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

func main()  {
	numbers := []int{64, 25, 12, 22, 11}
  fmt.Println("Before sorting:", numbers)
  SelectionSort(numbers)
  fmt.Println("After sorting:", numbers)
}
