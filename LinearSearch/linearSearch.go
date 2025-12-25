package main

import "fmt"

func LinearSearch(arr []int, target int) int  {
	for i, element := range arr {
		if element == target {
			return i // Return the index if the target is found
	 	}
	}

	return -1 // Return -1 if the target is not found
}

func main()  {
	data := []int{10, 50, 30, 70, 80, 20}
	target := 30
	resultIndex := LinearSearch(data, target)

	if resultIndex != -1 {
		fmt.Printf("Element found at index: %d\n", resultIndex)
  } else {
		fmt.Println("Element not found")
	}
}
