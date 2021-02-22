package main

import (
	"fmt"
)

func main() {
	fmt.Println(sorrt([]int{2, 5, 6, 1, 3}))
}

func sorrt(a []int) []int {

	for i := 0; i < len(a); i++ {
		swap := false
		for j := 0; j < len(a)-1; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swap = true
			} else {
				continue
			}
		}

		if !swap {
			break
		}
	}
	return a
}
