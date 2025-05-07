package main

import "fmt"

// closure is a func value that references variables outside it's body
// these variables persist as long as closure is referenced
// closures work with Lexical Scoping i.e they capture variables from surrounding context where they are defined
// this allows closures to access variables even after outer function has finished execution

func adder() func() int {
	i := 0
	fmt.Println("initial value: ", i)

	// anonymous func which increments val by 1 each time it's called
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}

func main() {
	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			fmt.Printf("subtracted %d from countdown: ", x)
			return countdown
		}
	}() // anonymous func which is immediately executed

	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))
	fmt.Println(subtractor(5))
}
