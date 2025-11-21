package main

import "fmt"

func main() {
	var x int = 65
	y := string(rune(x))
	fmt.Println(y) // Outputs "A" and not "65"

	// Converting string to []byte and []rune
	//
	var s string = "Hello, 🌞"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs) // Output: [72 101 108 108 111 44 32 240 159 140 158]
	fmt.Println(rs) // Output: [72 101 108 108 111 44 32 127774]
}
