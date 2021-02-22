package main

import "fmt"

func main() {
	fmt.Println(IsMonotonic([]int{-1, -5, -10, -1100, -900, -1101, -1102, -9001}))
}

func IsMonotonic(array []int) bool {
	// Write your code here.
	if len(array) == 0 || len(array) == 1 {
		return true
	}

	for i := 1; i < len(array); i++ {

		if array[i] > array[i-1] || array[i] == array[i-1] {
			continue
		} else {
			return false
		}

		if array[i] < array[i-1] || array[i] == array[i-1] {
			continue
		} else {
			return false
		}
	}

	return true
}
