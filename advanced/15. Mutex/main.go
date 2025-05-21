package main

import (
	"fmt"
	"sync"
)

// Mutual Exclusion is a synchronization primitive to prevent multiple goroutines from simultaneously accessing shared resourcs
// It ensures only 1 goroutine holds mutex at a given instant, thus avoiding race condition & data corruption
// principle used in concurrent programming to prevent multiple threads & processes accessing shared resources

// type counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {

// 	var wg sync.WaitGroup
// 	counter := &counter{} // pointer to instance of counter struct
// 	numGoroutines := 10

// 	// wg.Add(numGoroutines)

// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 				// counter.count++
// 			}
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("Final counter value: %d\n", counter.getValue())
// }

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
