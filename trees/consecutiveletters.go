package main

import (
	"fmt"
	"strings"
)

func main() {

	//fmt.Println(counter("Hello, playground"))
	fmt.Println(counter("AAAAABBBCCC"))

	fmt.Println(counter("AAAAABCDDD"))

	//fmt.Println(strings.Split(("Hello, playground"),""))
}

func Splitter(ab string) []string {

	var sli []string
	a := strings.Split(ab, "")

	for i := 0; i < len(a); i++ {
		sli = append(sli, a[i])
	}

	return sli
}

func Counter(sli []string) string {

	jaffa := make(map[string]int)

	count := 1

	for j := 0; j < len(str); j++ {
		if str[j] == str[j+1] {
			count += 1
		}

		if str[j] != str[j+1] {
			count
		}

	}

}
