package main

import (
	"fmt"
)

//Given a multiset of integers, return whether it can be partitioned into two subsets whose sums are the same.
//For example, given the multiset {15, 5, 20, 10, 35, 15, 10}, it would return true, since we can split it up into {15, 5, 10, 15, 10} and {20, 35}, which both add up to 55.
//Given the multiset {15, 5, 20, 10, 35}, it would return false, since we can't split it up into two subsets that add up to the same sum.

func main() {

	a := multiset([]int{1,5,5,11})
	fmt.Println("A", a)

}

func multiset(nums []int) bool {
//	var s1, s2 []int

//	for _, _ = range nums {
		for i := 0; i <= len(nums)-1; i++ {
			if nums[i] == sums(nums)-nums[i] {
				//fmt.Println(nums[i])
			}
		}
//	}
	return true
}

func sums(nums []int) int {
	//var nums []int
	sum := 0

	for i := 0; i <= len(nums)-1; i++ {
		sum = sum + nums[i]
		//fmt.Println("hh", sum)
	}
	return sum
}
