package main

import "fmt"

func modifyMap(m map[string]int) {
	m["key"] = 42 // This changes the original map
}

func modifySlice(s []int) {
	s[0] = 100       // This changes the original slice's data
	s = append(s, 5) // This only affects the local copy
}

func appendToSlice(s []int, val int) []int {
	return append(s, val)
}

func main() {
	// Maps example
	fmt.Println("=== Maps ===")
	myMap := map[string]int{"key": 10}
	fmt.Printf("Before modifyMap: %v\n", myMap)
	modifyMap(myMap)
	fmt.Printf("After modifyMap: %v\n", myMap) // prints map[key:42]

	// Slices example
	fmt.Println("\n=== Slices ===")
	mySlice := []int{1, 2, 3}
	fmt.Printf("Before modifySlice: %v\n", mySlice)
	modifySlice(mySlice)
	fmt.Printf("After modifySlice: %v\n", mySlice) // prints [100 2 3], not [100 2 3 5]

	// Proper way to append to a slice
	fmt.Println("\n=== Proper Append ===")
	mySlice2 := []int{1, 2, 3}
	fmt.Printf("Before appendToSlice: %v\n", mySlice2)
	mySlice2 = appendToSlice(mySlice2, 5)
	fmt.Printf("After appendToSlice: %v\n", mySlice2) // prints [1 2 3 5]
}
