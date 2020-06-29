package main

import (
	"fmt"
	"sort"
)

func main() {
	v := LargestProduct([]int{-10, -10, 5, 2,-7})
	v1 := LargestProduct([]int{17, 20, -2, 31, 52, 67, 24, 63, 05})
	v2 := LargestProduct([]int{-1, -2, 0, 1, 2, 7, 4, 3, 5})
	fmt.Println(v)
	fmt.Println(v1)
	fmt.Println(v2)
}

//Given a list of integers, return the largest product that can be made by multiplying any three integers.
//For example, if the list is [-10, -10, 5, 2], we should return 500, since that's -10 * -10 * 5.

func LargestProduct(nums []int)int {
	var end []int
	product := 1

	for k, _ := range nums {
		product = product * nums[k]
	}

	for i := 0; i <= len(nums)-1; i++{

		if nums[i] == 0{
			fmt.Println("This slice don't work due to zero")
			break
		}

		product1 := product / nums[i]
		end = append(end, product1)
		sort.Ints(end)
	}
	return end[len(end)-1]
}
