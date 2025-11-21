package main

import (
	"fmt"
)

func main() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("x+y =", x+y)
	fmt.Println("x-y =", x-y)
}
