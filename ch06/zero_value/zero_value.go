package main

import "fmt"

type Person struct {
	Name   string
	Age    *int // nil means age not provided, 0 means age is actually 0
	HasAge bool // alternative approach using boolean flag
}

func setDefault(age *int) {
	if age == nil {
		// Cannot do: *age = 18  // This would panic!
		// The pointer itself is passed by value, so reassigning it here will not help
		fmt.Println("Cannot set default on nil pointer")
	} else {
		*age = 18
		fmt.Println("Set age to 18")
	}
}

// GetAge demonstrates the comma ok idiom
func GetAge(p Person) (int, bool) {
	if p.Age != nil {
		return *p.Age, true
	}
	return 0, false
}

func main() {
	// Using nil pointer to indicate "no value"
	p1 := Person{Name: "Alice", Age: nil} // age not provided
	age := 0
	p2 := Person{Name: "Bob", Age: &age} // age is explicitly 0
	age25 := 25
	p3 := Person{Name: "Charlie", Age: &age25} // age is 25

	fmt.Printf("%s age pointer: %v\n", p1.Name, p1.Age)
	if p1.Age != nil {
		fmt.Printf("%s age value: %v\n", p1.Name, *p1.Age)
	}

	fmt.Printf("%s age pointer: %v\n", p2.Name, p2.Age)
	fmt.Printf("%s age value: %v\n", p2.Name, *p2.Age)

	fmt.Printf("%s age pointer: %v\n", p3.Name, p3.Age)
	fmt.Printf("%s age value: %v\n", p3.Name, *p3.Age)

	// Demonstrate setDefault behavior
	setDefault(p1.Age) // nil pointer
	setDefault(p2.Age) // pointer to 0
	fmt.Printf("After setDefault, Bob's age: %v\n", *p2.Age)

	// Using comma ok idiom
	if age, ok := GetAge(p1); ok {
		fmt.Printf("%s's age from GetAge: %v\n", p1.Name, age)
	} else {
		fmt.Printf("%s's age not set\n", p1.Name)
	}

	if age, ok := GetAge(p3); ok {
		fmt.Printf("%s's age from GetAge: %v\n", p3.Name, age)
	} else {
		fmt.Printf("%s's age not set\n", p3.Name)
	}
}
