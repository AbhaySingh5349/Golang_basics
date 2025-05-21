package main

import (
	"fmt"
	"sort"
)

// to sort a custom struct, we must define 3 methods: Len, Less, Swap

type Person struct {
	Name string
	Age  int
}

type SortBy func(p1, p2 *Person) bool

type personSorter struct {
	people   []Person
	criteria SortBy
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.criteria(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (criteria SortBy) Sort(people []Person) {
	ps := &personSorter{
		people:   people,
		criteria: criteria,
	}
	sort.Sort(ps)
}

func main() {
	// numbers := []int{5, 3, 4, 1, 2}
	// sort.Ints(numbers)
	// fmt.Println("Sorted numbers:", numbers)

	// nameSlice := []string{"John", "Anthony", "Steve", "Victor", "Walter"}
	// sort.Strings(nameSlice)
	// fmt.Println("Sorted strings:", nameSlice)

	// fruitSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	// sort.Slice(fruitSlice, func(i, j int) bool {
	// 	return fruitSlice[i][len(fruitSlice[i])-1] < fruitSlice[j][len(fruitSlice[j])-1]
	// })
	// fmt.Println("Sorted by last character:", fruitSlice)

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Anna", 35},
	}
	fmt.Println("Unsorted people:", people)

	ageAsc := func(p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}
	ageDesc := func(p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}
	lenName := func(p1, p2 *Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	SortBy(ageAsc).Sort(people)
	fmt.Println("Sorted by age (ascending):", people)

	SortBy(ageDesc).Sort(people)
	fmt.Println("Sorted by age (descending):", people)

	SortBy(name).Sort(people)
	fmt.Println("Sorted by name:", people)

	SortBy(lenName).Sort(people)
	fmt.Println("Sorted by length of name:", people)
}

// type ByAge []Person
// type ByName []Person

// func (a ByAge) Len() int {
// 	return len(a)
// }

// func (a ByAge) Less(i, j int) bool {
// 	return a[i].Age < a[j].Age
// }

// func (a ByAge) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }

// func (a ByName) Len() int {
// 	return len(a)
// }

// func (a ByName) Less(i, j int) bool {
// 	return a[i].Name < a[j].Name
// }

// func (a ByName) Swap(i, j int) {
// 	a[i], a[j] = a[j], a[i]
// }
