package main

import (
	"fmt"
	"time"
)

// timers let's us schedule event to occur periodically
// timeouts to limit how long a particular operation should wait
// delays to schedule operation to occur with some delay (eg. time.Sleep -> blocking in nature)
// to avoid memory leaks, always stop timers

// ======== BASIC TIMER USE
// func main() {
// 	// time.Sleep(time.Second)
// 	fmt.Println("Starting app.")
// 	timer := time.NewTimer(2 * time.Second) // non-blocking in nature
// 	fmt.Println("Waiting for timer.c")

// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}

// 	timer.Reset(time.Second)
// 	fmt.Println("Timer reset")

// 	<-timer.C // blocking in nature (channel which gives us time)
// 	fmt.Println("Timer expired")
// }

// ============= TIMEOUT
// func longRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(3 * time.Second) // creates a channel that will send a value after 3 seconds.
// 	done := make(chan bool)

// 	go func() {
// 		longRunningOperation()
// 		done <- true // send completion info
// 	}()

// 	// select statement waits for the first available channel to return
// 	select {
// 	case <-timeout:
// 		fmt.Println("operation timed out")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}
// }

// =========== SCHEDULING DELAYED OPERATIONS

// func main() {
// 	timer := time.NewTimer(2 * time.Second) // non blocking timer starts

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed by goroutine")
// 	}()

// 	fmt.Println("Waiting in main...")
// 	time.Sleep(4 * time.Second) // blocking timer starts (to let goroutine finish up it's task)
// 	fmt.Println("End of the program")
// }

func main() {
	timer1 := time.NewTimer(1 * time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	// handle multiple timer channels
	select {
	case <-timer1.C:
		fmt.Println("Timer1 expired")
	case <-timer2.C:
		fmt.Println("Timer2 expired")
	}
}
