package main

import (
	"errors"
	"fmt"
)

// parameters passed are copied values & any modification to them does not change original values

func add(a, b int) int {
	return a + b
}

func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func divide(a, b int) (int, int) {
	quotient, remainder := a/b, a%b
	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a > b", nil
	} else if b > a {
		return "b > a", nil
	}
	return "", errors.New("Invalid comparison !")
}

func variadicSum(statement string, nums ...int) (string, int) {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return statement, sum
}

func main() {
	fmt.Println("3+5 = ", add(1, 2))

	greet := func() {
		fmt.Println("Hello anonymous func")
	}

	greet()

	// functions as variable types (1st class objects / citizens)
	operation := add
	fmt.Println("3+5 = ", operation(1, 2))

	fmt.Println("3+5 = ", applyOperation(3, 5, add))

	multiplier := createMultiplier(5)
	fmt.Println("3*5 = ", multiplier(3))

	quotient, remainder := divide(10, 3)
	fmt.Printf("quotient: %v , remainder: %v\n", quotient, remainder)

	res, err := compare(1, 1)
	if err != nil {
		fmt.Println("comparison err: ", err)
	} else {
		fmt.Println("comparison res: ", res)
	}

	statement1, sum1 := variadicSum("sum of 1 to 4:", 1, 2, 3, 4)
	fmt.Println(statement1, sum1)

	numbers := []int{1, 2, 3, 4}
	statement2, sum2 := variadicSum("sum of 1 to 4:", numbers...)
	fmt.Println(statement2, sum2)
}
