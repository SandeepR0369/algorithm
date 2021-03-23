package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello, playground")
	//sum := 12
	//a := []int{10, 2, 5, 1, 7, 6, 4}
	
	fmt.Println(NumberOfWaysToMakeChange(12, []int{10, 2, 5, 1, 7, 6, 4}))
	
	
	
func NumberOfWaysToMakeChange (n int, denoms []int) int{

	result_slice := make([][]int, 0)
	no_of_ways := 0

	for i := 0; i < len(a); i++ {
		temp_slice := make([]int, 0)
		if denoms[i] == n {
			temp_slice = append(temp_slice, denoms[i])
			no_of_ways += 1
			continue
		}
		res := n
		for j := i; j < len(denoms); j++ {

			if res > 0 {

				if denoms[j] <= res {
					res = res - denoms[j]

					temp_slice = append(temp_slice, denoms[j])

				}

				if res == 0 {

					result_slice = append(result_slice, temp_slice)
					res = n
					no_of_ways += 1
					temp_slice = make([]int, 0)
				}

			}
		}
	}
	fmt.Println(no_of_ways)
	fmt.Println(result_slice)
}
