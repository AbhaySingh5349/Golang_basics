package main

import (
	"fmt"
	"time"
)

// Worker Pool is a Design Pattern to manage group of workers i.e goroutines

// ============= BASIC WORKER POOL PATTERN
// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker %d processing task %d\n", id, task)
// 		time.Sleep(time.Second) // Simulate work
// 		results <- task * 2     // sending processed task to "results channel"
// 	}
// }

// func main() {
// 	numWorkers, numJobs := 3, 10

// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	// Create workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)
// 	}

// 	// Send values to the tasks channel
// 	for i := range numJobs {
// 		tasks <- i
// 	}

// 	close(tasks) // close tasks channel once all values are sent to it

// 	// Collect the results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result:", result)
// 	}
// }

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate processing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d\n", req.numTickets, req.personID, req.cost)
		time.Sleep(time.Second) // simulate processing time
		results <- req.personID
	}

}

func main() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// start ticket processor/worker
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: (i + 1) * price}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
	}
}
