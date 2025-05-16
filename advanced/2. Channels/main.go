package main

import (
	"fmt"
)

// channels are way to communicate within goroutines, we cannot make channel work in our "any func without goroutine"
// unlike goroutines, "channels" are blocking in nature

func main() {

	//variable := make(chan type) '<-' receive & send operator

	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		greeting <- greetString // blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// receiving from channel is blocking in nature: if no data is available, receiver waits until data is sent into channel
	// receiver is part of "main goroutine", so "greeting channel" is communicating between "main & greeting goroutine"
	receiver := <-greeting // receiver is receiving data from "greeting" channel
	fmt.Println("receiver: ", receiver)

	receiver = <-greeting
	fmt.Println("receiver: ", receiver)

	for range 5 {
		receiver = <-greeting
		fmt.Println("receiver: ", receiver)
	}

	fmt.Println("All values in main goroutine received.")

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println("additional thread receiver: ", receiver)
	// 	receiver = <-greeting
	// 	fmt.Println("additional thread receiver: ", receiver)
	// }()

	// time.Sleep(1 * time.Second) // we need to make "main goroutine" wait if receiver is not on "main goroutine"

	fmt.Println("End of program.")
}
