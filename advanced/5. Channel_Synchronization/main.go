package main

import (
	"fmt"
	"time"
)

// ensures data is exchanged between goroutines, coordinates execution flow to avoid race-condition

// func main() {

// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second) // simulating task
// 		done <- struct{}{}          // sending value to channel
// 	}()

// 	// receiving value from channel (until it receives, main channel is blocked for un-buffered channel or full)
// 	<-done
// 	fmt.Println("Finished.")

// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		// fmt.Println("Sending value...")
// 		// time.Sleep(200 * time.Millisecond)
// 		ch <- 9 // Blocking until the value is received
// 		fmt.Println("Sent value")
// 	}()

// 	fmt.Println("in main after initiating goroutine")

// 	value := <-ch // Blocking until a value is sent
// 	fmt.Println(value)

// 	time.Sleep(500 * time.Millisecond) // let fmt.Println("Sent value") get executed before program termination
// }

// ========= SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE & RECEIVED
// func main() {
// 	numGoroutines := 3        // controls buffer
// 	done := make(chan int, 3) // buffer of 3 int value

// 	// all goroutines are executed consecutively which are assigned to available threads by "go runtime scheduler"
// 	// we cannot control assignement part, so any goroutine can finish first
// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // SENDING SIGNAL OF COMPLETION
// 		}(i)
// 	}

// 	// we need to ensure we have receivers for all goroutines else memory leak
// 	for range numGoroutines {
// 		<-done // Wait for each goroutine to finish, WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETION
// 	}

// 	fmt.Println("All goroutines are complete")
// }

// ========== 	SYNCHRONIZING DATA EXCHANGE
func main() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) // Channel closed before Goroutine could send a value to the channel

	// when we loop over channel, we create "receiver" each time (not creating receiver using <- operator)
	// but we might loop even when no val is to be received, so we need to "close" channel once all data is sent
	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // Loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed
}
