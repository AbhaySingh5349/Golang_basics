package main

import "fmt"

var globalVariable = "gv" // limited to package

func main() {
	var age int
	fmt.Println("age: ", age)

	var f_name string = "abhay"
	var m_name = "singh"

	// this notation cannot be used for "Global Variables"
	num := 9
	l_name := "bassi"

	globalVariable := "new_gv"

	fmt.Println("f_name: ", f_name, " , m_name: ", m_name, " , l_name: ", l_name, " with num: ", num, " with gv: ", globalVariable)
}
