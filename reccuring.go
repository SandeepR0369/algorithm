package main

import (
	"fmt"
	"strings"
)
//Given a string, return the first recurring character in it, or null if there is no recurring character.
//For example, given the string "acbbac", return "b". Given the string "abcdef", return null.

func main() {
	str := "sandeep"
	A := reccur(str)
	fmt.Println("first reccur letter:", A)
}

func reccur(str string) string {
	a := strings.Split(str, "")
	b := make(map[string]int)

	for _, v := range a {
		if val, ok := b[v]; ok {
			val = val + 1
			b[v] = val
			return v
		} else {
			b[v] = 1
		}
	}

/*	for m, n := range b {
		if n > 1 {
			return m
		}
	}
*/
	return "Null"
}
