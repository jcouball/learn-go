package main

import "fmt"

var myFuncVariable func(string) int

func sumOfBytes(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		result += int(str[i])
	}
	return result
}

func main() {
	myFuncVariable = sumOfBytes
	fmt.Println(myFuncVariable("James")) // Output: 496
}
