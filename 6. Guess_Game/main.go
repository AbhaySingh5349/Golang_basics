package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(10) + 1

	fmt.Println("welcome to num guessing game by chosing num between 1 - 10")
	fmt.Println("User Guess ?")

	var guess int
	for {
		fmt.Println("enter guess num: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("You Won!!")
			break
		} else if guess > target {
			fmt.Println("guess is higher than target")
		} else if guess < target {
			fmt.Println("guess is lower than target")
		}
	}
}
