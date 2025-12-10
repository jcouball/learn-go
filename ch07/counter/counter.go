package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

// Receiver is a pointer so the total and lastUpdated fields can be mutated
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// Receiver is a value receiver since it only reads fields
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}
