package main

import (
	"fmt"
	"sort"
)

//Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
//For example, given [100, 4, 200, 1, 3, 2], the longest consecutive element sequence is [1, 2, 3, 4]. Return its length: 4.

func main(){

}

func longest(nums []int)int{
	 var new []int

	 sort.Ints(nums)

	 for k, v := range nums(
		 for i := 0; i <=len(nums)-1; i++{
			 if nums[i+1]-nums[i]==1{
				new = new + 1
			 }
			 
		 }
	 )
}