package main

import "fmt"

func failedUpdate(px *int) {
	x := 10
	px = &x
}

func failedUpdate2(px *int) {
	x2 := 20
	px = &x2
}

func update(px *int) {
	*px = 20
}

func main() {
	// Example 1: Can't make a nil pointer non-nil
	var f *int // f is nil
	failedUpdate(f)
	fmt.Printf("After failedUpdate, f is still: %v\n", f) // prints nil

	// Example 2: Must dereference to mutate the value
	x := 10
	fmt.Printf("Before updates, x = %v\n", x)

	failedUpdate2(&x)
	fmt.Printf("After failedUpdate2, x = %v\n", x) // prints 10

	update(&x)
	fmt.Printf("After update, x = %v\n", x) // prints 20
}
