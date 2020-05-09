package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	y := Perfect(30)
	fmt.Println("Y", y)
}

//A number is considered perfect if its digits sum up to exactly 10.
//Given a positive integer n, return the n-th perfect number.
//For example, given 1, you should return 19. Given 2, you should return 28.

//MAYBE 3 -> 37 4 -> 46

func Perfect(num int)int {
	sum := 0
	total := 10
	res1 := strings.Split(strconv.Itoa(num), "")

	var res2 []int

	for m := 0; m <= len(res1)-1; m++ {
		if s, err := strconv.Atoi(res1[m]); err == nil {
			res2 = append(res2, s)
		}
	}
	for k, _ := range res2 {
		sum = sum + res2[k]
	}

	if sum > total {
		fmt.Println("Not Perfect Number")
	}
	if sum == total {
		fmt.Println("Perfect Number")
	}
	if sum < total {
		wq := append(res2, total-sum)
		wq1 := SliceJoin(wq)
		return wq1
	}
	return 0
}

func SliceJoin(slice []int) int {
	output := make([]string, len(slice))
	for k, v := range slice {
		output[k] = strconv.Itoa(v)
	}
	if b, err := strconv.Atoi(strings.Join(output, "")); err == nil {
		return b
	}
	return 0
}
