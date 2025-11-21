package main

import (
	"fmt"
	"maps"
)

func declare_nil_map() {
	var m map[string]int
	fmt.Println(m == nil) // Output: true
}

func declare_empty_map() {
	m := map[string]int{}
	fmt.Println(m == nil) // Output: false
	fmt.Println(m)        // Output: map[]
}

func declare_non_empty_map() {
	m := map[string][]string{
		"beatles": {"george", "john", "paul", "ringo"},
		"stones":  {"mick", "keith", "bryan", "bill", "charlie"},
	}
	fmt.Println(m)
}

func map_len() {
	myMap := map[string]int{"James": 58, "Jake": 27}
	fmt.Println(len(myMap)) // Output: 2
}

func read_and_write_map() {
	ages := map[string]int{}
	ages["John"] = 58
	ages["Mary"] = 32
	fmt.Println("John's age:", ages["John"]) // Output: John's age: 58
	fmt.Println("Mary's age:", ages["Mary"]) // Output: Mary's age: 32
}

func read_comma_ok() {
	ages := map[string]int{}
	ages["John"] = 58
	age, ok := ages["John"]
	fmt.Println(age, ok) // Output: 58 true
	age, ok = ages["Bob"]
	fmt.Println(age, ok) // Output: 0 false
}

func map_delete() {
	m := map[string]int{"James": 58, "Mary": 32}
	delete(m, "Mary")
	fmt.Println(m) // Output: map[James:58]
}

func map_clear() {
	m := map[string]int{"James": 58, "Mary": 32}
	fmt.Println(m, len(m)) // Output: map[James:58 Mary:32] 2
	clear(m)
	fmt.Println(m, len(m)) // Output: map[] 0
}

func compare_maps() {
	m1 := map[string]int{"James": 58, "Mary": 32}
	m2 := map[string]int{"Mary": 32, "James": 58}
	m3 := map[string]int{"James": 58, "Bob": 27}
	fmt.Println("maps.Equal(m1, m2):", maps.Equal(m1, m2)) // true
	fmt.Println("maps.Equal(m1, m3):", maps.Equal(m1, m3)) // false
}

func main() {
	declare_nil_map()
	declare_empty_map()
	declare_non_empty_map()
	map_len()
	read_and_write_map()
	read_comma_ok()
	map_delete()
	map_clear()
	compare_maps()
}
