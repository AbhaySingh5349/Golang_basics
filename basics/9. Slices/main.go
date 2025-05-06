package main

// slice is reference to an underline array which holds actual elements
// slice shares storage with array & other slices of that array
// allows dynamic re-sizing & powerfull operations

import (
	"fmt"
	"slices"
)

func main() {
	var nums_1 = []int{1, 2, 3}
	nums_2 := []int{4, 5, 6}

	fmt.Println(nums_1)
	fmt.Println(nums_2)

	nums_slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	nums_slice = a[1:4]
	nums_slice = append(nums_slice, 6, 7)

	fmt.Println("nums_slice: ", nums_slice)

	slice_copy := make([]int, len(nums_slice))
	copy(slice_copy, nums_slice)
	slice_copy[1] = 10

	fmt.Println("slice_copy: ", slice_copy)
	fmt.Println("nums_slice: ", nums_slice)

	if slices.Equal(nums_slice, slice_copy) {
		fmt.Println("equal slices")
	} else {
		fmt.Println("unequal slices")
	}

	slice_2 := nums_slice[2:4]
	fmt.Println("slice_2: ", slice_2)
	fmt.Println("capacity: ", cap(slice_2)) // len of slice + len of array beyond (not capacity of array but capacity of slice it can hold)
	fmt.Println("len: ", len(slice_2))      // count of elements within slice

	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		matrix[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			matrix[i][j] = i * j
		}
	}

	fmt.Println(matrix)
}
