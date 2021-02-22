package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, playground")
	//fmt.Println(MinNumberOfCoinsForChange(97, []int{1, 5, 10}))
	fmt.Println(MinNumberOfCoinsForChange(3, []int{2, 1}))
}

func MinNumberOfCoinsForChange(n int, denoms []int) int {

	coins := 0

	sort.Ints(denoms)
	for i := len(denoms) - 1; i >= 0; i-- {

		for j := len(denoms) - 1; j >= 0; j-- {
			if n == denoms[j] {
				return 1
			}
			fmt.Println("N", n)
			n = n - denoms[j]
			fmt.Println("N1", n)
			coins += 1

		}
		fmt.Println("N2", n)
		n -= denoms[i]
		coins += 1

		if n == 0 {
			return coins
		}

	}

	return coins

}