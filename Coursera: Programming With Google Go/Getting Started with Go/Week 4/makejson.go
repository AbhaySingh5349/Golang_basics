/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Prompt the user to enter a name
	var name string
	fmt.Print("Enter a name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Prompt the user to enter an address
	var address string
	fmt.Print("Enter an address: ")
	_, err = fmt.Scanln(&address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create a map and add the name and address to it
	data := map[string]string{
		"name":    name,
		"address": address,
	}

	// Marshal the map into JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON object
	fmt.Println(string(jsonData))
}
