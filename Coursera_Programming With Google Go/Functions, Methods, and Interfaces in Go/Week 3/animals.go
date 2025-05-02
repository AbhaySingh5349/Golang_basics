/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake.
Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the user makes a request.
*/

package main

import (
	"fmt"
)

// Animal struct to hold information about each animal
type Animal struct {
	food        string
	locomotion  string
	spokenSound string
}

// Eat method prints the animal's food
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move method prints the animal's locomotion
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak method prints the animal's spoken sound
func (a Animal) Speak() {
	fmt.Println(a.spokenSound)
}

func main() {
	// Define the three animals
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", spokenSound: "moo"},
		"bird":  {food: "worms", locomotion: "fly", spokenSound: "peep"},
		"snake": {food: "mice", locomotion: "slither", spokenSound: "hsss"},
	}

	// Infinite loop to accept user requests
	for {
		var animalName, infoRequested string

		// Prompt the user for input
		fmt.Print("> ")

		// Read user input
		fmt.Scan(&animalName, &infoRequested)

		// Get the corresponding animal
		animal, ok := animals[animalName]
		if !ok {
			fmt.Println("Invalid animal name. Please enter 'cow', 'bird', or 'snake'.")
			continue
		}

		// Process user request
		switch infoRequested {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid info requested. Please enter 'eat', 'move', or 'speak'.")
		}
	}
}
