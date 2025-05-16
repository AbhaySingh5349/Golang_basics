package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// "back ticks in struct": used for "struct field tags" which provide metadata about fields (useful when working with DB or converting struct <-> json)

type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}

func main() {

	person := Person{FirstName: "Abhay", EmailAddress: "abhay@gmail.com"}

	// Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 30, EmailAddress: "jane@fakemail.com", Address: Address{City: "New York", State: "NY"}}

	jsondata1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsondata1))

	jsonData1 := `{"full_name": "Jenny Doe", "emp_id": "0009", "age": 30, "address": {"city": "San Jose", "state": "CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData1), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(employeeFromJson)
	fmt.Println("Jenny's Age increased by 5 years", employeeFromJson.Age+5)
	fmt.Println("Jenny's city:", employeeFromJson.Address.City)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonByteList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshalling to JSON:", err)
	}

	fmt.Println("JSON List:", string(jsonByteList))

	// Handling unknown JSON structures
	jsonData2 := `{"name": "Abhay", "age": 30, "address": {"city": "New York", "state": "NY"}}`

	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData2), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])

}
