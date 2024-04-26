/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

/*
file.txt

Krishna Devi
Ajay Kumar
Archana Singh
Shivangi Singh
Abhay Singh
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the name struct
type Name struct {
	Fname string
	Lname string
}

func main() {
	file, err := os.Create("names.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write names to the file
	input_names := []string{"Krishna Devi", "Ajay Kumar", "Archana Singh", "Shivangi Singh", "Abhay Singh"}
	for _, name := range input_names {
		_, err := file.WriteString(name + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scanln(&filename)

	// Open the file
	file, err = os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a slice to hold the structs
	var names []Name

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into first name and last name
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Println("Error: Invalid line format")
			continue
		}

		// Create a struct and add it to the slice
		name := Name{
			Fname: parts[0],
			Lname: parts[1],
		}
		names = append(names, name)
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterate through the slice of structs and print the first and last names
	fmt.Println("Names read from file:")
	for _, name := range names {
		fmt.Printf("%s %s\n", name.Fname, name.Lname)
	}
}
