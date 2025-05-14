package main

import (
	"fmt"
	"time"
)

func main() {
	layout1 := "2006-01-02T15:04:05Z07:00"
	str1 := "2024-07-04T14:30:18Z"

	t1, err := time.Parse(layout1, str1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("t1: ", t1)

	str2 := "Jul 03, 2024 03:18 PM"
	layout2 := "Jan 02, 2006 03:04 PM"

	t2, err := time.Parse(layout2, str2)
	fmt.Println("t2: ", t2)
}
