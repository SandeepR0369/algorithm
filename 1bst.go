package main

import (
	"fmt"
	"math"
)

type Tree struct {
	value      int
	left_tree  *Tree
	right_tree *Tree
}

func main() {
	var number int
	var choice string
	var tree Tree
	//tree := new(Tree)
	// var lefttree Tree
	// var righttree Tree
	// tree.left_tree = &lefttree
	// tree.right_tree = &righttree
	fmt.Println("original tree:", tree)
	for {
		fmt.Println("do you want to add number to tree, press y to add, press n to print the tree")
		fmt.Scan(&choice)
		switch choice {
		case "y":
			fmt.Println("enter any number to insert to tree:")
			fmt.Scan(&number)
			insert(&tree, number)
			fmt.Printf("%d added to tree\n", number)
		case "n":
			printtree(&tree)
			return
		default:
			fmt.Println("invalid choice, please choose y or n")
		}
	}
}
func insert(tree *Tree, number int) {
	fmt.Println("tree.value: ", tree.value)
	if tree.value == 0 {
		tree.value = number
		return
	}
	if number < tree.value {
		fmt.Println("going here value is less:", tree.value)
		if tree.left_tree != nil {
			tree = tree.left_tree
			insert(tree, number)
		} else {
			var lefttree Tree
			tree.left_tree = &lefttree
			tree = tree.left_tree
			insert(tree, number)
		}
	}
	if number > tree.value {
		fmt.Println("going here value is more:", tree.value)
		if tree.right_tree != nil {
			tree = tree.right_tree
			insert(tree, number)
		} else {
			var righttree Tree
			tree.right_tree = &righttree
			tree = tree.right_tree
			insert(tree, number)
		}
	}
	return
}

func printtree(tree *Tree) {
	if tree.left_tree != nil {
		fmt.Println("L")
		printtree(tree.left_tree)
	}
	fmt.Println(tree.value)
	if tree.right_tree != nil {
		fmt.Println("R")
		printtree(tree.right_tree)
	}
	return
}

func (tree *Tree) FindClosestValue(target int) int {
	//var diff []int
	var d1, d2, d3 int
	mapdiff := make(map[int]int)

	if (*tree).value == target {
		return (*tree).value
	} else {
		//diff = append(diff, int(math.Abs(tree.value-target)))

		d1 = int(math.Abs(float64((*tree).value) - float64(target)))
		mapdiff[d1] = (*tree).value
	}

	/*if tree.right_tree.value == target {
		return tree.right_tree.value
	}*/

	//LEFT
	if (*tree).value < target {
		if (*tree).right_tree.value == target {
			return (*tree).right_tree.value
		} else {
			//diff = append(diff, int(math.Abs(tree.right_tree.value-target)))

			d2 = int(math.Abs(float64((*tree).right_tree.value) - float64(target)))
			mapdiff[d2] = (*tree).right_tree.value
		}

		if (*tree).left_tree.value == target {
			return (*tree).left_tree.value
		} else {
			//diff = append(diff, int(math.Abs(tree.left_tree.value-target)))

			d3 = int(math.Abs(float64((*tree).left_tree.value) - float64(target)))
			mapdiff[d3] = (*tree).left_tree.value
		}

	}
	//RIGHT
	if (*tree).value > target {
		if (*tree).right_tree.value == target {
			return (*tree).right_tree.value
		} else {
			//diff = append(diff, int(math.Abs((*tree).right_tree.value-float64(target))))
			d2 = int(math.Abs(float64((*tree).right_tree.value) - float64(target)))
			mapdiff[d2] = (*tree).right_tree.value
		}

		if (*tree).left_tree.value == target {
			return (*tree).left_tree.value
		} else {
			//diff = append(diff, int(math.Abs(tree.left_tree.value-target)))

			d3 = int(math.Abs(float64((*tree).left_tree.value) - float64(target)))
			mapdiff[d3] = (*tree).left_tree.value
		}
	}

	fmt.Println("MAPP", mapdiff)
	//sort.Ints()
	return -1
}
