package main

import (
	"fmt"
	"math/rand"
)

// Exercise instructions:
// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.

// generateRandomInts creates a slice of 100 random integers between 0 and 99 inclusive.
func generateRandomInts() []int {
	randomInts := []int{}

	for range 100 {
		randomInts = append(randomInts, rand.Intn(100))
	}

	return randomInts
}

func main() {
	randomInts := generateRandomInts()
	fmt.Println(randomInts)
}
