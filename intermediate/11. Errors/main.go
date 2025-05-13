package main

import (
	"errors"
	"fmt"
	"math"
)

// errors are represented by Error interface (built-in type to indicate error condition)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math err: square root of -ve num does not exists")
	}
	return math.Sqrt(x), nil
}

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("err: empty data")
	}

	// process data

	return nil
}

func inBuiltErrHandler() {
	if sqrt_val, err := sqrt(15); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt: ", sqrt_val)
	}

	if sqrt_val, err := sqrt(-15); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt: ", sqrt_val)
	}

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("data processed...")
	}
}

type myError struct {
	message string
}

// since we need to modify "msg" for error, so we pass by reference
// go bhas built-in error interface which has only 1 method "Error()" which returns "str" as custom err message
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

func readData() error {
	err := readConfig()
	if err != nil {
		// "%w" is formatting_verb for custom error string
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}

func main() {
	// inBuiltErrHandler()

	// err1 := eprocess()
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// println("")

	err2 := readData()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Data read successfully.")
}
