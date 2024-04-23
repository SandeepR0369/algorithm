package main

import "fmt"

func main() {
	fmt.Println(duplicate([]int{1, 2, 3, 4, 1}))
}

func duplicate(a []int) bool {
	seen := make(map[int]bool)
	for _, v := range a {

		if seen[v] {
			return true
		}
		seen[v] = true
	}

	return false
}
