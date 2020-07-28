package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	a := []string{"a", "b", "f", "h", "z", "a", "b", "f", "c", "a"}

	b := make(map[string]int)

		for _, v := range a {
			_, ok  := b[v] 
			if ok {
				b[v] += 1
			}else {
				b[v] = 1
			}
		}
	fmt.Println("AA", b)
}
