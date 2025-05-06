package main

import "fmt"

// recover is built-in func used only inside deferd functions to regain control of panicking go-routine

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered: ", r)
		}
	}()

	fmt.Println("Start process")
	panic("something went wrong !! panic...")
}

func main() {
	process()
	fmt.Println("returned from main")
}
