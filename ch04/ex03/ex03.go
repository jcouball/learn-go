package main

import "fmt"

// Exercise instructions:
//
// In main, declare an int variable called total. Write a for loop that uses a
// variable named i to iterate from 0 (inclusive) to 10 (exclusive). The body of the
// for loop should be as follows:
//
//   total := total + i
//   fmt.Println(total)
//
// After the for loop, print out the value of total. What is printed out? What is the
// likely bug in this code?
//

// Answer:
//
// The likely bug in this code is that total is re-declared within the loop and
// shadows the total variable outside the loop.
//
// This means that total is re-initialized to its zero value on each loop iteration causing
// the value of i (or i + 0) to be printed within the loop.
//
// Since the total variable inside the loop shadows the total variable outside the loop,
// the value of total is never updated after its initialization to its zero value.
//

func main() {
	total := 0
	for i := range 10 {
		total := total + i
		fmt.Println(total) // Prints the value of i
	}
	fmt.Println(total) // Prints 0
}
