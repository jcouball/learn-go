# Learn Go

## Example Code

Each chapter lives under `chNN/`, and every topic or exercise gets its own
folder with a small `package main`. To add a new example:

1. Create the chapter/topic directory, e.g. `ch04/interfaces`.
2. Add a `.go` file (any name works) with `package main` and a `main()` entry
     point.
3. Run `go run ./ch04/interfaces` or use the Makefile targets below.

The project Makefile auto-discovers every runnable package and exposes a few
helpers:

- `make list` &rarr; shows the safe target name (`ch03-slices`) and the package
    path (`./ch03/slices`).
- `make run-ch03-slices` &rarr; runs a specific topic/exercise; use the name
    from `make list`.
- `make build-all` / `make run-all` &rarr; build or run every example.
- `make ci` &rarr; runs `fmt`, `vet`, `test`, and `build-all`—use this before
    committing to ensure everything compiles.

## Chapter 1: Setting Up Your Environment

## Chapter 2: Predeclared Types and Declarations

Go defines predeclared types for booleans, numbers, strings, and a few built-in
interfaces like `error`, `any`, and `comparable`.

### Zero Value

The "zero value" is the default value given to any variable that is declared but not
assigned. Zero values for each type are:

- `bool`: `false`
- Numeric types default to `0`; complex numbers use `0+0i`
- `string`: `""` (empty string)
- `error`: `nil`
- `any`: `nil`
- `comparable`: inherits the zero value of the underlying type

### Literals

Types of literals:

- numeric: integers, floats, runes (single Unicode code points), and complex numbers
- string: interpreted (delimited with double quotes) and raw (delimited with
  backticks)
- boolean: `true` & `false`
- composite: arrays, slices, structs, maps, and interfaces
- function: `func(x int) int { return x * x }`

Interpreted string literals can not include unescaped backslashes, unescaped
newlines, or unescaped double quotes.

### Types

The types supported by Go are:

- Boolean: `bool`, zero value: false,
- Numeric
  - Integers: `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`,
    `uint64`, `byte` is an alias for `uint8`. `rune` is an alias for `int32`
  - Platform specific integers: `int`, `uint`, and `uintptr`
  - Floats: `float32`, `float64`
  - Complex: `complex64`, `complex128`
- String: `string`
- Error interface: `error`
- `any` (alias for `interface{}`)
- `comparable` (constraint for types supporting `==`/`!=`)

Strings in Go are immutable.

In Go, you must use explicit type conversion when variable types do not match,
including when only their numeric widths differ (e.g. assigning an int32 to an int64
requires explicit type conversion).

Literals are untyped.

For instance, boolean literals are untyped but only resolve to bool, whereas numeric
literals can become any compatible numeric type.

### Variables

#### Declaring with `var`

```go
var x int = 10
```

Omitting the type means the variable gets the literal's default type (`x` is of type
`int`):

```go
var x = 10
```

Omitting the value means to use the zero value (`name` is an empty string):

```go
var name string
```

Declare multiple variables (of the same type) at once. If the explicit type (`int`)
is left off, the type of each variable is inferred by each value. In this form, you
can only specify a single explicit type and it applies to all the variables being
declared.

```go
var x, y int = 10, 20
var name, age = "James", 58
var name string, age int = "James", 58 // INVALID: only 1 explicit type can be given
var name, age int = "James", 58 // INVALID: can not initialize an `int` with a `string`
```

Declare multiple variables (of different types) at once:

```go
var (
    firstName string = "James"
    lastName = "Couball" // Type inferred
    age int = 58
    minScore, maxScore = 90, 95 // Multiple values, type inferred
    father, mother string // Multiple values set to the zero value
)
```

#### Declaring with `:=`

Within a function, `:=` is a shorthand version of `var` that lets you declare and
initialize variables with type inference.

Declaring a single variable:

```go
var name = "James" // Using var
x := "James" // Equivalent using :=
```

