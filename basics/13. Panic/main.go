package main

import "fmt"

// stops current execution, unwinds stack & executes any defered funcs

func process(val int) {
	defer fmt.Println("defered 1")
	defer fmt.Println("defered 2")

	if val < 0 {
		fmt.Println("before panic")

		panic("value must be non-negative")

		// fmt.Println("after panic")
		// defer fmt.Println("defered 2")
	}
	fmt.Println("processing input: ", val)
}

func main() {
	process(10)
	process(-1)
}
