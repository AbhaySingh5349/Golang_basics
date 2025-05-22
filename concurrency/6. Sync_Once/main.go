package main

import (
	"fmt"
	"sync"
)

// ensures piece of code is executed only 1 time regardless of how many goroutines try to execute it
// usefull for initializing shared resources or performing set of tasks

var once sync.Once

func initialize() {
	fmt.Println("This will not be repeated no matter how many times we call this function using once.Do.")
}

func main() {

	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine # ", i)
			once.Do(initialize)

			// initialize()
		}()
	}
	wg.Wait()
}
