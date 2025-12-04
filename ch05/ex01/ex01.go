package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
Exercise instructions:

The simple calculator program doesn’t handle one error case: division by zero. Change
the function signature for the math operations to return both an int and an error. In
the div function, if the divisor is 0, return errors.New("division by zero") for the
error. In all other cases, return nil. Adjust the main function to check for this
error.
*/

func add(i int, j int) (int, error) {
	return i + j, nil
}

func sub(i int, j int) (int, error) {
	return i - j, nil
}

func mul(i int, j int) (int, error) {
	return i * j, nil
}

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue

		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
