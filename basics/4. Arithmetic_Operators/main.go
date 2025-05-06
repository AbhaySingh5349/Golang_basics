package main

import "fmt"

func main() {
	var a, b int = 10, 3
	var res int

	res = a + b
	fmt.Println("Addition: ", res)

	res = a - b
	fmt.Println("Subtraction: ", res)

	res = a * b
	fmt.Println("Multiplication: ", res)

	res = a / b
	fmt.Println("Division: ", res)

	res = a % b
	fmt.Println("Modulo: ", res)

	const pi float64 = 22 / 7.0
	fmt.Println("Division: ", pi)
}
