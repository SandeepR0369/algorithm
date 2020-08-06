package main

import (
	"fmt"
)

func main() {
	fmt.Println(Fibonacci(11))
}

func Fibonacci(a int) int {

	if a == 0 {
		return 0
	}
	if a == 1 {
		return 1
	}

	x := Fibonacci(a-1) + Fibonacci(a-2)
	return x
}
