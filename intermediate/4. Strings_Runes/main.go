package main

import (
	"fmt"
	"unicode/utf8"
)

// string: sequence of bytes, immutable in nature
// Rune: dealing with individual characters or doing character-level manipulation (like iterating over runes in a string).

func main() {
	msg_1 := "Hello,\tGo"
	msg_2 := "Hello, \rGo!" // move pointer to 1st char of line

	fmt.Println(msg_1)
	fmt.Println(msg_2)

	fmt.Println("ASCII val of 1st char: ", msg_1[0])
	fmt.Printf("1st char %c\n", msg_1[0])

	for i, ch := range msg_1 {
		fmt.Printf("char at idx %d is %c\n", i, ch)
	}

	fmt.Println("Rune Count in str_1: ", utf8.RuneCountInString(msg_1))

	greeting, name := "hello", "abhay"
	fmt.Println(greeting + " " + name)

	str1, str2 := "Apple", "Banana"
	fmt.Println(str1 < str2) // ASCII values are compared

	var ch rune = 'a'
	fmt.Printf("rune: %c", ch)

	rune_str := string(ch)
	fmt.Printf("Type of rune_str %T", rune_str)
}
