package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	//array := [][]int{{{{1}, {2}, {3}, {4}}, {{12}, {13}, {14}, {5}}, {{11}, {16}, {15}, {6}}, {{10}, {9}, {8}, {7}}}}
	array := [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}
	fmt.Println("A", array)
	fmt.Println(IsSquare(array))
	fmt.Println(SpiralTraverse(array))
}

func SpiralTraverse(array [][]int) []int {
	// Write your code here.
	var jaffa []int

	m := len(array)
	//n := len(array[0])

	if IsSquare(array) == false {
		return nil
	}

	for i := 0; i < m; i++ {
		//for j := 0; j < len(array[i]); j++ {
		for j := 2i && j <= m {	
		//if i
			jaffa = append(jaffa, array[i][j])	
		}

	//	for k := 


	}
	return jaffa
}


func IsSquare(array [][]int) bool {
	m := len(array)

	for i := 0; i < m; i++ {
		if m == len(array[i]) {
			continue
		} else {
			return false
		}
	}
	return true
}

