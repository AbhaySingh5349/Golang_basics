package main

import (
	"context"
	"fmt"
	"time"
)

// "context" can carry key-val pairs. These are instances of struct i.e Object
// object which carry info about "deadlines", "cancellation signals" & "req. scoped values in APIs & goroutines"
// they manage "lifecycle of processes" & to signal when operation should be aborted

// ======= DIFFERENCE BETWEEN CONTEXT.TODO AND CONTEXT.BACKGROUND
// func main() {

// 	todoContext := context.TODO() // placeholder for context which does not carry any signals or deadlines (when we are unsure of which context to use)
// 	contextBkg := context.Background() // returns an empty, base context. It's the root context from which all other contexts should be derived (when we're sure we don’t need timeout, cancellation, or values)

// 	ctx := context.WithValue(todoContext, "name", "John")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))
// }

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done(): // if any value is received i.e cancellation / completion signal, operation is done
		return "Operation canceled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

func main() {
	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO():", result)

	ctx = context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // "WithTimeout" needs a parent context (context get's cancelled after 1 sec but it still retains data it contains)
	defer cancel()                                         // cancel before end of main

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result from timeout context:", result)

	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result after timeout:", result)
}

// func doWork(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Work cancelled:", ctx.Err())
// 			return
// 		default:
// 			fmt.Println("Working...")
// 		}
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func main() {
// 	ctx := context.Background()
// 	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // timer of the context starts here. You have this specified time duration to use this context. After this time duration, the context will send a cancelation signal.

// 	ctx, cancel := context.WithCancel(ctx) // timer of the context starts here. You have this specified time duration to use this context. After this time duration, the context will send a cancelation signal.

// 	go func() {
// 		time.Sleep(2 * time.Second) // simulating a heavy task. Consider this a heavy time consuming operation
// 		cancel()                    // manually canceling only after the task is finished
// 	}()

// 	ctx = context.WithValue(ctx, "requestID", "hdsjf3234324234")
// 	ctx = context.WithValue(ctx, "name", "John")
// 	ctx = context.WithValue(ctx, "IP", "dsd.34.4332.34")
// 	ctx = context.WithValue(ctx, "OS", "Operating System")

// 	go doWork(ctx)

// 	time.Sleep(3 * time.Second)

// 	requestID := ctx.Value("requestID")
// 	if requestID != nil {
// 		fmt.Println("Request ID:", requestID)
// 	} else {
// 		fmt.Println("No request ID found.")
// 	}

// 	logWithContext(ctx, "This is a test log message")
// }

// func logWithContext(ctx context.Context, message string) {
// 	requestIDVal := ctx.Value("requestID")
// 	log.Printf("RequestID: %v - %v", requestIDVal, message)
// }
