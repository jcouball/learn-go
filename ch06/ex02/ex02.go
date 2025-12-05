package main

import "fmt"

/*
Chapter 6, Exercise 2

Write two functions:

- The UpdateSlice function takes in a []string and a string. It sets the last
  position in the passed-in slice to the passed-in string. At the end of UpdateSlice,
  print the slice after making the change.
- The GrowSlice function also takes in a []string and a string. It appends the string
  onto the slice. At the end of GrowSlice, print the slice after making the change.

Call these functions from main. Print out the slice before each function is called
and after each function is called.

Do you understand why some changes are visible in main and why some changes are not?
*/

/*
Solution

The changes in UpdateSlice are visibile in main because the slice length and capacity
was not changed.

The changes in GrowSlice are not visible in main because the length and possibly
capacity in the slice header (which contains length, capacity, pointer) was changed.
Since the slice header is passed by value, a local copy of the slice is made in
GrowSlice when the length is changed. Because of this, there is no way to update the
slice in the calling function.
*/

func UpdateSlice(strings []string, lastElement string) {
	fmt.Println("UpdateSlice: ", strings)
	strings[len(strings)-1] = lastElement
	fmt.Println("UpdateSlice: ", strings)
}

func GrowSlice(strings []string, lastElement string) {
	strings = append(strings, lastElement)
	fmt.Println("GrowSlice: ", strings)
}

func main() {
	fmt.Println("For UpdateSlice")
	slice1 := []string{"One", "Two", "Three"}
	fmt.Println("Before: ", slice1) // Output: [One Two Three]
	UpdateSlice(slice1, "New")
	fmt.Println("After: ", slice1) // Output: [One Two New]

	fmt.Println("For GrowSlice")
	slice2 := []string{"One", "Two", "Three"}
	fmt.Println("Before: ", slice2) // Output: [One Two Three]
	GrowSlice(slice2, "New")
	fmt.Println("After: ", slice2) // Output: [One Two Three]
}
