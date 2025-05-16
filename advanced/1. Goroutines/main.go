package main

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave main thread and run in background and come back to join main thread once functions are finished/ready to return any value
// Goroutines do not stop the program flow and are non blocking
// Goriutine ~ Thread
// similar to async/await in JS

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Num: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Char: ", string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in doWork.")
}

func main() {
	fmt.Println("Beginning program.")

	go sayHello()
	fmt.Println("After sayHello() function.")

	var err error

	go func() {
		err = doWork()
	}()

	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}