Declaring multiple variables:

```go
var name, age = "James", 58 // Using var
name, age := "James", 58 // Equivalent using :=
```

You should only declare multiple variables on the same line when assigning multiple
values returned from a function or the comma ok idiom.

### Constants

`const`  is a way to declare a value as immutable. It is used just like `var`.

Constants give names to literals. They can only hold values that the compiler can
figure out at compile time. Go does not provide a way to specify that a value
calculated at runtime is immutable.

There are no immutable arrays, slices, maps, or structs and no way to declare that a
field in a struct is immutable.

Constants can be typed or untyped. An untyped constant works exactly like a literal.

### Naming Variables and Constants

Variable and constant names must start with an underscore or letter and can contain
underscores, letters, and digits. Here, letter and digit means any Unicode character
that is considered a letter or digit.

Idiomatic Go uses camel case for both variable and constant names.

## Chapter 3: Composite Types

### Arrays

Arrays are rarely used directly in Go. They have many limitations. Unless you know
the exact length you need ahead of time, use slices instead.

Ways to declare an array:

```go
// Declare the size of the array and type of each element
var x [3]int // elements are initialized to zero value
// With non-zero values
var x [3]int{10, 20, 30}
// Sparse array
var x [12]int{1, 5: 4, 6, 10: 100, 15} // unspecified elements are initialized to the zero value
// Infer the size by the number of elements
var x [...]int{10, 20, 30}
```

Arrays can be compared with `==` and `!=`. Arrays are equal if they are the same
length and contain the same values in the same order.

You can declare multidimensional arrays by stacking dimensions. The following
statement declares an array of length 2 whose element type is an array of ints of
length 3:

```go
var x [2][3]int
```

Reading past the end of an array or using a negative index causes a panic.

The built-in function `len` takes in an array and returns its length.

Since an array's *length* is part of its type and all types must be known at compile
time, you cannot use a variable to specify the size of an array.

This also means that you can't write a function that works with arrays of any size
and you can't assign arrays of different sizes to the same variable.

### Slices

A slice is a data structure that holds a sequence of values that can grow as needed.

The length of a slice is not part of its type.

When declaring a slice, you do not specify the size of the slice:

```go
// a slice initialially containing 3 integers
var x = []int{10, 20, 30}
// Multidimensional slice
var x [][]int
```

When declaring a slice, you want to minimize the number of times the slice needs to
grow.

If it isn't expected to grow at all, use a `var` declaration with no assigned value
to create a `nil` slice:

```go
var data []int
```

If you have some starting values or if a slice's values aren't going to change, then
use a slice literal:

```go
data := []string{"first", "second", "third"}
```

If you have a good idea of how large your slice needs to be, but don't know what
those values will be when you are writing the program, use `make`.

- If you are using a slice as a buffer, then specify a non-zero length
- If you are sure you know the exact size you want, you can specify the length and
  index into the slice to set the values
- In other situations, use `make` with a zero length and a specified capacity

```go
length := 10
capacity := 100
data := make([]string, length, capacity)
fmt.Println(data) // [          ]
```

You read and write slices using bracket syntax just like arrays. Reading past the end
of the slice or with a negative index causes a panic.

```go
x[0] = 10
fmt.Println(x[0])
```

The zero value for a slice is `nil`. A `nil` slice contains nothing.

You cannot use `==` or `!=` to compare slices. Use `slices.Equal` or
`slices.EqualFunc` instead.

#### `len` (slices)

The built-in function `len` takes a slice or array and returns its length. When given
`nil`, it returns `0`.

`len` on an array is compile-time constant (the length of the array) while on a slice
it’s computed at runtime.

#### `append`

`append` is used to grow slices:

```go
var x = []int{1, 2, 3}
x = append(x, 4, 5, 6)
fmt.Println(x) // Output: [1 2 3 4 5 6]
```

