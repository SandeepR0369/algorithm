package main

import (
	"fmt"
)

//Given an array of numbers and a number k, determine if there are three entries in the array which add up to the specified number k. For example, given [20, 303, 3, 4, 25] and k = 49, return true as 20 + 4 + 25 = 49.

func main(){

	fmt.Println(numbers([]int{20, 303, 3, 4, 25}, 49))
}
/*

func numbers(nums []int, k int)bool{

//	var a, b, c int

var b, c int

//	a = 0
//	b = 0
//	c = 0
//	k = a + b + c

	var seclist, thirdlist []int

	for i := 0; i < len(nums); i++{

		if nums[i] > k || nums[i] == k{
			fmt.Println("3", b, nums[i])
			continue
		} 
		fmt.Println("1", nums[i])
			b = k - nums[i]
			//fmt.Println("2", b, nums[i])	

			seclist = append(seclist, b)
		
		for j := 0; j < len(seclist); j++{

			if seclist[j] > b || seclist[j] == b{
				continue
			} else {
				c = b - seclist[j]
				thirdlist = append(thirdlist, c)
			}
	}

	for l := 0; l < len(thirdlist); l++{
		if thirdlist[l] == c {
			return true
		}
	}
	
	}

	return false
}

*/
func numbers(nums []int, k int)bool{

	//	var a, b, c int
	
	var b, c int
	
	//	a = 0
	//	b = 0
	//	c = 0
	//	k = a + b + c
	
		var seclist, thirdlist []int
	
		for i := 0; i < len(nums); i++{
	
			if nums[i] > k || nums[i] == k{
				//fmt.Println("3", b, nums[i])
				continue
			} 
			//fmt.Println("1", nums[i])
				b = k - nums[i]
			fmt.Println("2", b, nums[i])	
	
				seclist = append(seclist, b)	
		}

		//for j := 0; j < len(seclist); j++{
			for j != i {
			if seclist[j] > b || seclist[j] == b{
				continue
			} else {
				c = b - seclist[j]
				thirdlist = append(thirdlist, c)
			}
	}

	for l := 0; l < len(thirdlist); l++{
		if thirdlist[l] == c {
			return true
		}
	}
	
		return false
	}


	func numbers(nums []int, k int)bool{

		//	var a, b, c int
		
		var b, c int
		
		//	a = 0
		//	b = 0
		//	c = 0
		//	k = a + b + c
		
			var seclist, thirdlist []int
		
			for i := 0; i < len(nums); i++{
		
				if nums[i] > k || nums[i] == k{
					//fmt.Println("3", b, nums[i])
					continue
				} 
				//fmt.Println("1", nums[i])
					b = k - nums[i]
				fmt.Println("2", b, nums[i])	
		
					seclist = append(seclist, b)	
			}
	
			//for j := 0; j < len(seclist); j++{
				for j != i {
				if nums[j] > b || nums[j] == b{
					continue
				} else {
					c = b - nums[j]
					thirdlist = append(thirdlist, c)
				}
		}
	
		for l := 0; l < len(thirdlist); l++{
			if thirdlist[l] == c {
				return true
			}
		}
		
			return false
		}	