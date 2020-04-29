package main

import (
	"fmt"
	"strings"
)

var str string
var n int

func main() {
	fmt.Print("Please Enter the text you want to check: ")
	fmt.Scanf("%s", &str)
	if palindrome(strings.ToUpper(str)) == true {
		fmt.Println(str, "is Palindrome")
	} else {
		fmt.Println(str, "not Palindrome")
	}
}

func reversal(ac string) string {
	var str1 string
	res1 := strings.Split(ac, "")
	for _, v := range res1 {
		str1 = v + str1
	}
	/*for i := len(res1)-1; i >= 0; i--{                       //This is also correct
		str1 += res1[i]

	}*/
	return str1
}

func palindrome(ac string) bool {
	//if strings.ToUpper(ac) == strings.ToUpper(reversal(str)){   //This is also correct
	r_str := reversal(str)
	if strings.EqualFold(r_str, ac) {
		return true
	}
	return false
}
