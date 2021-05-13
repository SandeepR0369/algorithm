package main

import (
	"fmt"
	"sort"
)

/*-> Find the next highest number by rearranging the digits of the given number. If it is not possible then return -1
Input : 4321
Output: -1
//54123
Input : 3142
Output: 3214

Input : 8579
Output: 8597*/

func main() {

	fmt.Println(rearrange([]int{4, 1, 3, 2}))
	fmt.Println(rearrange([]int{5, 1, 4, 3, 2}))
	fmt.Println(rearrange([]int{4, 3, 2, 1}))

}

func rearrange(a []int) []int {

	swap := 0
	for j := len(a) - 1; j >= 1; j-- {

		if a[j] < a[j-1] {
			continue
		}
		if a[j] > a[j-1] {
			a[j], a[j-1] = a[j-1], a[j]
			sort.Ints(a[j:])
			swap += 1

		}

		if swap == 1 {
			return a
		}

	}

	return []int{}

}

/*
func main() {
	fmt.Println(rearrange(51432))

}

func rearrange(a int) []int {

	b := strconv.Itoa(a)

	c := strings.Split(b, "")

	sli := make([]int, len(c))

	for i := range c {
		sli[i], _ = strconv.Atoi(c[i])
	}
	
	swap := 0
	for j := len(sli) - 1; j >= 1; j-- {

		if sli[j] < sli[j-1] {
			continue
		}
		if sli[j] > sli[j-1] {
			sli[j], sli[j-1] = sli[j-1], sli[j]
			sort.Ints(sli[j:])
			swap += 1

		}

		if swap == 1 {
			return sli
		}

	}

	return []int{}

}
*/
