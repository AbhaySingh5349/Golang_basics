package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println("numbers: ", numbers)

	numbers[0] = 5
	numbers[4] = 10
	fmt.Println("numbers: ", numbers)

	fruits := [4]string{"apple", "banana", "orange", "grapes"}
	fmt.Println("fruits: ", fruits)

	for idx, val := range fruits {
		fmt.Printf("idx: %d, val: %s\n", idx, val)
	}
	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(matrix)

	originalArr := [3]int{1, 2, 3}
	deepCopiedArr := originalArr // deep copy

	deepCopiedArr[0] = 4

	fmt.Println("originalArr: ", originalArr)
	fmt.Println("deepCopiedArr: ", deepCopiedArr)

	var shallowCopiedArr *[3]int // shallow copy
	shallowCopiedArr = &originalArr

	shallowCopiedArr[0] = 5

	fmt.Println("originalArr: ", originalArr)
	fmt.Println("shallowCopiedArr: ", shallowCopiedArr)
}
