package main

import (
	"errors"
	"fmt"
	"log"
)

/*
type BST struct {
	Value int

	Left  *BST
	Right *BST
}*/

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}
/*
func (tree *BST) InsertValue(target int) int {
	// Write your code here.
	return -1
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	return -1
}*/

func main() {
	//Insert("1", 3)
	values := []string{"d", "b", "c", "e", "a"}
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	tree := &Node{}
	for i := 0; i < len(values); i++ {
		err := tree.Insert(values[i], data[i])
		if err != nil {
			log.Fatal("Error inserting value '", values[i], "': ", err)
		}
	}
	fmt.Println("Tree", tree)
}

func (n *Node) Insert(value, data string) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {

	case value == n.Value:
		return nil

	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}
