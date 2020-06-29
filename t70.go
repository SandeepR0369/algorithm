package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//	Missing([]int{1, 43, 64, 2, 5, 67, 334, 23})
	a := Perfect(30314)
	fmt.Println(a)

}

//A number is considered perfect if its digits sum up to exactly 10.
//Given a positive integer n, return the n-th perfect number.
//For example, given 1, you should return 19. Given 2, you should return 28.

func Perfect(num int) int{
	
	res1 := strings.Split(strconv.Itoa(num), "")
	fmt.Println("R1", res1)

	var res2 []int
		for m := 0; m <= len(res1)-1; m++ {
			if s, err := strconv.Atoi(res1[m]); err == nil {
				res2 = append(res2, s)
			}
		}
	//	fmt.Printf("%T, %v", res2, res2)
	fmt.Println("R2", res2)

	sum := 0
	for k, _ := range res2{
		sum = sum+res2[k]
		if sum == 10{
			continue 
		}else if sum > 10 {
			return sum - 10
		}else if sum < 10 {
			return 10 - sum
		}
		fmt.Println("S1", sum)
	}
	fmt.Println("S1", sum)
	return 0
	}
