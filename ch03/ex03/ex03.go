package main

import "fmt"

// Write a program that defines a struct called `Employee` with three fields:
// `firstName`, `lastName`, and `id`. The first two fields are of type string, and
// the last field (`id`) is of type `int`. Create three instances of this struct
// using whatever values you’d like. Initialize the first one using the struct
// literal style without names, the second using the struct literal style with names,
// and the third with a `var` declaration. Use dot notation to populate the fields in
// the third struct. Print out all three structs.

func main() {
	type employee struct {
		firstName string
		lastName  string
		id        int
	}

	emp1 := employee{"John", "Doe", 10001}
	fmt.Println("emp1: ", emp1)

	emp2 := employee{firstName: "Jane", lastName: "Austin", id: 10002}
	fmt.Println("emp2: ", emp2)

	var emp3 employee
	emp3.firstName = "Bob"
	emp3.lastName = "Marley"
	emp3.id = 10003
	fmt.Println("emp3: ", emp3)
}
