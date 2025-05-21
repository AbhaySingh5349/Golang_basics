package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{} // using empty struct in channel: when we want to count or signal something without storing any actual data
	refillTime time.Duration // how often new token is added to channel
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit), // buffered channel
		refillTime: refillTime,
	}
	for range rateLimit {
		rl.tokens <- struct{}{} // filling channel with initial set of tokens
	}

	go rl.startRefill() // goroutine running in background which re-fills empty slots in token channel

	return rl
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime) // creates new token periodically
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C: // receive values from ticker
			select {
			case rl.tokens <- struct{}{}: // sending empty struct i.e token after every tick if there's an empty slot in tokens channel
			default:
			}
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens: // remove value from token channel if it exists
		return true
	default:
		return false
	}
}

func main() {
	// 5 req / sec accepted
	rateLimiter := NewRateLimiter(5, time.Second)

	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}

//1 0 ms		First Request Allowed	5 tokens left
//2 200 ms		Second Request Allowed	4 tokens left
//3 400 ms		Third Request Allowed	3 tokens left
//4 600 ms		Fourth Request Allowed	2 tokens left
//5 800 ms		Fifth Request Allowed	1 tokens left
//6 1000 ms		Sixth Request Allowed	(not 0) 1 tokens left the startRefill function executes and adds one more token
//7 1200 ms		Seven Request Denied	tokens left
//8 1400 ms		Eight Request Denied	tokens left
//9 1600 ms		Ninth Request Denied	tokens left
//10 1800 ms		Tenth Request Denied
