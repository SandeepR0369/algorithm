package main

import "fmt"

func main() {
	input := []int{1, 6, 8, 5, 9, 4, 7, 2}
	output := summation(input)
	fmt.Println(output) // Output: [1 14 18 9]
}

func summation(arr []int) []int {
	abc := []int{}
	sum := 0
	count := 1
	j := 1
	abc = append(abc, arr[0])
	for j < len(arr) {
		for i := count; i < count+1; i++ {
			sum += arr[i]
			count += 1
		}
		abc = append(abc, sum)
	}
	return abc
}
