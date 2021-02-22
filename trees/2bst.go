package main

import (
//	"errors"
	"fmt"
//	"log"
)


type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func printNode(n *BST){
	fmt.Println("Value: ", n.Value)
}

func main(){
	x := &BST{
		Value: 6,
	}

	printNode(x)
}