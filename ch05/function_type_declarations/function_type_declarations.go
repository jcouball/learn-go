package main

import "fmt"

type transformStringToInt func(string) int

func sumOfBytes(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		result += int(str[i])
	}
	return result
}

func countVowels(str string) int {
	count := 0
	vowels := "aeiouAEIOU"
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(vowels); j++ {
			if str[i] == vowels[j] {
				count++
				break
			}
		}
	}
	return count
}

// processString accepts any function matching the transformStringToInt signature
func processString(str string, transformer transformStringToInt) {
	result := transformer(str)
	fmt.Printf("Result: %d\n", result)
}

func main() {
	// Use the function type for variable declarations
	var operation transformStringToInt

	operation = sumOfBytes
	fmt.Println(operation("James")) // Output: 496

	operation = countVowels
	fmt.Println(operation("James")) // Output: 2

	// Use the function type as a parameter
	processString("Hello", sumOfBytes)  // Output: Result: 500
	processString("Hello", countVowels) // Output: Result: 2
}
