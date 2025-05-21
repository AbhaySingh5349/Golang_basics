package main

import (
	"fmt"
	"sync"
	"time"
)

// Wait Group is a synchronization primitive provide by "sync package"
// it is used to wait for collection of goroutines to complete their execution
// it helps in achieving synchronization & coordination of concurrent task ensuring all tasks are completed before moving forward
// it also help with resource management by performing cleanups after concurrent operations

// ============ BASIC EXAMPLE WITHOUT USING CHANNELS
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // when task is finished, dec counter in WG for num of goroutines it needs to wait for

// 	// wg.Add(1) // WRONG PRACTICE as execution happens very fast, so might be possible that "main goroutine" already reach "wg.Wait()"

// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate some time spent on processing the task
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	// WG knows it needs to wait for 3 calls of "wg.Done()" else it will not let execution flow after "wg.Wait()"
// 	wg.Add(numWorkers) // THIS IS THE CORRECT WAY OF ADDING NUM_WORKERS to COUNTER

// 	// Launch workers
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() // blocking mechanism
// 	fmt.Println("All workers finished")
// }

// EXAMPLE WITH CHANNELS
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Printf("WorkerID %d starting.\n", id)

// 	time.Sleep(time.Second) // simulate some work
// 	for task := range tasks {
// 		results <- task * 2
// 	}
// 	fmt.Printf("WorkerID %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5

// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)

// 	// we want to receive values in real time & not when all goroutines are done, so we extract "wg.Wait()" outside main thread
// 	// hence we print results as they are generated
// 	go func() {
// 		wg.Wait()
// 		close(results) // closes "results channel" once all go-routines are finished
// 	}()

// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}
// }

// CONSTRUCTION EXAMPLE
type Worker struct {
	ID   int
	Task string
}

// simulates a worker performing a task
// we will be modifiying fields of "worker struct"
func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func main() {
	var wg sync.WaitGroup

	// Define tasks to be performed by workers
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1) // adding worker count
		go worker.PerformTask(&wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Construction is finished
	fmt.Println("Contruction finished")
}
