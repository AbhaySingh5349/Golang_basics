package main

import (
	"fmt"
	"time"
)

// buffereing let's channel to hold limited num of values before "blocking sender"
// buffered channels are useful for managing data flow & concurrency
// they do not require immediate receiver
// buffered channels allow "async communication" i.e it allows "senders" to continue working without blocking untill buffer is full
// with help of Load Balancers, they manage data transfer between producer & consumer

// func main() {
// 	// =========== BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("End of program.")
// }

func main() {
	// ================== BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
	// make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("Goroutine 2 second timer started.")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch) // ends <- starts (as soon as "<-ch" is encountered, "ch <- 3" executes & print statement might not complete)
	}()
	// fmt.Println("Blocking starts")
	ch <- 3 // Blocks because the buffer is full
	// fmt.Println("Blocking ends")
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
}
