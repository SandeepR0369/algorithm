package main

import (
	"fmt"
	"sort"
	"reflect"
)

func main() {
	v := LargestProduct([]int{-10, -10, 5, 2, -7})
//	v1 := LargestProduct([]int{17, 20, -2, 31, 52, 67, 24, 63, 05})
//	v2 := LargestProduct([]int{-1, -2, 0, 1, 2, 7, 4, 3, 5})
	fmt.Println(v)
//	fmt.Println(v1)
//	fmt.Println(v2)
}

//Given a list of integers, return the largest product that can be made by multiplying any three integers.
//For example, if the list is [-10, -10, 5, 2], we should return 500, since that's -10 * -10 * 5.

func LargestProduct(nums []int) int {
	var end []int
	var product1 int

	product := 1

	for k, _ := range nums {
		if nums[k] != 0 {
			product = product * nums[k]
	//	}else if nums[k] == 0 {
	//		product = product * (nums[k] - (nums[k] == 0))
		}
		//product = product * nums[k]
		fmt.Println("1VAER",reflect.TypeOf(nums[k]))
		fmt.Println("2VAER",reflect.TypeOf(nums[k] == 0))
	}
	

	for i := 0; i <= len(nums)-1; i++ {
		/*	if nums[i] == 0{
			product1 = product / nums[i]
		}*/

		if nums[i] != 0 {
	//		fmt.Println("E", end)
			product1 = product / nums[i]
			end = append(end, product1)
	//		fmt.Println("E1", end)
		} else if nums[i] == 0 {
	//		fmt.Println("E2", end)
			product1 = product / 1
			end = append(end, product1)
		}

		sort.Ints(end)
	}
	return end[len(end)-1]
}
