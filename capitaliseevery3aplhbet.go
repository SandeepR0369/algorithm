package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
}

func CapitalizeEveryThirdAlphanumericChar(s string) string {

	jaffa := strings.Split(s, "")

	j := 0

	newstr := ""
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")

	for i := 0; i < len(jaffa); i++ {

		j += 1

		res := re.Match([]byte(jaffa[i]))

		if j%3 != 0 {
			//fmt.Println(j, jaffa[i])
			newstr = newstr + strings.ToLower(jaffa[i])

		}

		if j%3 == 0 {
			newstr = newstr + strings.ToUpper(jaffa[i])

		}

		if !res {
			//fmt.Println("not alphanumeric:", jaffa[i])
			j = j - 1

		}

	}

	return newstr
}
