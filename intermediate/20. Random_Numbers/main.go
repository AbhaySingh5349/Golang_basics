package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// seeding a random num generator
	// const_seed_val := rand.New(rand.NewSource(11))
	// fmt.Println("const_val: ", const_seed_val.Intn(101)) // generate fixed num between [0, 101)

	// var_seed_val := rand.New(rand.NewSource(time.Now().Unix()))
	// fmt.Println("var_seed_val: ", var_seed_val.Intn(101)) // even though we have seeded, but it's varying with time

	// fmt.Println(rand.Intn(6) + 5)
	// fmt.Println(val.Intn(101))
	// fmt.Println(rand.Float64()) // between 0.0 and 1.0

	for {
		// Show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// show the results
		fmt.Printf("You rolled a %d and a %d.\n", die1, die2)
		fmt.Println("Total:", die1+die2)

		// Ask of the user wants to roll again
		fmt.Println("Do you want to roll again? (y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}
