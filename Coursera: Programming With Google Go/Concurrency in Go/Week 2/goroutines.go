/*
Race condition occurs because the outcome of the program becomes unpredictable due to the interleaving of goroutines.
Each goroutine reads the value of counter, then increments it, and writes it back.
However, this read-modify-write sequence is not atomic, meaning that other goroutines can interleave between these steps.
As a result, the final value of counter depends on the order in which these interleavings occur.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(routine_num string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("routine: ", routine_num, "->", i+1)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers("1", &wg)
	go printNumbers("2", &wg)

	wg.Wait()

	fmt.Println("Done printing numbers.")
}
