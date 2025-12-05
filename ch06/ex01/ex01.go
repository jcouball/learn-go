package main

import "fmt"

/*
Chapter 6, Exercise 1

Create a struct named Person with three fields: FirstName and LastName of type string
and Age of type int.

Write a function called MakePerson that takes in firstName, lastName, and age and
returns a Person.

Write a second function MakePersonPointer that takes in firstName, lastName, and age
and returns a *Person.

Call both from main.

Compile your program with go build -gcflags="-m". This both compiles your code and
prints out which values escape to the heap.

Are you surprised about what escapes?
*/

/*

Solution

My guess is that the Person structure returned from MakePersonPointer escapes the
stack because even though it is a pointer to a struct that only has primitive values,
a pointer to a Person is returned which makes it incompatible with the stack.

Here is what the compiler output:

```
$ ex01 $ go build -gcflags="-m"
# learn-go/ch06/ex01
./ex01.go:72:6: can inline MakePerson
./ex01.go:76:6: can inline MakePersonPointer
./ex01.go:82:18: inlining call to MakePerson
./ex01.go:83:25: inlining call to MakePersonPointer
./ex01.go:88:24: inlining call to MakePerson
./ex01.go:88:13: inlining call to fmt.Println
./ex01.go:89:31: inlining call to MakePersonPointer
./ex01.go:89:13: inlining call to fmt.Println
./ex01.go:72:17: leaking param: firstName to result ~r0 level=0
./ex01.go:72:35: leaking param: lastName to result ~r0 level=0
./ex01.go:76:24: leaking param: firstName
./ex01.go:76:42: leaking param: lastName
./ex01.go:77:9: &Person{...} escapes to heap
./ex01.go:83:25: &Person{...} does not escape
./ex01.go:88:13: ... argument does not escape
./ex01.go:88:24: ~r0 escapes to heap
./ex01.go:89:13: ... argument does not escape
./ex01.go:89:31: &Person{...} escapes to heap
```

Things the compiler output tells me that i didn't guess:

- The returned person escapes when it is passed to `fmt.Println` because `fmt.Println`
  takes an `interface{}` which requires heap allocation
- String data in fields like firstName and lastName escape the stack

*/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	// These DON'T escape
	p1 := MakePerson("John", "Doe", 25)
	p2 := MakePersonPointer("Betty", "Boop", 19)
	_ = p1
	_ = p2

	// These DO escape (move to heap) - because of fmt.Println
	fmt.Println(MakePerson("John", "Doe", 25))
	fmt.Println(MakePersonPointer("Betty", "Boop", 19))
}
