package main

import "fmt"

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	timesTwo := makeMult(2)
	timesThree := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(timesTwo(i), timesThree(i))
	}
}
