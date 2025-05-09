package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	elements []T // list of elements
}

// pointer receiver since we are modifying elements
func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

// pointer receiver since we are modifying elements
func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var default_val T
		return default_val, false
	}

	top := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1] // new slice from begining to len-1 idx
	return top, true
}

// we can still pass by reference to not create unnecessary copies of stack
func (s Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) display() {
	if len(s.elements) == 0 {
		fmt.Println("Stack is empty")
		return
	}

	for _, val := range s.elements {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func main() {
	x1, y1 := 1, 2
	x1, y1 = swap(x1, y1)
	fmt.Println(x1, ":", y1)

	x2, y2 := "Abhay", "Singh"
	x2, y2 = swap(x2, y2)
	fmt.Println(x2, ":", y2)

	fmt.Println("**********************")

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)

	intStack.display()

	intStack.pop()

	fmt.Println("intStack.isEmpty(): ", intStack.isEmpty())
	intStack.display()

	intStack.pop()
	intStack.pop()

	intStack.display()
	fmt.Println("intStack.isEmpty(): ", intStack.isEmpty())
}
