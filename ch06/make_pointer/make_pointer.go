package main

import "fmt"

// makePointer is a helper function to turn a constant value into a pointer
func makePointer[T any](value T) *T {
	return &value
}

func main() {
	ptr := makePointer(42)        // returns *int pointing to 42
	strPtr := makePointer("hi")   // returns *string pointing to "hi"
	boolPtr := makePointer(true)  // returns *bool pointing to true
	floatPtr := makePointer(3.14) // returns *float64 pointing to 3.14

	fmt.Printf("ptr: %v, value: %v\n", ptr, *ptr)
	fmt.Printf("strPtr: %v, value: %v\n", strPtr, *strPtr)
	fmt.Printf("boolPtr: %v, value: %v\n", boolPtr, *boolPtr)
	fmt.Printf("floatPtr: %v, value: %v\n", floatPtr, *floatPtr)
}
