package main

import "fmt"

// specifies allowable operations on a channel (eg. sending, receiving) hence Type Safety
// they enforce & document intended use of channel in function & goroutines
// they are not used as independent variables which we can declare

func main() {

	ch := make(chan int) // bi-directional channel
	producer(ch)
	consumer(ch)
}

// Send only channel
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
