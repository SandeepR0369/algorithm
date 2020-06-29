package main

import (
	"fmt"
	"strings"
)

func main() {
	e := reversal("Hello, playground Sandeep")
	fmt.Println(e)
}

func reversal(a string) string {
	var b, c []string
	b = strings.Split(a, "")

	for i := len(b) - 1; i >= 0; i-- {
		c = append(c, b[i])
	}

	d := strings.Join(c, "")

	return d

}
