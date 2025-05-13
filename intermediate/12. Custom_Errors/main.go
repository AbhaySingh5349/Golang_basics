package main

import (
	"errors"
	"fmt"
)

// custom errors can encapsulate specific details about what went wrong by providing enhanced context (making it easier to debug)
// these are "types" which implement "err interface" which requires implementation of "Error()" method

// we can "pass by reference" to use same instance of "cutomError" struct as it's memory efficient
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.err)
}

type customError struct {
	code    int
	message string
	err     error // wrapped error
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Operation completed successfully!")
}

func doSomething() error {
	err := doSomethingElse() // eg. saving a file created an error
	if err != nil {
		// accessing "memory address" & not using "duplicate copy" of struct which modifies struct outside scope of function
		return &customError{
			code:    500,
			message: "Something went wrong",
			err:     err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}
