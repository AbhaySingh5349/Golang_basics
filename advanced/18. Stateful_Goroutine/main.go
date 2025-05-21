package main

import (
	"fmt"
	"time"
)

// stateful goroutine maintains & updates which affects its behaviour

type StatefulWorker struct {
	count int
	ch    chan int
}

func (w *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch: // receiving value from channel
				w.count += value
				fmt.Println("Current count:", w.count)

			}
		}
	}()
}

// listens to values coming on channel
func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func main() {
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}

	stWorker.Start()

	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
