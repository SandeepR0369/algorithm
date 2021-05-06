package main

import "fmt"

/*
Write a function which takes an array and return majority element (if it exists) and error (return error if n majority element exist). A majority element in an array A[] of size n is an element that appears more than n/2 times (and hence there is at most one such element).
 {3, 3, 4, 2, 4, 4, 2, 4, 4}
 {3, 3, 4, 2, 4, 4, 2, 4}"

*/

func main() {
	fmt.Println(major([]int{3, 3, 4, 2, 4, 4, 2, 4, 4}))
	//major([]int{3, 3, 4, 2, 4, 4, 2, 4, 4})
}

func major(a []int) int {
	//var count int
	//fmt.Println(count)

	ele, count := a[0], 0
	for i := 0; i < len(a); i++ {
		if count == 0 {
			ele, count = a[i], 1
		} else {
			if a[i] == ele {
				count++
			} else {
				count--
			}

		}
	}
	return ele
}

