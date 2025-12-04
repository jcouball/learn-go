package main

import (
	"fmt"
	"sort"
)

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"John", "Smith", 37},
		{"Bob", "Dole", 23},
		{"Fred", "Pont", 18},
	}
	fmt.Println(people) // Output: [{John Smith 37} {Bob Dole 23} {Fred Pont 18}]

	// sort by last name
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people) // Output: [{Bob Dole 23} {Fred Pont 18} {John Smith 37}]

	// sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people) // Output: [{Fred Pont 18} {Bob Dole 23} {John Smith 37}]
}
