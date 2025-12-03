package main

import (
	"fmt"
	"math/rand"
)

// Exercise instructions:
//
//
// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
//
// For each value in the slice, apply the following rules:
//
// - If the value is divisible by 2, print “Two!”
// - If the value is divisible by 3, print “Three!”
// - If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
// - Otherwise, print “Never mind”.
//

// generateRandomInts creates a slice of 100 random integers between 0 and 99 inclusive.
func generateRandomInts() []int {
	randomInts := []int{}

	for range 100 {
		randomInts = append(randomInts, rand.Intn(100))
	}

	return randomInts
}

// create a slice of strings that maps the given slice of ints to a string described
// in the instructions
func generateDivisibilityStrings(ints []int) []string {
	result := []string{}

	for _, v := range ints {
		switch {
		case v%2 == 0 && v%3 == 0:
			result = append(result, "Six!")
		case v%2 == 0:
			result = append(result, "Two!")
		case v%3 == 0:
			result = append(result, "Three!")
		default:
			result = append(result, "Never mind")
		}
	}

	return result
}

func main() {
	fmt.Println(generateDivisibilityStrings(generateRandomInts()))
}
