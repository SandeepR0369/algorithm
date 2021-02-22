package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Write a function that takes a natural number as input and returns the number of digits the input has.
//Constraint: don't use any loops.

func main(){
	fmt.Println(numberofint(10))
}
/*
func numberofint(natural int)int{
counter := 0

	if natural <= 0 {
		return 0
	}

	if natural > 0 {
		count := natural/10
		counter = counter+1
	}
return counter
}
*/

func numberofint(natural int)int {
	
	//natural := 123
	
	str := strconv.Itoa(natural)
	//fmt.Println(str)
	
	sl := strings.Split(str, "")
	//fmt.Println(sl)
	
	return len(sl)
}
