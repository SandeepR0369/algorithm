package main

import (
	"fmt"
	"math"
)

type Input struct {
	root *BST
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

//need to find the closest value
//check the root value if it matches, if not store the difference
//compare the left and right side values move that side

func (t *Input) Insert(data int) {
	if t.root == nil {
		t.root = &BST{Value: data}
	} else {
		t.root.insert(data)
	}
}

func (tree *BST) FindClosestValue(target int) int {
	var diff []int
	var d1, d2, d3 int
	mapdiff := make(map[int]int)

	if tree.Value == target {
		return tree.Value
	} else {
		//diff = append(diff, int(math.Abs(tree.Value-target)))

		d1 = int(math.Abs(float64(tree.Value) - float64(target)))
		mapdiff[d1] = tree.Value
	}

	/*if tree.Right.Value == target {
		return tree.Right.Value
	}*/

	//LEFT
	if tree.Value < target {
		if tree.Right.Value == target {
			return tree.Right.Value
		} else {
			//diff = append(diff, int(math.Abs(tree.Right.Value-target)))

			d2 = int(math.Abs(float64(tree.Right.Value) - float64(target)))
			mapdiff[d2] = tree.Right.Value
		}

		if tree.Left.Value == target {
			return tree.Left.Value
		} else {
			//diff = append(diff, int(math.Abs(tree.Left.Value-target)))

			d3 = int(math.Abs(float64(tree.Left.Value) - float64(target)))
			mapdiff[d3] = tree.Left.Value
		}

	}

	//RIGHT
	if tree.Value > target {
		if tree.Right.Value == target {
			return tree.Right.Value
		} else {
			//diff = append(diff, int(math.Abs(tree.Right.Value-float64(target))))

			d2 = int(math.Abs(float64(tree.Right.Value) - float64(target)))
			mapdiff[d2] = tree.Right.Value
		}

		if tree.Left.Value == target {
			return tree.Left.Value
		} else {
			//diff = append(diff, int(math.Abs(tree.Left.Value-target)))

			d3 = int(math.Abs(float64(tree.Left.Value) - float64(target)))
			mapdiff[d3] = tree.Left.Value
		}
	}
	fmt.Println("MAPP", mapdiff)
	//sort.Ints()

	return -1
}

func main() {
	var t Input

	t.Insert(12)
	t.Insert(6)
	t.Insert(7)
	t.Insert(14)
	t.Insert(17)

	fmt.Println(FindClosestValue(10))

}

func diff(tree *BST, target int) {

}
