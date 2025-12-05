package main

import (
	"fmt"
	"os"
	"time"
)

/*
Chapter 6, Exercise 3

Write a program that builds a []Person with 10,000,000 entries (they could all be the
same names and ages). See how long it takes to run.

Change the value of GOGC and see how that affects the time it takes for the program
to complete. Set the environment variable GODEBUG=gctrace=1 to see when garbage
collections happen and see how changing GOGC changes the number of garbage
collections.

What happens if you create the slice with a capacity of 10,000,000?

*/

/*
Solution

Running with default GOGC takes about 0.5s:

```console
[james] ex03 $ go run ex03.go
Execution time: 587.328625ms
```

Running with GOGC=20 takes about 0.75s and does many more garbage collections:

```console
[james] ex03 $ GOGC=20 go run ex03.go
Execution time: 726.313042ms
```

GOGC=100 (default): 36 garbage collections triggered

GOGC=20: 71 garbage collections triggered

Setting GOGC=20 nearly doubled the number of GC cycles (71 vs 36). This makes sense
because GOGC controls the percentage of heap growth before triggering a new GC cycle.
With GOGC=20, the garbage collector runs when the heap grows by only 20% (instead of
100% with the default), resulting in more frequent but smaller GC pauses.

Preallocating the slice with a capacity of 10,000,000 reduces the execution time to
about 70ms, regardless of GOGC setting, and results in fewer garbage collections (16
vs. 36).
*/

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Fprintf(os.Stderr, "\nExecution time: %v\n", time.Since(start))
	}()

	people := make([]Person, 0, 10_000_000)

	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{"John", "Doe", 30})
	}

	_ = people
}
