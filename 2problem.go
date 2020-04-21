package main

import (
	"fmt"
)

//need to take a slice and return with product of other indexes
//example if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6]

func main() {
	z1 := Product([]int{1, 2, 3, 4, 5})
	fmt.Println(z1)

	z2 := Product([]int{1, 43, 64, 2, 5, 67, 334, 23})
	fmt.Println(z2)

	z3 := Product([]int{1, 2, -1, 3, -2})
	fmt.Println(z3)
}

func Product(list []int) []int {
	res := 1
	for _, val := range list {
		res = res * val
	}

	for i, val := range list {
		list[i] = res / val
	}

	return list
}
