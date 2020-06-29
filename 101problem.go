package main

import (
	"fmt"
)

//Given an even number (greater than 2), return two prime numbers whose sum will be equal to the given number.

func main() {
	y:= Primes(17)
	fmt.Println(y)

}

func Primes( a int) (b,c int){
	//var a int
	if a%2 != 0 {
		fmt.Println("Please check the Number, its not Even")
	} else if a%2 == 0 && a > 2{
		fmt.Println("Number is Perfect")
	}
}
