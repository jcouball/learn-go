package main

import "fmt"

func main() {
	var b byte = 255
	b += 1

	var smallI int32 = 2_147_483_647
	smallI += 1

	var bigI uint64 = 18_446_744_073_709_551_615
	bigI += 1

	fmt.Println("b = ", b)
	fmt.Println("smallI = ", smallI)
	fmt.Println("bigI = ", bigI)
}
