package main

import (
	"fmt"
	"strings"
)

// strings are "sequence of bytes"

func main() {

	str := "Hello Go Programming"
	fmt.Println("len(str): ", len(str))

	fmt.Printf("str[0]: %c\n", str[0])
	fmt.Println("split: ", strings.Split(str, " "))
	fmt.Println("str Contains Go ? ", strings.Contains(str, "Go"))
	fmt.Println("count of 'o': ", strings.Count(str, "o"))

	replaced_str := strings.Replace(str, "Go", "World", 1)
	fmt.Println(str, " -> ", replaced_str)

	fmt.Println("uppercase: ", strings.ToUpper(str))

	fmt.Println("Repeat: ", strings.Repeat("foo ", 3))

	str1, str2 := "Hello", "World"
	res := str1 + "_" + str2
	fmt.Println("append res: ", res)

	fruites := []string{"apple", "mango", "grapes"}
	fmt.Println("joined: ", strings.Join(fruites, ", "))

	// STRING BUILDER (memory efficient & avoids memory leaks)
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result := builder.String()
	fmt.Println(result)

	// Using Writerune to add a character
	builder.WriteRune(' ') // runes are character, so ' '
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}
