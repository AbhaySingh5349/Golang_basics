package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Enter 1st floating point number:")
	fmt.Scanln(&input)

	// Parse the input as a float64
	num1, err1 := strconv.ParseFloat(input, 64)
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}

	// Truncate the floating point number
	truncated_num1 := int(math.Trunc(num1))

	// Print the truncated integer
	fmt.Println("1st Truncated integer:", truncated_num1)

	fmt.Println("Enter 2nd floating point number:")
	fmt.Scanln(&input)

	// Parse the input as a float64
	num2, err2 := strconv.ParseFloat(input, 64)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	// Truncate the floating point number
	truncated_num2 := int(math.Trunc(num2))

	// Print the truncated integer
	fmt.Println("2nd Truncated integer:", truncated_num2)
}
