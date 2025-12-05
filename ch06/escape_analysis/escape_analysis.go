package main

import "fmt"

type Person struct {
	Age int
}

func stackExample() {
	x := 10              // x is on the stack
	y := Person{Age: 30} // y is on the stack
	fmt.Printf("stackExample: x=%v, y=%v\n", x, y)
	// Both x and y are freed when the function returns
}

func heapExample() *Person {
	p := Person{Age: 30} // p escapes, so it's allocated on the heap
	return &p            // Returning pointer moves p to heap
}

func main() {
	stackExample()

	person := heapExample()
	fmt.Printf("heapExample returned: %v\n", person)

	fmt.Println("\nTo see escape analysis, run:")
	fmt.Println("  go build -gcflags=\"-m\" ./ch06/escape_analysis")
	fmt.Println("or")
	fmt.Println("  go run -gcflags=\"-m\" ./ch06/escape_analysis")
}
