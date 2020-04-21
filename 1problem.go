package main

import (
	"fmt"
)

//Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
//For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

func main() {
	y, t := sum([]int{1, 43, 64, 2, 5, 67, 334, 23}, 69)
	fmt.Println(y, t)
	x := Isitcorrect([]int{1, 43, 64, 2, 5, 67, 334, 23}, 169)
	fmt.Println(x)
}

func sum(two []int, sum int) (int, int) {
	for i, val := range two {
		for j, val2 := range two {
			if i == j {
				continue
			} else if val+val2 == sum {
				return i, j
			}
		}
	}
	return -1, -1
}

func Isitcorrect(two []int, sum int) bool {
	for i, val := range two {
		for j, val2 := range two {
			if i == j {
				continue
			} else if val+val2 == sum {
				return true
			}
		}
	}
	return false
}
