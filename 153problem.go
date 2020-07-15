package main

import (
	"fmt"
	"strings"
)

//Find an efficient algorithm to find the smallest distance (measured in number of words) between any two given words in a string.
//For example, given words "hello", and "world" and a text content of "dog cat hello cat dog dog hello cat world",
//return 1 because there's only one word "cat" in between the two words.

func main() {
	a1 := distance("fuck catie hello cat dog dog hello cat world", "hello", "world")
	fmt.Println(a1)
}

func distance(str, d, e string) int {

	if d == e {
		return 0
	}

	b := strings.Split(str, " ")


	for m, _ := range b {
		if b[m] == d {
			index1 = append(index1, m)
		}
		if b[m] == e{
			index2 = append(index2, m)
		}
		
	}
	

	for i, _ := range b {

		if b[i] == d {
			fmt.Println(i,d)
		}

		if b[i] == d {
			for m, _ := range b {

				if b[m] == e {
					return m - i - 1
				}
			}
		}
	}
	return 0
}
