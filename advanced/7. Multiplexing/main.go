package main

import (
	"fmt"
)

// multiplexing is process of handling concurrent operations within a goroutine for multiple channels simultaneously
// it's similar to "switch" case
// "select statement" facilitates multipleplexing by allowing goroutine to wait on multiple channel operations & react to whichever operation is ready first

// func main() {

// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()
// 	go func() {
// 		ch2 <- 2
// 	}()

// 	time.Sleep(500 * time.Millisecond)

// 	// only 1 case is executed at a given instant
// 	// num of channels for which we want data
// 	for range 2 {
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("Received from ch1:", msg)
// 			time.Sleep(time.Second)
// 		case msg := <-ch2:
// 			fmt.Println("Received from ch2:", msg)
// 			time.Sleep(time.Second)
// 		default:
// 			fmt.Println("No channels ready...")
// 		}
// 	}

// 	fmt.Println("End of program")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Received:", msg)
// 	case <-time.After(3 * time.Second): // return channel of type time.Time
// 		fmt.Println("Timeout.")
// 	}
// }

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for {
		select {
		case msg, isOpen := <-ch:
			// if channel is closed but we arw still trying to receive values, it returns default val of that data_type
			if !isOpen {
				fmt.Println("Channel closed")
				// clean up activities
				return
			}
			fmt.Println("Received:", msg)
		}
	}
}
