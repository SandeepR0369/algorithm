package main

import (
	"fmt"
	"strings"
)

//Given a string of words delimited by spaces, reverse the words in string. For example, given "hello world here", return "here world hello"
//Follow-up: given a mutable string representation, can you perform this operation in-place?

func main() {
	x := wordreversal("Given a string of words")
	fmt.Println(x)
}

func wordreversal(str string) string {
	var a, b []string
	a = strings.Split(str, " ")
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	d := strings.Join(b, " ")
	return d
}
