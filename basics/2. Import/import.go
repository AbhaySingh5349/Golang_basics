package main

import (
	"fmt"
	custom_http "net/http"
)

func main() {
	fmt.Println("Hello, Go Standard libraries")

	resp, err := custom_http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("HTTP response status: ", resp.Status)
}
