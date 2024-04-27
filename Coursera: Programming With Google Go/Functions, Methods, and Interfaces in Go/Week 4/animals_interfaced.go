package main

import (
	"fmt"
)

// Animal interface defines the methods Eat(), Move(), and Speak()
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct represents a cow
type Cow struct {
	name string
}

// Eat method for Cow
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move method for Cow
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak method for Cow
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird struct represents a bird
type Bird struct {
	name string
}

// Eat method for Bird
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move method for Bird
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak method for Bird
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake struct represents a snake
type Snake struct {
	name string
}

// Eat method for Snake
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move method for Snake
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak method for Snake
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		var command string
		fmt.Print("> ")
		fmt.Scanln(&command)

		switch command {
		case "newanimal":
			var name, animalType string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter animal type (cow, bird, snake): ")
			fmt.Scanln(&animalType)

			switch animalType {
			case "cow":
				animals[name] = Cow{name}
			case "bird":
				animals[name] = Bird{name}
			case "snake":
				animals[name] = Snake{name}
			default:
				fmt.Println("Invalid animal type. Please choose cow, bird, or snake.")
			}

			fmt.Println("Created it!")
		case "query":
			var name, info string
			fmt.Print("Enter name: ")
			fmt.Scanln(&name)
			fmt.Print("Enter info (eat, move, speak): ")
			fmt.Scanln(&info)

			animal, ok := animals[name]
			if !ok {
				fmt.Println("Animal not found.")
				continue
			}

			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid info. Please choose eat, move, or speak.")
			}
		default:
			fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
		}
	}
}
