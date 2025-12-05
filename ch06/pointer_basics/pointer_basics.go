package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	// Address operator (&) examples
	x := "hello"
	pointerToX := &x

	person := Person{Name: "Alice"}
	pointerToName := &person.Name // addressable expression

	numbers := [3]int{1, 2, 3}
	pointerToFirst := &numbers[0] // addressable expression

	fmt.Printf("pointerToX points to: %v\n", pointerToX)
	fmt.Printf("pointerToName points to: %v\n", pointerToName)
	fmt.Printf("pointerToFirst points to: %v\n", pointerToFirst)

	// Indirection operator (*) examples
	y := 10
	pointerToY := &y
	fmt.Printf("pointerToY address: %v\n", pointerToY)
	fmt.Printf("*pointerToY value: %v\n", *pointerToY)
	z := 5 + *pointerToY
	fmt.Printf("5 + *pointerToY = %v\n", z)

	// Pointer type declaration
	a := 10
	var pointerToA *int
	pointerToA = &a
	fmt.Printf("pointerToA points to: %v, value: %v\n", pointerToA, *pointerToA)

	// Using new() function
	var b = new(int)
	fmt.Printf("b == nil: %v\n", b == nil)
	fmt.Printf("*b: %v\n", *b)

	// Pointer to a variable
	var c string
	d := &c
	fmt.Printf("d points to empty string: %v\n", *d == "")
}
