package main

import "fmt"

func main()  {
	dynamicArray := []int{}
	slice := make([]int, 0, 3)

	dynamicArray = append(dynamicArray, 10)
	dynamicArray = append(dynamicArray, 20)
	dynamicArray = append(dynamicArray, 30)

	slice = append(slice, 10)
	slice = append(slice, 20)
	slice = append(slice, 30)

	fmt.Printf("Length: %d, Capacity: %d, Value: %v\n", len(dynamicArray), cap(dynamicArray), dynamicArray)
	fmt.Printf("--------------------------------------------------------------\n")
	fmt.Printf("Length: %d, Capacity: %d, Value: %v\n", len(slice), cap(slice), slice)
}
