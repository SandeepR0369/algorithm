package main

import (
	"fmt"
	"sort"
)

func main() {
	y := Missing([]int{-1, 0, -2, 1, 2, 7, 4, 3, 5})
	z := Missing([]int{-3, -2, -1, 33, 65, 67})

	fmt.Println(y)
	fmt.Println(z)
	z1 := Missing([]int{2, 3, 4, 11, 13})
	fmt.Println(z1)

	z2 := Missing([]int{3, 0, 3, 1, 4})
	fmt.Println(z2)

	z3 := Missing([]int{3, 0, 2, 1, 4})
	fmt.Println(z3)

}

//find the first missing positive integer in linear time and constant space
//example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3
//check what is first number and then next number isn't increment it should return increment

func Missing(numbers []int) int {
	var j []int
	sort.Ints(numbers)
	for i, val := range numbers {
		if numbers[i] < 0 {
			continue
		} else if numbers[i] >= 0 {
			j = append(j, val)
		}
	}

	for _, _ = range j {
		for m := 0; m < len(j)-1; m++ {
			if j[m+1]-j[m] == 1 {
				continue
			} else if j[m+1]-j[m] > 1 {
				return j[m] + 1
			}
		}
		if j[len(j)-1]-j[len(j)-1-1] == 1 {
			return j[len(j)-1] + 1
		}
	}
	return 0
}
