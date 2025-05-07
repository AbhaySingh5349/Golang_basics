package main

import "fmt"

func modifyValue(ptr *int) {
	*ptr++
}

func main() {
	var a int = 10

	var ptr *int
	ptr = &a

	fmt.Println("a: ", a)
	fmt.Println("ptr: ", ptr)
	fmt.Println("*ptr: ", *ptr) // de-referencing a pointer

	modifyValue(ptr)
	fmt.Println("modified a: ", a)
}
