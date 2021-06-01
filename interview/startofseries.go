package main

import (
	"fmt"
	"strconv"
	"strings"
)

//https://play.golang.org/p/7xfR29n56LC

//Return smallest number with same number of digits
//Example: 1999 return 1000

func main() {
	fmt.Println(Solution(11))
}

func Solution(N int) int {

	b := strconv.Itoa(N)
	c := strings.Split(b, "")
	sli := make([]int, len(c))

	for i := range c {
		sli[i], _ = strconv.Atoi(c[i])
	}

	if len(sli) == 1 {
		return 0
	}

	for z := 0; z < len(sli)-1; z++ {

		if sli[0] != 1 {
			sli[0] = 1
		}

		if sli[z+1] == 0 {
			continue
		}

		if sli[z+1] != 0 {
			sli[z+1] = 0
		}
	}

	sli2 := make([]string, len(c))

	for k := range sli {
		sli2[k] = strconv.Itoa(sli[k])
	}

	e := strings.Join(sli2, "")

	d, _ := strconv.Atoi(e)
	return d

}

