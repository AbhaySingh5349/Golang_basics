package main

import (
	"flag"
	"fmt"
	"os"
)

// go run .\main.go -name "Abhay Singh" -age 25
// go run .\main.go -help (to get some info on flags)

func main() {

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args {
		fmt.Println("Argument", i, ":", arg)
	}

	// Define flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of the user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)

}
