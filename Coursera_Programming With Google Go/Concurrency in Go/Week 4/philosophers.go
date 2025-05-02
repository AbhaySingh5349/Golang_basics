/*
1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5
7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Chopstick represents a mutex for a chopstick.
type Chopstick struct{ sync.Mutex }

// Host represents a host that manages access to chopsticks.
type Host struct {
	concurrentEaters int // Number of philosophers currently eating
	sync.Mutex           // Mutex for synchronizing access to concurrentEaters
}

// requestPermission requests permission from the host to eat.
// It returns true if the permission is granted, false otherwise.
func (h *Host) requestPermission() bool {
	h.Lock()
	defer h.Unlock()
	if h.concurrentEaters < 2 {
		h.concurrentEaters++
		return true
	}
	return false
}

// releasePermission releases the permission obtained from the host.
func (h *Host) releasePermission() {
	h.Lock()
	defer h.Unlock()
	h.concurrentEaters--
}

// Philosopher represents a philosopher with a number and two chopsticks.
type Philosopher struct {
	number                        int        // Philosopher's number
	leftChopstick, rightChopstick *Chopstick // Chopsticks on the left and right of the philosopher
	host                          *Host      // Host for managing access to chopsticks
}

// eat simulates the eating behavior of a philosopher.
func (p *Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		// Request permission to eat from the host
		if p.host.requestPermission() {
			// Acquire chopsticks
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()

			// Start eating
			fmt.Printf("Philosopher %d starting to eat\n", p.number)
			time.Sleep(100 * time.Millisecond) // Simulate eating time

			// Finish eating
			fmt.Printf("Philosopher %d finishing eating\n", p.number)

			// Release chopsticks
			p.rightChopstick.Unlock()
			p.leftChopstick.Unlock()

			// Release permission to eat
			p.host.releasePermission()
		} else {
			// If permission is not granted, wait for a while
			fmt.Printf("Philosopher %d is waiting for permission to eat.\n", p.number)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	// Create chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// Create host
	host := &Host{}

	// Create philosophers
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
			host:           host,
		}
	}

	// Simulate philosophers dining
	var wg sync.WaitGroup
	for _, p := range philosophers {
		wg.Add(1)
		go p.eat(&wg)
	}
	wg.Wait()
}
