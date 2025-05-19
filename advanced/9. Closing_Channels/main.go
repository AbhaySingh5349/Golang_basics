package main

import "fmt"

// we close channels to signal completion idicating no more data will be send over it (but can still receive in "buffered channel") which helps goroutines which are receiving data
// it also prevents resource leaks
// close channels using goroutine from "Sender" & not from "Receiver"
// closing already closed channel leads to "runtime panic"
// only 1 goroutine should be responsible for clsoing channel (avoids panic & race conditions)

// === Simple closing channel example
// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch) // close channel after sending all values
// 	}()

// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }

// RECEIVING FROM A CLOSED CHANNEL
// func main() {

// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch
// 	if !ok {
// 		fmt.Println("Channel is closed")
// 		return
// 	}
// 	fmt.Println(val)
// }

// RANGE OVER CLOSED CHANNEL
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for val := range ch {
// 		fmt.Println("Received:", val)
// 	}
// }

// clsoing a closed channel
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		close(ch)
// 		close(ch)
// 	}()
// 	time.Sleep(time.Second)
// }

func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

// receive & send only channels
func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	// ch2 will be holding some values
	for val := range ch2 {
		fmt.Println(val)
	}
}
