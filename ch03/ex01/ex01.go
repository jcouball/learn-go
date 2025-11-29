package main

import "fmt"

// Chapter 3, Exercise 1
//
// Write a program that defines a variable named greetings of type slice of strings
// with the following values: "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
// Create a subslice containing the first two values; a second subslice with the
// second, third, and fourth values; and a third subslice with the fourth and fifth
// values. Print out all four slices.

func main() {
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	fmt.Println("greetings: ", greetings)

	subslice1 := greetings[0:2]
	fmt.Println("subslice1: ", subslice1)

	subslice2 := greetings[1:4]
	fmt.Println("subslice2: ", subslice2)

	subslice3 := greetings[3:5]
	fmt.Println("subslice3: ", subslice3)
}
