package main

import (
	"flag"
	"fmt"
	"os"
)

// way to organise cmd line interfaces in hierarchical structures allowing diff. functionalities be grouped under main commands

// while using :sub-commands, we need to use = for assigning values
// go run .\main.go firstSub -processing=true -bytes=2048

func main() {

	// common flag for multiple sub-commands
	stringFlag := flag.String("user", "Guest", "Name of the user")
	flag.Parse()
	fmt.Println("stringFlag: ", *stringFlag)

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	// stores pointer to actual value
	firstFlag := subcommand1.Bool("processing", false, "Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subcommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:", *firstFlag) // de-referencing
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language:", *flagsc2)
	default:
		fmt.Println("no subcommand entered!")
		os.Exit(1)
	}

}
