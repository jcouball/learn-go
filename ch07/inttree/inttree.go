package main

import "fmt"

type IntTree struct {
	val         int
	left, right *IntTree
}

// This method works even when called on a nil *IntTree
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

// InOrder traversal for printing
func (it *IntTree) InOrder() {
	if it == nil {
		return
	}
	it.left.InOrder()
	fmt.Printf("%d ", it.val)
	it.right.InOrder()
}

func main() {
	var tree *IntTree
	// Inserting into nil tree
	tree = tree.Insert(5)
	tree = tree.Insert(3)
	tree = tree.Insert(7)
	tree = tree.Insert(1)
	tree = tree.Insert(9)

	fmt.Print("Tree values (in-order): ")
	tree.InOrder()
	fmt.Println()
}
