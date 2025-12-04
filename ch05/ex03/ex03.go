package main

import "fmt"

/*
Chapter 5, Exercise 3

Write a function called prefixer that has an input parameter of type string and
returns a function that has an input parameter of type string and returns a
string. The returned function should prefix its input with the string passed into
prefixer. Use the following main function to test prefixer:

func main() {
    helloPrefix := prefixer("Hello")
		fmt.Println(helloPrefix("Bob")) // should print Hello Bob
		fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
*/

func prefixer(prefix string) func(str string) string {
	return func(str string) string {
		return prefix + " " + str
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // Output: Hello Bob
	fmt.Println(helloPrefix("Maria")) // Output: Hello Maria
}
