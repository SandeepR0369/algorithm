package main

import (
	"fmt"
)

//Given an array nums of integers, return how many of them contain an even number of digits.
//Example 1:
//Input: nums = [12,345,2,6,7896]  Output: 2  {12 -> 2digits, 345 -> 3digits, 2, 6 -> 1digit, 7896 -> 4digits}

func main() {
	fmt.Println(EvenDigits([]int{12, 345, 2, 6, 7896, 1771}))
	fmt.Println(NumEvenDigits([]int{12, 345, 2, 6, 7896, 1771}))

}

func EvenDigits(nums []int) []bool {
	var out []bool
	for i := 0; i < len(nums); i++ {
		if IsEven(Split(nums[i])) == true {
			out = append(out, true)
		} else {
			out = append(out, false)
		}
	}
	return out
}

func NumEvenDigits(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if IsEven(Split(nums[i])) == true {
			count += 1
		} else {
			continue
		}
	}
	return count
}

func IsEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func Split(m int) int {
	count := 0
	for m != 0 {	
		m = m / 10
		count += 1
	}
	return count
}
