package main

// embeddings allow us to inherit fields & methods from another structs

import "fmt"

type person struct {
	name string
	age  int
}

// Employee inherits from Person
// fileds of embedded struct are promoted to outer struct

type Employee1 struct {
	person // Anonymous field
	empId  string
	salary float64
}

type Employee2 struct {
	employeeInfo person // Embedded struct Named field
	empId        string
	salary       float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

// methods can be over-ridden by redefining them in outer struct
func (e Employee1) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.name, e.empId, e.salary)
}

func (e Employee2) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.employeeInfo.name, e.empId, e.salary)
}

func main() {

	emp1 := Employee1{
		person: person{name: "John", age: 30},
		empId:  "E001",
		salary: 40000,
	}

	fmt.Println("Name:", emp1.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age:", emp1.age)   // Same as above
	fmt.Println("Emp ID:", emp1.empId)
	fmt.Println("Salary:", emp1.salary)

	emp1.introduce() // over-riding "introduce" method of "person"

	fmt.Print("\n****************************************\n")

	emp2 := Employee2{
		employeeInfo: person{name: "John", age: 30},
		empId:        "E002", salary: 50000,
	}

	fmt.Println("Name:", emp2.employeeInfo.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age:", emp2.employeeInfo.age)   // Same as above
	fmt.Println("Emp ID:", emp2.empId)
	fmt.Println("Salary:", emp2.salary)

	emp2.introduce()

}
