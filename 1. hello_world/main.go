// package = project = worksapce (collection of common src code files / to group together code with similar purpose)

// used for creating executable type package to spit out runnable file
package main

import "fmt"

// entry point of program
// without main method, compiler cannot generate binary executable
func main() {
	fmt.Print("hello world golang")
}

// with "go run main.go", -> compiles ".go file" into temporary executable file in memory & after compilation runs code

// in deployment process, we compile our code once & store it for future execution
// go build main.go -> generates binary executable and simply run it by "main.exe"
