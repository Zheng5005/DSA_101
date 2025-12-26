package main

import "fmt"

func insetionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j - 1] > arr[j]; j-- {
			arr[j], arr[j - 1] = arr[j - 1], arr[j]
		}
	}

	return arr
}

func main()  {
	numbers := []int{64, 25, 12, 22, 11}
  fmt.Println("Before sorting:", numbers)
  insetionSort(numbers)
  fmt.Println("After sorting:", numbers)
}
