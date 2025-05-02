/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user to enter a string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')

	// Remove the newline character from the input
	input = strings.TrimSuffix(input, "\n")

	// Convert the input string to lowercase for case-insensitive comparison
	input = strings.ToLower(input)

	fmt.Printf("string is %s\n", input)

	// Check if the input string starts with 'i', ends with 'n', and contains 'a'
	if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
