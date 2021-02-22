package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, v int) *Node {
	if root == nil {
		root = &Node{v, nil, nil}
	} else if v < root.value {
		root.left = insert(root.left, v)
	} else {
		root.right = insert(root.right, v)
	}

	return root
}

func traverse(root *Node) {
	if root == nil {
		return
	}

	traverse(root.left)
	traverse(root.right)
}

func main() {
	var root *Node
	const SIZE = 2000000
	var a [SIZE]int

	fmt.Printf("Generating random array with %v values...\n", SIZE)

	start := time.Now()

	for i := 0; i < SIZE; i++ {
		a[i] = rand.Intn(10000)
	}

	end := time.Since(start)

	fmt.Printf("Done. Took %s\n\n", end)
	fmt.Printf("Filling the tree with %v nodes...\n", SIZE)

	start = time.Now()

	for i := 0; i < SIZE; i++ {
		root = insert(root, a[i])
	}

	end = time.Since(start)

	fmt.Printf("Done. Took %s\n\n", end)
	fmt.Printf("Traversing all %v nodes in tree...\n", SIZE)

	start = time.Now()

	traverse(root)

	end = time.Since(start)

	fmt.Printf("Done. Took %s\n\n", end)
}