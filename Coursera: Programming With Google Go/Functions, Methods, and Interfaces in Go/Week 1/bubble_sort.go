package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	numsSlice := make([]int, 0, 3)

	// Prompt user for input
	fmt.Println("Enter integers separated by spaces, type 'X' to finish:")

	// Read input from user
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Check if the user wants to exit
		if strings.ToUpper(input) == "X" {
			break
		}

		// Convert input to integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to finish.")
			continue
		}

		// Append number to slice
		numsSlice = append(numsSlice, num)
	}

	// Perform Bubble Sort
	BubbleSort(numsSlice)

	// Print sorted integers
	fmt.Println("Sorted integers:", numsSlice)
}

// BubbleSort sorts a slice of integers using the bubble sort algorithm
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

// Swap swaps the position of two adjacent elements in a slice
func Swap(nums []int, i int) {
	nums[i], nums[i+1] = nums[i+1], nums[i]
}
