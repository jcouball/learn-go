package main

import "fmt"

func main() {
	a := 20
	f := func() {
		fmt.Println(a) // Output: 20
		a = 30
	}
	f()
	fmt.Println(a) // Output: 30
}
