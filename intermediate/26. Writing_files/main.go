package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("writeByte.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
		return
	}

	// executed when "main" func returns
	defer file.Close() // defer can be used for "cleanup" actions

	// write data to file
	data := []byte("Hello World\nLet's Go!!")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data has been written to file successfully.")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Writing to writeString.txt complete.")
}
