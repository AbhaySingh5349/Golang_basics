package main

import "fmt"

func forLoop() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	numbers := []int{1, 2, 3, 4, 5}
	for idx, val := range numbers {
		fmt.Printf("idx: %d, val: %d \n", idx, val)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even: ", i)
		} else {
			fmt.Println("odd: ", i)
		}
	}

	for i := range 10 {
		fmt.Print(i)
	}
}

func whileLoop() {
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}
}

func main() {
	// forLoop()
	whileLoop()
}
