package main

import (
	"fmt"
)

// Reduce a number to 0 in minimum steps
//1) Decrease the number by 1, or
//2) Replace the number by its largest prime divisor(not 1 or itself)
// Example: 9 to 0
// 9 replaced with 3, 3 to 2, 2 to 1, 1 to 0 {Output: 4}

func main() {

	fmt.Println(MinSteps(32))

}

func MinSteps(number int) int {

	step := 0

	for k := number; k >= 0; k-- {

		for i := number - 1; i >= 2; i-- {

			if ifprime(i) == true {

				if number%i == 0 {
					//fmt.Println("1", i)
					number = i
					step += 1

				} else {
					continue
				}
			}
		}
		//fmt.Println("2", number)
		number -= 1

		step += 1

		if number == 0 {
			return step
		} else {
			continue
		}

		/*if number > 0 {
			continue
		}*/

	}

	return step

}

func ifprime(number int) bool {
	isPrime := true
	if number == 0 || number == 1 {
		return false
	} else {
		for i := 2; i <= number/2; i++ {
			if number == i {
				continue
			}

			if number%i == 0 {

				isPrime = false
				return false

			}
		}
		if isPrime == true {
			return true
		}
	}
	return false
}
