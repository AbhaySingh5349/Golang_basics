package main

import "fmt"

// defer is mechanism that allows to postpone execution of function until surrounding function returns
// useful feature to ensure cleanup actions or finalizing tasks are performed regardless of how function exits (normally or panic)

// defer func is always part of other func i.e a surrounding func that encloses defer func
// defer statements are evaluated as soon as they are encountered

func process(val int) {
	defer fmt.Println("deferred val statement executed: ", val)
	defer fmt.Println("1st deferred statement executed")
	defer fmt.Println("2nd deferred statement executed")
	defer fmt.Println("3rd deferred statement executed")

	val++
	fmt.Println("normal statement executed")
	fmt.Println("normal val statement executed: ", val)
}

func main() {
	process(0)
}
