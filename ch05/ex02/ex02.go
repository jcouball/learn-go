package main

import (
	"fmt"
	"io"
	"os"
)

/*
Chapter 5, Exercise 2

Write a function called fileLen that has an input parameter of type string and
returns an int and an error. The function takes in a filename and returns the number
of bytes in the file. If there is an error reading the file, return the error. Use
defer to make sure the file is closed properly.

When implementing, use os.Open to open the file, file.Close to close the file, and
io.ReadAll to read the contents of the file.
*/

func fileLen(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	return len(data), nil
}

func main() {
	length, err := fileLen("README.md")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Length:", length)
	}
}
