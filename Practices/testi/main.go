package main

import "fmt"

func main() {
	fmt.Println(add(2, 7))
}

func add(a, b int) int {
	out := a + b
	return out
}
