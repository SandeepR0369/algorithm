package main

import (
	"fmt"
)

//Generating Prime number upto given input

func main() {

	fmt.Println(GenPrime(15))

}

func GenPrime(number int) []int {

	primediv := make([]int, 0)

	for i := 2; i < number; i++ {
		if isprime(i) {
			primediv = append(primediv, i)
		}
	}

	//fmt.Println(primediv)

	return primediv

}

func isprime(number int) bool {
	isPrime := true
	if number == 0 || number == 1 {
		return false
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				isPrime = false
				return false
			}
		}
	}
	return isPrime
}
