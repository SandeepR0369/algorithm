package main

import (
	"fmt"
	"strings"
)

func main() {
	a1 := wordreverse("Hello jaffa playground")
	a2 := wordreverse("I am a boy")
	fmt.Println(a1)
	fmt.Println(a2)

}

func wordreverse(a string) string {
	var c []string

	b := strings.Split(a, " ")

	for i := len(b) - 1; i >= 0; i-- {

		c = append(c, b[i])

	}

	d := strings.Join(c, " ")

	return d

}
