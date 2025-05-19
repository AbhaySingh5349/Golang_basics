package main

import (
	"fmt"
	"time"
)

// non-blocking operations on a channel allows "goroutine" to perform send/receive without getting stuck if channel is not ready
// they help in maintaining responsiveness & help goroutines from getting blocked indefinitely (avoid deadlocks)

func main() {

	ch := make(chan int)

	// === NON BLOCKING RECEIVE OPERATION (immediately executes "default" case if no channel is ready)
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No messages available.")
	}

	// === NON BLOCKING SEND OPERATION (when we don't have a receiver)
	select {
	case ch <- 1:
		fmt.Println("Sent message.")
	default:
		fmt.Println("Channel is not ready to receive.")
	}

	// === NON BLOCKING OPERATION IN REAL TIME SYSTEMS

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				// if value received on "quit channel"
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
}
