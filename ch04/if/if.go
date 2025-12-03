package main

import (
	"fmt"
	"math/rand"
)

func basicIfStatement() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

func conditionalWithShortStatement() {
	if n := rand.Intn(10); n%2 == 0 {
		fmt.Println("Even number:", n)
	} else {
		fmt.Println("Odd number:", n)
	}
}

func main() {
	basicIfStatement()
	conditionalWithShortStatement()
}
