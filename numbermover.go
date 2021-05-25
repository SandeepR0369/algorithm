package main

import (
	"fmt"
)

func main() {
	fmt.Println(Numbermover([]int{6, 0, 8, 2, 3, 0, 4, 0, 1}, 0))
	fmt.Println(Numbermover([]int{6, 0, 8, 2, 3, 0, 4, 0, 1}, 6))
}

func Numbermover(sli []int, target int) []int {

	for j := 0; j < len(sli); j++ {
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] == target {
				sli[i], sli[i+1] = sli[i+1], sli[i]
			}
		}
	}
	return sli
}

