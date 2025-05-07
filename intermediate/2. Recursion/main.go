package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n/10 == 0 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func main() {
	fmt.Println("factorial(4): ", factorial(4))
	fmt.Println("sumOfDigits(1234): ", sumOfDigits(1234))
}
