package main

import (
	"fmt"
	"sort"
)

func main() {
	a := high1([]int{51, 21, 108, 40, 7, 34})
	fmt.Println(a)

	b := high2([]int{109, 51, 9, 21, 108, 40, 7, 34})
	fmt.Println(b)

	c := low([]int{109, 51, 9, 21, 108, 40, 7, 34})
	fmt.Println(c)
}

func high1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func high2(nums []int) int {
	for m := 0; m < len(nums); m++ {
		swap := false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return nums[len(nums)-1]
}

func low(nums []int) int {
	for m := 0; m < len(nums); m++ {
		swap := false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return nums[0]
}