`append` can also concatenate two slices:

```go
var x = []int{1, 2, 3}
var y = []int{4, 5, 6}
x = append(x, y...)
fmt.Println(x) // Output: [1 2 3 4 5 6]
```

`append` always increases the length of a slice.

#### `cap`

The built-in function `cap` returns the current capacity of a slice. This is the size
a slice can grow to without having to reallocate the slice.

When `cap` is given an array, it returns the length of the array.

#### `make`

The built-in function `make` allows you to specify the type, length, and capacity of
a slice:

```go
// A slice with 5 elements each with the zero value
x := make([]int, 5)
// A slice with 0 elements and a capacity of 10
x := make([]int, 0, 10)
x = append(x, 1, 2, 3, 4) // slice would have length of 4
```

#### `clear` (slices)

The built-in function `clear` sets all of the slice's elements to their zero value.
The length of the slice remains unchanged:

```go
s := []string{"first", "second", "third"}
fmt.Println(s, len(s))
clear(s)
fmt.Println(s, len(s))
```

#### Slice Expressions

A slice expression creates a slice from a slice or an array.

A slice expression is a expression after an array or slice name enclosed in brackets:
`x[start:end]`. `start` is the position in `x` to start the slice. If omitted it
defaults to `0`. `end` is one past the last position to include. `start` and `end`
must be integers.

Examples:

```go
x := []string{"first", "second", "third", "forth"}
fmt.Println("x[:2]:", x[:2])   // Output: ["first", "second"]
fmt.Println("x[1:]:", x[1:])   // Output: ["second", "third", "forth"]
fmt.Println("x[1:3]:", x[1:3]) // Output: ["second", "third"]
fmt.Println("x[:]:", x[:]) // Output: ["first", "second", "third", "forth"]
```

When you make a slice from a slice, you are not making a copy. Instead, you now have
two variables that are sharing memory. Changes to an element in a slice affect all
slices that share that element.

```go
x := []string{"first", "second", "third", "forth"}
y := x[1:3] // ["second", "third"]
y[0] = "2nd"
x[2] = "3rd"
fmt.Println("x:", x) // Output: ["first", "2nd", "3rd", "forth"]
fmt.Println("y:", y) // Output: ["2nd", "3rd"]
```

Avoid modifying slices after they have been sliced or if they were produced by
slicing. Use a three-part slice expression (aka full slice expression) to prevent
append from sharing capacity between slices.

```go
x := []string{"first", "second", "third", "forth"}
y := x[1:3:3] // ["second", "third"]
y = append(y, "4th")
fmt.Println("x:", x) // Output: ["first", "second", "third", "forth"]
fmt.Println("y:", y) // Output: ["second", "third", "4th"]
```

#### `copy`

The built-in function `copy` creates a slice that is independent of the original
slice.

```go
x := []string{"first", "second", "third", "forth"}
y := make([]string, 2)
copy(y, x[1:3]) // y is ["second", "third"]
y[0] = "2nd" // y is altered without altering x
x[2] = "3rd"
fmt.Println("x:", x) // Output: ["first", "second", "3rd", "forth"]
fmt.Println("y:", y) // Output: ["2nd", "third"]
```

Copy copies as many values as it can from the source to destination, limited by
whichever slice length is smaller. The number of elements copied is returned. The
capacity of either slice does not matter.

You can copy between two slices that cover overlapping sections of an underlying
slice:

```go
x := []string{"first", "second", "third", "forth"}
copy(x[1:], x[0:2])
fmt.Println("x:", x) // Output: ["first", "first", "second", "forth"]
```

#### Converting between Arrays and Slices

To convert an array to a slice, use the `[:]` slice expression on the array:

```go
x := [4]string{"first", "second", "third", "forth"}
y := x[:] // y is a slice
```

To convert a slice to an array, use a type conversion:

