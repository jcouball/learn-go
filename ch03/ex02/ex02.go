package main

import "fmt"

// Write a program that defines a string variable called message with the value "Hi
// 👩 and and 👨" and prints the fourth rune in it as a character, not a number.

func main() {
	message := "Hi 👩 and and 👨"
	fmt.Println("message: ", message)
	runes := []rune(message)
	fmt.Printf("Fourth rune: %c\n", runes[3])
}
