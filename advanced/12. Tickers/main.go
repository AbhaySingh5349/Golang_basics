package main

import (
	"fmt"
	"time"
)

// tickers are used for performing recurring tasks periodically (eg. polling, regular updates)

// func main() {

// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at:", tick)
// 	// }

// 	// i := 0
// 	// for range ticker.C {
// 	// 	i++
// 	// 	fmt.Println(i)
// 	// }

// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}
// }

// ========= SCHEDULING LOGGING, PERIODIC TASKS, POLLING FOR UPDATES
// func periodicTask() {
// 	fmt.Println("Performing periodic task at:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second) // timer signals stops after 5 sec

	defer ticker.Stop() // runs after execution is done in main

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stopping ticker.")
			return
		}
	}
}
