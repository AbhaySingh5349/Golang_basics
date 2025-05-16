package main

import (
	"fmt"
	"time"
)

// by default channels are unbuffered i.e no storage
// unbuffered channels need immediate receiver, so we send values in separate goroutine & keep receiver in different goroutine

// buffereing let's channel to hold limited num of values before "blocking sender"
// buffered channels are useful for managing data flow & concurrency
// buffered channels allow "async communication" i.e it allows "senders" to continue working without blocking untill buffer is full
// with help of Load Balancers, they manage data transfer between producer & consumer

func main() {
	ch := make(chan int)

	// go func() {
	// 	time.Sleep(3 * time.Second)

	// 	fmt.Println(<-ch)
	// 	fmt.Println("3 second Goroutine finished")
	// }()
	// ch <- 1

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2 second Goroutine finished")
	}()

	go func() {
		// ch <- 1
		time.Sleep(3 * time.Second)
		fmt.Println("3 second Goroutine finished")
	}()

	receiver := <-ch // throws panic err after 3 sec as no value is received
	fmt.Println(receiver)

	fmt.Println("End of program")
}
