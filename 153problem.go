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
	var index1, index2 []int
	if d == e {
		return 0
	}
	b := strings.Split(str, " ")
	for m, _ := range b {
		if b[m] == d {
			index1 = append(index1, m)
		}
		if b[m] == e {
			index2 = append(index2, m)
		}
	}

	var fin []int

	for i := 0; i <= len(index1)-1; i++ {
		for j := 0; j <= len(index2)-1; j++ {
			fin = append(fin, index2[j]-index1[i])
		}
	}
	sort.Ints(fin)
	return fin[0] - 1
}

func Absolutedistance(str, d, e string) float64 {
	var index1, index2 []int
	if d == e {
		return 0
	}

	b := strings.Split(str, " ")

	for m, _ := range b {
		if b[m] == d {
			index1 = append(index1, m)
		}
		if b[m] == e {
			index2 = append(index2, m)
		}
	}

	var fin []float64
	for i := 0; i <= len(index1)-1; i++ {
		for j := 0; j <= len(index2)-1; j++ {
			fin = append(fin, math.Abs(float64(index2[j]-index1[i])))
		}

	}
	sort.Float64s(fin)
	return fin[0] - 1
}
