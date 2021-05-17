package main

import (
	"fmt"
	"strconv"
	"strings"
)
//A strong number is a special number whose sum of the factorial digits is equal to the original number.
func main() {

	/*StrongNumber(145)
	StrongNumber(141)
	StrongNumber(140)
	fmt.Println(fact(0))*/

	for i := 1; i <= 100000; i++ {
		StrongNumber(i)
	}

}

func StrongNumber(n int) {

	b := strconv.Itoa(n)
	c := strings.Split(b, "")
	sli := make([]int, len(c))

	for i := range c {
		sli[i], _ = strconv.Atoi(c[i])
	}
	//var s2 []int
	sum := 0
	for j := 0; j < len(sli); j++ {

		//s2 = append(s2, fact(sli[j]))
		sum = sum + fact(sli[j])
	}

	if n == sum {
		fmt.Println("Strong Number", n, sum)
	}
	/*else {
		fmt.Println("Not Strong Number", n, sum)
	}*/
}

func fact(n int) int {
	out := 1

	if n == 0 {
		return 1
	}

	for i := n; i >= 1; i-- {
		out *= i
	}
	return out
}

