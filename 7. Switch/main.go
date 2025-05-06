package main

import "fmt"

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case float64:
		fmt.Println("x is float")
	case string:
		fmt.Println("x is string")
	default:
		fmt.Println("x is invalid type")
	}
}

func main() {
	checkType(20)
	checkType(3.14)
	checkType("hello")
	checkType(true)

	day := "Mon"

	switch day {
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("Weekday")
	case "Sat", "Sun":
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	num := 10

	switch {
	case num < 10:
		fmt.Println("Num < 10")
	case num >= 10:
		fmt.Println("Num >= 10")
	}

	switch {
	case num < 10:
		fmt.Println("Num < 10")
	case num >= 10:
		fmt.Println("Num >= 10")
		fallthrough
	case num == 10:
		fmt.Println("Num is 10")
	}
}