```go
x := []string{"first", "second", "third", "forth"}
y := [4]string(x) // ["first", "second", "third", "forth"]
```

The size of the array must be known at compile time.

The size of the array may be smaller than the length of the slice. In this case, the
array will contain only the given number of elements from the slice.

```go
x := []string{"first", "second", "third", "forth"}
z := [2]string(x) // ["first", "second"]
```

A runtime panic will occur if the size of the array is larger than the length of the
slice.

```go
x := []string{"first", "second", "third", "forth"}
y := [5]string(x) // Panics
```

### Strings

In Go, a string is implemented as a series of bytes instead of a series of UTF-8 code
points.

This means that using `len`, index expressions and slice expressions can be
problematic when the string contains multi-byte code points.

Converting an int to a string treats the int like a UTF-8 code point.

```go
var x int = 65
var y string(x)
fmt.Println(y) // Outputs "A" and not "65"
```

A string can be converted back and forth to a slice of bytes or a slice of runes.

```go
var s string = "Hello, 🌞"
var bs []byte = []byte(s)
var rs []rune = []rune(s)
fmt.Println(bs) // Output: [72 101 108 108 111 44 32 240 159 140 158]
fmt.Println(rs) // Output: [72 101 108 108 111 44 32 127774]
```

Rather than use the slice and index expressions with strings, you should extract
substrings and code points from strings using the functions in the `strings` and
`unicode/utf8` packages in the standard library.

### Maps

A map is a hash map implemented as a hash table under the hood.

The map type is written: `map[keyType]valueType`

The zero value for a map is `nil`. A `nil` map has a length of `0`. Attempting to
read a `nil` map always returns the zero value for the map’s value type. However,
attempting to write to a `nil` map variable causes a panic.

Here are a few ways to declare maps:

```go
// Create a map variable that’s set to its zero value
var map1 map[string]int // returns nil
// Create an empty map using := with a map literal
map2 := map[string]int{} // returns an empty map with len 0
// Create a non-empty map using := with a map literal
map3 := map[string][]string {
    "beatles": []string{"george", "john", "paul", "ringo"},
    "stones": []string{"mick", "keith", "bryan", "bill", "charlie"}
}

fmt.Println(map1) // Output: map[]
fmt.Println(map2) // Output: map[]
fmt.Println(map3) // Output: map[beatles:[george john paul ringo] stones:[mick keith bryan bill charlie]]
```

If you know how many key-value pairs you plan to insert into a map, you can use make
to create a map with a specific initial size:

```go
ages := make(map[int][]string, 10)
```

Maps created with `make` still have a length of `0`, and they can grow past the
initially specified size.

Maps automatically grow as you add key-value pairs to them.

Maps are not comparable. You can check if they are equal to nil, but you cannot check
if two maps have identical keys and values using == or differ using !=.

#### `len` (maps)

Returns the number of key-value pairs in a map.

```go
myMap := map[string]int{"James": 58, "Jake": 27}
fmt.Println(len(myMap)) // Output: 2
```

#### Reading and Writing

Assign a value to a map key by putting the name within brackets and using = to
specify the value. Read the value assigned to a map key by putting the key within
brackets.

```go
ages := map[string]int{}
ages["John"] = 58
ages["Mary"] = 32
fmt.Println("John's age:", ages["John"]) // Output: John's age: 58
fmt.Println("Mary's age:", ages["Mary"]) // Output: Mary's age: 32
```

Reading the value of a key that isn't in the map will return the zero value for the
map's value type. You can tell if the key exists using "the comma ok idiom."

```go
ages := map[string]int{}
ages["John"] = 58
age, ok := ages["John"]
fmt.Println(age, ok) // Output: 58 true
age, ok = ages["Bob"]
fmt.Println(age, ok) // Output: 0 false
```

#### `delete`

The built-in function `delete` removes key-value pairs from a map:

