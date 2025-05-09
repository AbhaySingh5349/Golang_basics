package main

// methods are functions associated to a "type" (not necessarily a struct)
// they enable us to define behaviours that operate on instances of type
// they are declared with a "receiver" which is "type" on which method operates on
// 2 types of "receivers": value & pointer

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver (when we do not want to modify receiver instance)
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver (if method wants to modify receiver instance or if we want to avoid copying large type)
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor // r.length = r.length * factor
	r.width *= factor
}

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is", area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle with a factor of 2 is", area)

	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.Area())
	fmt.Println(s.Rectangle.Area())

	num := MyInt(-5)
	num1 := MyInt(9)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())
}

type MyInt int

// Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}
