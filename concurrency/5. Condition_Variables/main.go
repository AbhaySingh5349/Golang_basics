package main

import (
	"fmt"
	"sync"
	"time"
)

// condition variable is synchronization primitive which allows goroutines to wait for conditions to met while holding lock
// it signals goroutines about state change
// condition variable is used in conjunction with mutex & help in achieving complex synchronization scenario

// goroutine must acquire "mutex" before waiting or signaling condition variable
// "Broadcast" wakes up all gorouties waiting on condition variable

const bufferSize = 5

type buffer struct {
	items []int      // slice of values to hold in buffer
	mu    sync.Mutex // mutex to ensure "exclusive access" to buffer
	cond  *sync.Cond // condition variable for synchronization (allows goroutines to wait for signal changes in state)
}

func newBuffer(size int) *buffer {
	// initialize buffer instance
	b := &buffer{
		items: make([]int, 0, size), // empty slice with max capacity of size
		// cond: sync.NewCond(&b.mu), we cannot do this since we need reference to instance "b"
	}
	b.cond = sync.NewCond(&b.mu) // initialize condition with pointer to buffers mutex (has Lock(), Unlock() methods)
	return b
}

func (b *buffer) produce(item int) {
	b.mu.Lock()         // locking to ensure "exclusive access"
	defer b.mu.Unlock() // release lock after func ends

	// goroutine wait condition (check if buffer has reached capacity)
	for len(b.items) == bufferSize {
		b.cond.Wait() // makes goroutine sleep
	}

	b.items = append(b.items, item) // appending item to slice
	fmt.Println("Produced:", item)

	b.cond.Signal() // wakes up cosnumer informing about new item
}

func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait()
		// This function stops doing anything and waits for other function to append to the slice
	}

	item := b.items[0]
	b.items = b.items[1:] // extarct 1st value & creating new slice
	fmt.Println("Consumed:", item)

	b.cond.Signal()

	return item
}

func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 20 {
		b.produce(i + 1)
		time.Sleep(100 * time.Millisecond)
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 20 {
		b.consume()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup // producer & consumer use it to synchronize completion of goroutines

	wg.Add(2) // goroutines
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}