```go
m := map[string]int{"James": 58, "Mary": 32}
delete(m, "Mary")
fmt.Println(m) // Output: map[James:58]
```

If the key isn't present in the map or if the map is `nil`, nothing happens. The
`delete` function does not return a value.

#### `clear` (maps)

The built-in function `clear` sets the map's length to zero:

```go
m := map[string]int{"James": 58, "Mary": 32}
fmt.Println(m, len(m)) // Output: map[James:58,Mary:32] 2
clear(m)
fmt.Println(m, len(m)) // Output: map[] 0
```

#### Comparing Maps

Use `maps.Equal` or `maps.EqualFunc` see if two different maps contain the same
key-value pairs:

```go
m1 := map[string]int{"James": 58, "Mary": 32}
m2 := map[string]int{"Mary": 32, "James": 58}
m3 := map[string]int{"James": 58, "Bob": 27}
fmt.Println("maps.Equal(m1, m2):", maps.Equal(m1, m2)) // true
fmt.Println("maps.Equal(m1, m3):", maps.Equal(m1, m3)) // false
```

#### Maps as Sets

Go doesn't include a set type, but you can use a map to simulate some of its
features. Store the set value type in the key of the `map` and store a `bool` for the
value.

### Structs

In Go, a `struct` is a composite type that groups together zero or more named fields,
each with its own type, into a single unit. It’s the language’s way to define custom
data records: `type Person struct { Name string; Age int }`. Instances can then hold
related values under one variable, with fields accessed via dot notation.

A struct definition uses the `type` keyword, a name, the `struct` keyword, and braces
containing the fields. Inside the braces, list each field on its own line as
`FieldName FieldType`.

```go
type person struct {
    name string
    age  int
    pet  string
}
```

Structs defined in a function can only be used within that function.

Once a struct is declared, you can declare variables of that type:

```go
var fred person
```

Struct variables with no value assigned gets the zero value for the struct type. A
zero value struct has every field set to the field's zero value.

A struct literal can be assigned to a variable as well:

```go
bob := person{}
```

Field values can be given in a struct literal as a comma delimited list of fields or
in map literal style:

```go
type person struct {
    name string
    age  int
}
james := person{"James", 58}
bob := person{name: "Bob", age: 32}
```

When giving values using the comma delimited list format, a value for every field in
the struct must be given in the order they were declared in the struct definition.

When using the keyed (map-style) struct literal, specify each field as `FieldName:
value`. You may list fields in any order and omit ones you want left at their zero
value. A multi-line struct literal must end every line, even the last field,
with a trailing comma.

A field in a struct is accessible with dot notation:

```go
type person struct {
    name string
    age  int
}
james := person{"James", 58}
fmt.Println("Name: ", james.name) // Output: Name: James
fmt.Println("Age: ", james.age)   // Output: Age: 58
```

#### Anonymous Structs

You can declare that a variable implements a struct type without first giving the struct type a name. The type is unique to that declaration and can't be reused elsewhere.

```go
var person struct {
    name string
    age  int
}

person.name = "James"
person.age = 58

fmt.Println(person) // Output: {James 58}
```

You can also assign values with a struct literal:

```go
person := struct {
    name string
    age  int
}{
    name: "Mary",
    age: 32,
}

fmt.Println(person)
```

#### Comparing and Converting Structs

Structs that are entirely composed of comparable types are comparable. Those with
slice, map, function, and channel fields are not.

You can't compare structs of different types.

Go allows type conversion of one struct type to another if the fields of both structs
have the same names, order, and types.

Two different types of anonymous structs can be compared without a type conversion if
the fields of both structs have the same names, order, and type. You can also assign
between named and anonymous struct types if the fields of both structs have same
names, order, and types.

```go
type firstPerson struct {
    name string
    age  int
}
f := firstPerson{name: "James", age: 58}
var g struct {
    name string
    age  int
}

// The following compiles
g = f
fmt.Println(f == g) // Output: true
```
