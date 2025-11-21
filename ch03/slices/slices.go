package main

import "fmt"

func append_elements() {
	// Appending elements to a slice
	//
	var x = []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x) // Output: [1 2 3 4 5 6]
}

func append_slice() {
	// Appending one slice to another
	//
	var x = []int{1, 2, 3}
	var y = []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x) // Output: [1 2 3 4 5 6]
}

func clear_slice() {
	// Clearing a slice
	//
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s)) // Output: [first second third] 3
	clear(s)
	fmt.Println(s, len(s)) // Output: [  ] 3
}

func slice_expressions() {
	// Slice Expressions
	//
	x := []string{"first", "second", "third", "forth"}
	fmt.Println("x[:2]:", x[:2])   // Output: ["first", "second"]
	fmt.Println("x[1:]:", x[1:])   // Output: ["second", "third", "forth"]
	fmt.Println("x[1:3]:", x[1:3]) // Output: ["second", "third"]
	fmt.Println("x[:]:", x[:])     // Output: ["first", "second", "third", "forth"]
}

func slice_expression_share_underlying_array() {
	// Slice expressions share the same underlying array
	//
	x := []string{"first", "second", "third", "forth"}
	y := x[1:3] // ["second", "third"]
	y[0] = "2nd"
	x[2] = "3rd"
	fmt.Println("x:", x) // Output: ["first", "2nd", "3rd", "forth"]
	fmt.Println("y:", y) // Output: ["2nd", "3rd"]
}

func full_slice_expression() {
	// Full slice expression to limit capacity
	//
	length := 10
	capacity := 100
	data := make([]string, length, capacity)
	fmt.Println(data) // Output: [          ]
}

func full_slice_expression_to_prevent_sharing_capacity() {
	// Full slice expression to prevent sharing capacity
	//
	x := []string{"first", "second", "third", "forth"}
	y := x[1:3:3] // ["second", "third"]
	y = append(y, "4th")
	fmt.Println("x:", x) // Output: ["first", "2nd", "3rd", "forth"]
	fmt.Println("y:", y) // Output: ["second", "third", "4th"]
}

func copy_slice() {
	// Copying a slice
	//
	x := []string{"first", "second", "third", "forth"}
	y := make([]string, 2)
	copy(y, x[1:3]) // ["second", "third"]
	y[0] = "2nd"
	x[2] = "3rd"
	fmt.Println("x:", x) // Output: ["first", "second", "3rd", "forth"]
	fmt.Println("y:", y) // Output: ["2nd", "third"]
}

func copy_overlapping_slices() {
	// Copying overlapping slices
	//
	x := []string{"first", "second", "third", "forth"}
	copy(x[1:], x[0:2])
	fmt.Println("x:", x) // Output: ["first", "first", "second", "third"]
}

func array_cant_be_larger_than_slice() {
	// Array can't be larger than the length of the slice
	//
	// x := []string{"first", "second", "third", "forth"}
	// y := [5]string(x) // Panics
	// fmt.Println("y:", y)
}

func main() {
	append_elements()
	append_slice()
	clear_slice()
	slice_expressions()
	slice_expression_share_underlying_array()
	full_slice_expression()
	full_slice_expression_to_prevent_sharing_capacity()
	copy_slice()
	copy_overlapping_slices()
	array_cant_be_larger_than_slice()
}
