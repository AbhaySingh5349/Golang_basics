package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// reading data efficiently

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio packageee!\nHow are you doing?"))

	// Reading byte slice
	data := make([]byte, 20)
	ub_idx, err := reader.Read(data) // store 20 bytes of data from reader
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", ub_idx, data[:ub_idx])

	line, err := reader.ReadString('\n') // reads from point where it left earlier till next line break
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	fmt.Println("Read line:", line)

	// struct which is wrapper around "io.writer" & provides bufferening for efficient writing
	writer := bufio.NewWriter(os.Stdout)

	// Writing byte slice
	write_data := []byte("Hello, bufio package!\n")
	n, err := writer.Write(write_data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	// Writing string
	str := "This is a string.\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}
