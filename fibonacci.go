package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibByLoop(50))
//	fmt.Println(Fibonacci(50))
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


func fibByLoop(n int) int {
	fibBox := []int{0, 1}
	for  i := 0; i < n; i++ {
		fmt.Println("F", i, fibBox[i], fibBox[i+1])
		v := fibBox[i] + fibBox[i+1]
		fmt.Println("V",v)
		fibBox = append(fibBox, v)
	}

	fmt.Println("J",fibBox[n])

	result := int(fibBox[n])
	//defer trackTime(time.Now(), result, "Loop")
	
	return result
}