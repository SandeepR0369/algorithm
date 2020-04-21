package main

import (
	"fmt"
	//"sort"
)

func main() {
	//sum([1,43,64,2,5,67,334,23], 69)
	y, t := sum([]int{1, 43, 64, 2, 5, 67, 334, 23}, 69)
	fmt.Println(y, t)
	x := Isitcorrect([]int{1, 43, 64, 2, 5, 67, 334, 23}, 169)
	fmt.Println(x)
/*	zz := Product([]int{1, 43, 64, 2, 5, 67, 334, 23})
	fmt.Println(zz)
	z1 := Product([]int{1, 2, 3, 4, 5})
	fmt.Println(z1)
	z2 := []int{1, 43, 64, 2, 5, 67, 334, 23}
	sort.Ints(z2)
	fmt.Println(z2)*/
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
