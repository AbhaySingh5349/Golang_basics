package main

import "fmt"

func main() {

	// Printing Functions
	// fmt.Print("Hello ")
	// fmt.Print("World!")
	// fmt.Print(12, 456)

	// fmt.Println("Hello ")
	// fmt.Println("World!")
	// fmt.Println(12, 456)

	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// Formatting Functions
	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Print("\nfmt.Sprint: ", s)

	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Print("\nfmt.Sprintln: ", s)

	// sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	// fmt.Println(sf)

	// Scanning Functions
	// var name string
	// var age int

	// fmt.Print("Enter your name and age:")
	// fmt.Scan(&name, &age)
	// fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Error Formatting Functions

	err := checkAge(19)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
