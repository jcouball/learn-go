package main

import "fmt"

func anonymous_struct() {
	var person struct {
		name string
		age  int
	}

	person.name = "James"
	person.age = 58

	fmt.Println(person) // Output: {James 58}
}

func anonymous_struct_literal() {
	person := struct {
		name string
		age  int
	}{
		name: "Mary",
		age:  32,
	}

	fmt.Println(person) // Output: {Mary 32}
}

func comparing_converting_structs() {
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{name: "James", age: 58}
	g := struct {
		name string
		age  int
	}(f)

	// The following compiles
	fmt.Println(f == g) // Output: true
}

func main() {
	anonymous_struct()
	anonymous_struct_literal()
	comparing_converting_structs()
}
